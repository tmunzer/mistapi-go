package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// MSPsLicenses represents a controller struct.
type MSPsLicenses struct {
	baseController
}

// NewMSPsLicenses creates a new instance of MSPsLicenses.
// It takes a baseController as a parameter and returns a pointer to the MSPsLicenses.
func NewMSPsLicenses(baseController baseController) *MSPsLicenses {
	mSPsLicenses := MSPsLicenses{baseController: baseController}
	return &mSPsLicenses
}

// ClaimMspLicence takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.ResponseClaimLicense data and
// an error if there was an issue with the request or response.
// Claim an Order by Activation Code
func (m *MSPsLicenses) ClaimMspLicence(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.CodeString) (
	models.ApiResponse[models.ResponseClaimLicense],
	error) {
	req := m.prepareRequest(ctx, "POST", fmt.Sprintf("/api/v1/msps/%v/claim", mspId))
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Response when the key is invalid (or already used)"},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ResponseClaimLicense
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseClaimLicense](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListMspLicenses takes context, mspId as parameters and
// returns an models.ApiResponse with models.License data and
// an error if there was an issue with the request or response.
// Get List of Msp Licenses
func (m *MSPsLicenses) ListMspLicenses(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[models.License],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/licenses", mspId),
	)
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
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.License
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.License](decoder)
	return models.NewApiResponse(result, resp), err
}

// MoveOrDeleteMspLicenseToAnotherOrg takes context, mspId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Move or Delete MSP Licenses
func (m *MSPsLicenses) MoveOrDeleteMspLicenseToAnotherOrg(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.MspLicenseAction) (
	*http.Response,
	error) {
	req := m.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/msps/%v/licenses", mspId),
	)
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
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ListMspOrgLicenses takes context, mspId as parameters and
// returns an models.ApiResponse with models.License data and
// an error if there was an issue with the request or response.
// Get List of MSP Licences
func (m *MSPsLicenses) ListMspOrgLicenses(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[models.License],
	error) {
	req := m.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/msps/%v/stats/licenses", mspId),
	)
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
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.License
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.License](decoder)
	return models.NewApiResponse(result, resp), err
}