package mistapigo

import (
	"context"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// AdminsLookup represents a controller struct.
type AdminsLookup struct {
	baseController
}

// NewAdminsLookup creates a new instance of AdminsLookup.
// It takes a baseController as a parameter and returns a pointer to the AdminsLookup.
func NewAdminsLookup(baseController baseController) *AdminsLookup {
	adminsLookup := AdminsLookup{baseController: baseController}
	return &adminsLookup
}

// Lookup takes context, body as parameters and
// returns an models.ApiResponse with models.ResponseLoginLookup data and
// an error if there was an issue with the request or response.
// Login Lookup
func (a *AdminsLookup) Lookup(
	ctx context.Context,
	body *models.EmailString) (
	models.ApiResponse[models.ResponseLoginLookup],
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/login/lookup")
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "user does not exist"},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	var result models.ResponseLoginLookup
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseLoginLookup](decoder)
	return models.NewApiResponse(result, resp), err
}
