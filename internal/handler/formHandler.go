package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sarthakjha/Formy/internal/model"
	"github.com/sarthakjha/Formy/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/sheets/v4"
)

type question struct{
	Question string `json:"question_string"`
}
type formRequest struct {
	QuestionString []question `json:"questions"`
	IsGmailNotificationEnabled bool `json:"is_gmail_notif_enabled" bson:"is_gmail_notif_enabled"`
	IsSheetEnabled bool `json:"is_sheet_enabled" bson:"is_sheet_enabled"`
	Title string `json:"title" bson:"title"`
}

func (handler *Handler) CreateForm(w http.ResponseWriter, r *http.Request){
	formSubmission := formRequest{}
	err := json.NewDecoder(r.Body).Decode(&formSubmission)
	if err!=nil{
		log.Println("ERROR: cant decode questions")
		w.WriteHeader(500)
		fmt.Fprint(w,"internal error")
		return
	}
	questionObject := []model.Question{}
	for i := 0; i < len(formSubmission.QuestionString); i++ {
		question := model.Question{
			QuestionId: primitive.NewObjectID(),
			QuestionString: formSubmission.QuestionString[i].Question,
		}
		questionObject = append(questionObject,question)
	}
	
	// create spreadsheet here if option is checked and set columns
	sheetLink := ""
	if formSubmission.IsSheetEnabled {
		s := handler.SheetsClient.Spreadsheets.Create(&sheets.Spreadsheet{
			Properties: &sheets.SpreadsheetProperties{
				Title: formSubmission.Title,
			},
		})
		sheet,err := s.Do() 

		// handling spreadsheet's permission
		filePermission := &drive.Permission{
			Type: "anyone",
			Role: "writer",
		}
		d := handler.DriveClient.Permissions.Insert(sheet.SpreadsheetId,filePermission)
		_,err = d.Do()
		if err!=nil{
			log.Println("ERROR: ", err.Error())
		}
		// handler.SheetsClient.Spreadsheets.Values.a
		sheetLink = sheet.SpreadsheetUrl
	}
	
	err=repository.AddFormToDatabase(handler.Db,questionObject, formSubmission.IsSheetEnabled,formSubmission.IsGmailNotificationEnabled, sheetLink)
	if err!=nil{
		log.Println("ERROR: error saving questions")
		w.WriteHeader(500)
		fmt.Fprint(w,"internal error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "question recieved")
}