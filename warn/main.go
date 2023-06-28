package main

import (
	"context"
	"log"
	"os"
	"strings"
	"warn/warn/models"
	"warn/warn/states"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {

	log.Print("Starting")

	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + 
		" user=" + user + 
		" password=" + password + 
		" dbname=" + database + 
		" port=" + port 

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
	  panic("failed to connect database: " + err.Error())
	}

	log.Print("Connected to DB")


	// Migrate the schema
	
	db.AutoMigrate(&models.WarnNoticeModel{})

	log.Print("Finished Migration")

	state := os.Getenv("STATE")

	if strings.Compare(state, "maryland") == 0 {
		states.MarylandHandler()
		log.Print("Finished state")
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            state,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}