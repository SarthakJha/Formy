package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sarthakjha/Formy/internal/api"
)


func main()  {
	if err := godotenv.Load("prod.env"); err!= nil {
		log.Fatalln(err.Error())
	}

	r := mux.NewRouter()
	api.SetupRoutes(r)
	server := http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", "5001"),
		Handler: r,
	}

	go func(){
		log.Println("server listening to port: ", server.Addr)
		if err:=server.ListenAndServe(); err!=nil {
			log.Fatalln("ERROR: ",err.Error())
		}
	}()

	sig := make(chan os.Signal,1)

	signal.Notify(sig, os.Interrupt)
	<-sig

	// graceful shutdown
	ctx,cancel := context.WithCancel(context.Background())
	defer server.Shutdown(ctx)
	defer cancel()

	log.Println("SERVER SHUTDOWN")
}