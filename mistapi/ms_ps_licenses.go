// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"net/http"
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

// ClaimMspLicense takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.ResponseClaimLicense data and
// an error if there was an issue with the request or response.
// Claim a license order for this MSP by submitting the activation code from the order.
func (m *MSPsLicenses) ClaimMspLicense(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.CodeString) (
	models.ApiResponse[models.ResponseClaimLicense],
	error) {
	req := m.prepareRequest(ctx, "POST", "/api/v1/msps/%v/claim")
	req.AppendTemplateParams(mspId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Response when the key is invalid (or already used)"},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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
// Return MSP license entitlement, subscription, amendment, and usage summary information.
func (m *MSPsLicenses) ListMspLicenses(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[models.License],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/licenses")
	req.AppendTemplateParams(mspId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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
// returns an *Response and
// an error if there was an issue with the request or response.
// Perform an MSP license action, such as amending license quantity to an organization, deleting a subscription, undoing an amendment, or annotating a subscription.
func (m *MSPsLicenses) MoveOrDeleteMspLicenseToAnotherOrg(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.MspLicenseAction) (
	*http.Response,
	error) {
	req := m.prepareRequest(ctx, "PUT", "/api/v1/msps/%v/licenses")
	req.AppendTemplateParams(mspId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ListMspOrgLicenses takes context, mspId as parameters and
// returns an models.ApiResponse with models.License data and
// an error if there was an issue with the request or response.
// Return license entitlement, subscription, amendment, and usage statistics for organizations managed by this MSP.
func (m *MSPsLicenses) ListMspOrgLicenses(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[models.License],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/stats/licenses")
	req.AppendTemplateParams(mspId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	var result models.License
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.License](decoder)
	return models.NewApiResponse(result, resp), err
}
