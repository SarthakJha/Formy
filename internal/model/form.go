package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type Form struct{
	FormId primitive.ObjectID `json:"id" bson:"_id"`
	Questions []primitive.ObjectID`json:"question_ids" bson:"question_ids"`
}