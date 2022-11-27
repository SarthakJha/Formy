package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sarthakjha/Formy/internal/model"
	"github.com/sarthakjha/Formy/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/sheets/v4"
)

type question struct{
	Question string `json:"question_string"`
}
type questionRequest struct {
	QuestionString []question `json:"questions"`
	IsGmailNotificationEnabled bool `json:"is_gmail_notif_enabled" bson:"is_gmail_notif_enabled"`
	IsSheetEnabled bool `json:"is_sheet_enabled" bson:"is_sheet_enabled"`
}

func (handler *Handler) CreateForm(w http.ResponseWriter, r *http.Request){
	/**
		{
			questions: [{
				question_string: "lalalal?"
			}]
		}
	*/
	
	questions := questionRequest{}
	err := json.NewDecoder(r.Body).Decode(&questions)
	if err!=nil{
		log.Println("ERROR: cant decode questions")
		w.WriteHeader(500)
		fmt.Fprint(w,"internal error")
		return
	}
	questionObject := []model.Question{}
	for i := 0; i < len(questions.QuestionString); i++ {
		question := model.Question{
			QuestionId: primitive.NewObjectID(),
			QuestionString: questions.QuestionString[i].Question,
		}
		questionObject = append(questionObject,question)
	}
	// create spreadsheet here if option is checked and set columns
	sheetLink := ""
	if questions.IsSheetEnabled {
		s := handler.SheetsClient.Spreadsheets.Create(&sheets.Spreadsheet{})
		sheet,err := s.Do() 
		if err!=nil{
			log.Println("ERROR: ", err.Error())
		}
		// handler.SheetsClient.Spreadsheets.Values.a
		sheetLink = sheet.SpreadsheetUrl
	}
	
	err=repository.AddFormToDatabase(handler.Db,questionObject, questions.IsSheetEnabled,questions.IsGmailNotificationEnabled, sheetLink)
	if err!=nil{
		log.Println("ERROR: error saving questions")
		w.WriteHeader(500)
		fmt.Fprint(w,"internal error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "question recieved")
}