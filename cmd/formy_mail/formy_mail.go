package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sarthakjha/Formy/internal/model"
	"github.com/sarthakjha/Formy/internal/queue"
)

func main(){
	if err := godotenv.Load("../../test.env"); err!= nil {
		log.Fatalln(err.Error())
	}

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