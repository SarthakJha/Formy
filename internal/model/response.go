package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type Response struct{
	ResponseId primitive.ObjectID
	QuestionId primitive.ObjectID 
	ResponseText string
}