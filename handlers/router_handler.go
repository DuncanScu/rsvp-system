package handlers

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	CanExecute(events.APIGatewayProxyRequest) bool
	Execute(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	handlers := []Handler{
		&GetPartyMembersHandler{},
		&GetPartyEventsHandler{},
		&RsvpToEventHandler{},
	}

	for _, handler := range handlers {
		if handler.CanExecute(request) {
			slog.Info("Executing handler: %T", handler)
			return handler.Execute(ctx, request)
		}
	}

	return events.APIGatewayProxyResponse{}, nil
}
