package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "mistapi/errors"
    "mistapi/models"
)

// SelfAlarms represents a controller struct.
type SelfAlarms struct {
    baseController
}

// NewSelfAlarms creates a new instance of SelfAlarms.
// It takes a baseController as a parameter and returns a pointer to the SelfAlarms.
func NewSelfAlarms(baseController baseController) *SelfAlarms {
    selfAlarms := SelfAlarms{baseController: baseController}
    return &selfAlarms
}

// ListAlarmSubscriptions takes context as parameters and
// returns an models.ApiResponse with []models.ResponseSelfSubscription data and
// an error if there was an issue with the request or response.
// Get List of all the subscriptions
func (s *SelfAlarms) ListAlarmSubscriptions(ctx context.Context) (
    models.ApiResponse[[]models.ResponseSelfSubscription],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/self/subscriptions")
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
    var result []models.ResponseSelfSubscription
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ResponseSelfSubscription](decoder)
    return models.NewApiResponse(result, resp), err
}
