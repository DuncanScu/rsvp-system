package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/DuncanScu/rsvp_system/database"
	"github.com/DuncanScu/rsvp_system/dtos"
	"github.com/DuncanScu/rsvp_system/entities/enums/status"
	"github.com/DuncanScu/rsvp_system/initializers"
	usecases "github.com/DuncanScu/rsvp_system/use_cases"
	"github.com/aws/aws-lambda-go/events"
)

type RsvpToEventHandler struct {
}

func (h *RsvpToEventHandler) CanExecute(r events.APIGatewayProxyRequest) bool {
	segmented_path := strings.Split(r.Path, "/")
	isUpdateEventRsvpPath := segmented_path[1] == "event" && segmented_path[3] == "rsvp"

	return isUpdateEventRsvpPath && r.HTTPMethod == "POST"
}

func (h *RsvpToEventHandler) Execute(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var err error
	var dto dtos.RsvpToEventDto

	err = json.Unmarshal([]byte(request.Body), &dto)
	eventId, found := request.PathParameters["event_id"]
	if !found {
		return events.APIGatewayProxyResponse{}, errors.New("event_id not found in path parameters")
	}

	parsedEventId, err := strconv.ParseUint(eventId, 10, 0)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	uc := &usecases.RsvpToEvent{
		RsvpRepo: &database.RsvpDb{Db: initializers.DB},
	}

	err = uc.Execute(uint(parsedEventId), dto.UserId, status.StatusType(dto.Status))
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{}, nil
}
