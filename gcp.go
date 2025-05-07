package main

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func setUpGoogleSheetsAPI() *sheets.Service {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: .env file not found. Falling back to environment variables.")
	}

	// load b64 encoded GCP Sheets api credentials
	encodedCredentials := os.Getenv("GCP_SHEETS_CREDENTIALS_BASE64")
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		log.Fatalf("Failed to decode base64 credentials: %v", err)
	}

	config, err := google.JWTConfigFromJSON(decodedCredentials, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		log.Fatalf("Failed to parse credentials: %v", err)
	}

	ctx := context.Background()

	// Use service account or OAuth2 credentials
	client := config.Client(ctx)

	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
	}

	return sheetsService
}
