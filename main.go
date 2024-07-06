package main

import (
	"github.com/DuncanScu/rsvp_system/handlers"
	"github.com/DuncanScu/rsvp_system/initializers"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	initializers.ConnectToDB()
	initializers.MigrateDb()
	// initializers.SeedDb()
}

func main() {
	// res, err := handlers.HandleRequest(nil, events.APIGatewayProxyRequest{
	// 	Path:       "/party/{unique_id}/members",
	// 	HTTPMethod: http.MethodGet,
	// 	PathParameters: map[string]string{
	// 		"unique_id": "VVoveui9QC",
	// 	},
	// })

	// res, err := handlers.HandleRequest(nil, events.APIGatewayProxyRequest{
	// 	Path:       "/party/{unique_id}/events",
	// 	HTTPMethod: http.MethodGet,
	// 	PathParameters: map[string]string{
	// 		"unique_id": "VVoveui9QC",
	// 	},
	// })

	// res, err := handlers.HandleRequest(context.Background(), events.APIGatewayProxyRequest{
	// 	Path:       "/event/0/rsvp",
	// 	HTTPMethod: http.MethodPost,
	// 	PathParameters: map[string]string{
	// 		"event_id": "0",
	// 	},
	// 	Body: `{"user_id":1,"response":"Pending"}`,
	// })

	lambda.Start(handlers.HandleRequest)
}
