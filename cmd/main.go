package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)


func main()  {
	if err := godotenv.Load("prod.env"); err!= nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("hello world")
}