package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/DuncanScu/rsvp_system/database"
	"github.com/DuncanScu/rsvp_system/initializers"
	usecases "github.com/DuncanScu/rsvp_system/use_cases"
	"github.com/aws/aws-lambda-go/events"
)

type GetPartyEventsHandler struct{}

func (h GetPartyEventsHandler) CanExecute(r events.APIGatewayProxyRequest) bool {
	segmented_path := strings.Split(r.Path, "/")
	isPartyPath := segmented_path[1] == "party" && segmented_path[3] == "events"
	return isPartyPath &&
		r.HTTPMethod == http.MethodGet &&
		len(r.PathParameters) == 1
}

func (h GetPartyEventsHandler) Execute(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uniqueId, found := request.PathParameters["unique_id"]
	if !found {
		return events.APIGatewayProxyResponse{}, errors.New("unique_id not found in path parameters")
	}

	uc := &usecases.GetEventsByParty{
		EventsRepo: &database.EventDb{Db: initializers.DB},
		PartyRepo:  &database.PartyDb{Db: initializers.DB},
	}

	partyEvents, err := uc.Execute(uniqueId)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	b, err := json.Marshal(partyEvents)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, err
}
