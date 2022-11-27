package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sarthakjha/Formy/internal/model"
	"github.com/sarthakjha/Formy/internal/queue"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main(){
	if err := godotenv.Load("prod.env"); err!= nil {
		log.Fatalln(err.Error())
	}
	ctx_auth, cancel_auth := context.WithCancel(context.Background())
	defer cancel_auth()
	credBytes, err:= base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_OAUTH"))
	if err != nil{
		log.Fatalln("ERROR: decoding auth creds")
	}
	config,err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil{
		log.Fatalln("ERROR: decoding auth creds1, ", err.Error())
	}
	
	client := config.Client(ctx_auth)



	srv, err := sheets.NewService(ctx_auth, option.WithHTTPClient(client))

	fmt.Println(srv)
	redisClient := queue.ConnectQueue("localhost:6379")
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	subs := redisClient.Subscribe(ctx, string(queue.EMAIL_NOTIF))

	resp := make(chan model.Response)

	go func() {
		fmt.Println("routine working")
		// buisness logic
		for a := range(resp) {
			fmt.Println(a)
		}
	}()

	for{
		fmt.Println("reciecving start..")
		msg,err := subs.ReceiveMessage(ctx)
		if err !=nil{
			log.Fatalln("ERROR: ", err.Error())
		}
		responseobj := model.Response{}
		if err:= json.Unmarshal([]byte(msg.Payload), &responseobj);err!=nil{
			log.Fatalln("ERROR: ", err.Error())
		}
		resp<- responseobj
	}

	
	
}