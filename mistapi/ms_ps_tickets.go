package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// MSPsTickets represents a controller struct.
type MSPsTickets struct {
    baseController
}

// NewMSPsTickets creates a new instance of MSPsTickets.
// It takes a baseController as a parameter and returns a pointer to the MSPsTickets.
func NewMSPsTickets(baseController baseController) *MSPsTickets {
    mSPsTickets := MSPsTickets{baseController: baseController}
    return &mSPsTickets
}

// ListMspTickets takes context, mspId, start, end, duration as parameters and
// returns an models.ApiResponse with []models.Ticket data and
// an error if there was an issue with the request or response.
// Get List of Tickets of a MSP
func (m *MSPsTickets) ListMspTickets(
    ctx context.Context,
    mspId uuid.UUID,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.Ticket],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/tickets", mspId),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result []models.Ticket
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Ticket](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountMspTickets takes context, mspId, distinct as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count tickets
func (m *MSPsTickets) CountMspTickets(
    ctx context.Context,
    mspId uuid.UUID,
    distinct *models.MspTicketsCountDistinctEnum) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/tickets/count", mspId),
    )
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}
