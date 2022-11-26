package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct{
	QuestionId primitive.ObjectID `json:"id" bson:"_id"`
	FormId primitive.ObjectID `json:"form_id" bson:"form_id"` 
	QuestionString string `json:"question_string" bson:"question_string"`
}