package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type Form struct{
	FormId primitive.ObjectID `json:"id" bson:"_id"`
	Questions []primitive.ObjectID`json:"question_ids" bson:"question_ids"`
	IsGmailNotificationEnabled bool `json:"is_gmail_notif_enabled" bson:"is_gmail_notif_enabled"`
	IsSheetEnabled bool `json:"is_sheet_enabled" bson:"is_sheet_enabled"`
	GoogleSheetLink string `json:"google_sheet_link,omitempty" bson:"google_sheet_link,omitempty"`
	Title string `json:"title" bson:"title"`
	SheetId string `json:"sheet_id" bson:"sheet_id"`
}