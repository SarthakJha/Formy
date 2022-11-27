package repository

import (
	"context"
	"log"
	"os"

	"github.com/sarthakjha/Formy/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func AddFormToDatabase(client *mongo.Client, questions []model.Question, isSheetEnabled,isGmailNotifEnabled bool)error{

	formId := primitive.NewObjectID()
	questionIds := []primitive.ObjectID{}
	for i := 0; i < len(questions); i++ {
		questions[i].FormId = formId
		// NOTE:- Need atomic failsafe mechanism
		addQuestionToDatabse(client,questions[i])
		questionIds = append(questionIds,questions[i].QuestionId)
	}

	form := model.Form{
		FormId: formId,
		Questions: questionIds,
		IsGmailNotificationEnabled: isGmailNotifEnabled,
		IsSheetEnabled: isSheetEnabled,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("FORM_COLLECTION")).InsertOne(ctx, form)
	if err!=nil{
		log.Println("ERROR SAVING FORM"+err.Error())
		return err
	}
	return nil
}

func addQuestionToDatabse(client *mongo.Client, question model.Question){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("QUESTION_COLLECTION")).InsertOne(ctx, question)
	if err!=nil{
		log.Println("ERROR SAVING QUESTION: "+err.Error())
	}
}

func AddResponseToDatabase(client *mongo.Client, response model.Response)error{
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("RESPONSE_COLLECTION")).InsertOne(ctx, response)	
	if err!=nil{
		log.Println("ERROR SAVING RESPONSE: "+err.Error())
		return err
	}
	return nil
}