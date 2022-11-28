package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/sarthakjha/Formy/internal/model"
	"github.com/sarthakjha/Formy/internal/queue"
)

func main(){
	if err := godotenv.Load("prod.env"); err!= nil {
		log.Fatalln(err.Error())
	}
	
	redisClient := queue.ConnectQueue("localhost:6379")
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	subs := redisClient.Subscribe(ctx, string(queue.EMAIL_NOTIF))

	resp := make(chan model.Response)

	go func() {
		log.Println("routine working")
		// buisness logic
		for a := range(resp) {
			fmt.Println(a)
			from := os.Getenv("EMAIL_ADDR")
			password := os.Getenv("EMAIL_PASS")
		
			toEmailAddress := a.ResponseText
			to := []string{toEmailAddress}
		
			host := "smtp.gmail.com"
			port := "587"
			address := host + ":" + port
		
			subject := "no-reply Thanks for the submission\n"
			body := "Thanks!"
			message := []byte(subject + body)
		
			auth := smtp.PlainAuth("", from, password, host)
		
			err := smtp.SendMail(address, auth, from, to, message)
			if err != nil {
				log.Println("ERROR: ", err.Error() )
			}
		}
	}()

	for{
		log.Println("reciecving start..")
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