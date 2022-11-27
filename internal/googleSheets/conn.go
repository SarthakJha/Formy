package googleSheets

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ConnectToGoogleServices()(*sheets.Service,error){
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
	return sheets.NewService(ctx_auth, option.WithHTTPClient(client))

}

func ConnectToDrive()(*drive.Service,error){
	ctx_auth, cancel_auth := context.WithCancel(context.Background())
	defer cancel_auth()
	credBytes, err:= base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_OAUTH"))
	if err != nil{
		log.Fatalln("ERROR: decoding auth creds")
	}
	config,err := google.JWTConfigFromJSON(credBytes, drive.DriveFileScope)
	if err != nil{
		log.Fatalln("ERROR: decoding auth creds1, ", err.Error())
	}
	
	client := config.Client(ctx_auth)
	return drive.NewService(ctx_auth, option.WithHTTPClient(client))

}
