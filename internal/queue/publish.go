package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/sarthakjha/Formy/internal/model"
)


func PublishResponseForGoogleSheet(res model.Response, queue *redis.Client){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	payload,err := json.Marshal(res)
	if err != nil {
		log.Fatalln("ERROR: converting payload to json")
	}
	s := queue.Publish(ctx,string(SHEETS),payload)
	if s.Err() != nil {
		log.Fatalln("ERROR1: ", s.Err().Error())
	}
}

func PublishResponseForEmailNotif(res model.Response, queue *redis.Client){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	payload,err := json.Marshal(res)
	if err != nil {
		log.Fatalln("ERROR: converting payload to json")
	}
	s := queue.Publish(ctx,string(EMAIL_NOTIF),payload)
	if s.Err() != nil {
		log.Fatalln("ERRO2R: ", s.Err().Error())
	}
}