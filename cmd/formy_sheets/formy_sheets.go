package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sarthakjha/Formy/internal/googleSheets"
	"github.com/sarthakjha/Formy/internal/model"
	"github.com/sarthakjha/Formy/internal/queue"
	"google.golang.org/api/sheets/v4"
)

func main(){
	if err := godotenv.Load("prod.env"); err!= nil {
		log.Fatalln(err.Error())
	}

	redisClient := queue.ConnectQueue("localhost:6379")
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	subs := redisClient.Subscribe(ctx, string(queue.SHEETS))

	resp := make(chan []model.Response)
	conn,err := googleSheets.ConnectToGoogleServices()
	if err != nil{
		log.Default().Fatalln("ERROR: connecting to google services")
	}
	go func() {
		fmt.Println("routine working")
		// buisness logic
		for a := range(resp) {
			sheetId :=  a[0].Form.SheetId
			respsTextArr := make([][]interface{},len(a))
			for _,j := range(a){
				respsTextArr[0] = append(respsTextArr[0], j.ResponseText)
			}
			addCall:=conn.Spreadsheets.Values.Append(sheetId, "Sheet1",&sheets.ValueRange{
				MajorDimension: "ROWS",
				Values: respsTextArr,
			}).ValueInputOption("RAW")
			_,err:=addCall.Do()
			if err != nil {
				log.Println("ERROR: ",err.Error())
			}
			log.Println("row appended")
		}
	}()

	for{
		msg,err := subs.ReceiveMessage(ctx)
		if err !=nil{
			log.Fatalln("ERROR: ", err.Error())
		}
		responseobj := []model.Response{}
		if err:= json.Unmarshal([]byte(msg.Payload), &responseobj);err!=nil{
			log.Fatalln("ERROR: ", err.Error())
		}
		resp<- responseobj
	}

	
	
}