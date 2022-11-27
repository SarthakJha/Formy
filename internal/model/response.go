package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type Response struct{
	ResponseId primitive.ObjectID `json:"id" bson:"_id"`
	Form Form `json:"form" bson:"form"` 
	QuestionId primitive.ObjectID `json:"question_id" bson:"question_id"`
	ResponseText string	`json:"response_string" bson:"response_string"`
}