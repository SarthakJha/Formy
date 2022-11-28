package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sarthakjha/Formy/internal/model"
	"github.com/sarthakjha/Formy/internal/queue"
	"github.com/sarthakjha/Formy/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type formResponseRequest struct{
	Response []res `json:"response"`
}

type res struct{
	ResponseString string `json:"response_string"`
	Question model.Question `json:"question_id"`
	Form model.Form `json:"form"`
}
func (handler *Handler) CreateResponse(w http.ResponseWriter, r *http.Request)  {
	
	res := formResponseRequest{}
	err:=json.NewDecoder(r.Body).Decode(&res)
	if err!=nil{
		log.Println("ERROR: cant decode response")
		w.WriteHeader(500)
		fmt.Fprint(w,"internal error")
		return
	}
	responseArr := []model.Response{}

	for i := 0; i < len(res.Response); i++ {
		if err!=nil{
			log.Println("ERROR: issue with objectID conversion")
			w.WriteHeader(500)
			fmt.Fprint(w,"internal error")
			return
		}
		responseObject := model.Response{
			Question: res.Response[i].Question,
			ResponseId: primitive.NewObjectID(),
			ResponseText: res.Response[i].ResponseString,
			Form: res.Response[i].Form,
		}
		err = repository.AddResponseToDatabase(handler.Db,responseObject)
		responseArr = append(responseArr, responseObject)

		// handle publishing of messages
		if responseObject.Form.IsGmailNotificationEnabled &&  responseObject.Question.IsResponseEmail {
			queue.PublishResponseForEmailNotif(responseObject,handler.Queue)
			log.Println("message for email published")
		}

		if err!=nil{
			log.Println("ERROR: cant save response")
			w.WriteHeader(500)
			fmt.Fprint(w,"internal error")
			return
		}
	}
	if res.Response[0].Form.IsSheetEnabled {
		queue.PublishResponseForGoogleSheet(responseArr,handler.Queue)
		log.Println("message for sheet published")
	}
	
	fmt.Fprint(w,"saved your responses")

	
}