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

// MSPsSSORoles represents a controller struct.
type MSPsSSORoles struct {
	baseController
}

// NewMSPsSSORoles creates a new instance of MSPsSSORoles.
// It takes a baseController as a parameter and returns a pointer to the MSPsSSORoles.
func NewMSPsSSORoles(baseController baseController) *MSPsSSORoles {
	mSPsSSORoles := MSPsSSORoles{baseController: baseController}
	return &mSPsSSORoles
}

// ListMspSsoRoles takes context, mspId as parameters and
// returns an models.ApiResponse with []models.SsoRoleMsp data and
// an error if there was an issue with the request or response.
// List MSP SSO role definitions that map identity-provider role assertions to MSP privilege scopes.
func (m *MSPsSSORoles) ListMspSsoRoles(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[[]models.SsoRoleMsp],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/ssoroles")
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

	var result []models.SsoRoleMsp
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.SsoRoleMsp](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateMspSsoRole takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.SsoRoleMsp data and
// an error if there was an issue with the request or response.
// Create an MSP SSO role definition with a display name and the MSP privileges granted when the role is matched during SSO.
func (m *MSPsSSORoles) CreateMspSsoRole(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.SsoRoleMsp) (
	models.ApiResponse[models.SsoRoleMsp],
	error) {
	req := m.prepareRequest(ctx, "POST", "/api/v1/msps/%v/ssoroles")
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

	var result models.SsoRoleMsp
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SsoRoleMsp](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteMspSsoRole takes context, mspId, ssoroleId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an MSP SSO role definition so it can no longer grant MSP privileges during SSO.
func (m *MSPsSSORoles) DeleteMspSsoRole(
	ctx context.Context,
	mspId uuid.UUID,
	ssoroleId uuid.UUID) (
	*http.Response,
	error) {
	req := m.prepareRequest(ctx, "DELETE", "/api/v1/msps/%v/ssoroles/%v")
	req.AppendTemplateParams(mspId, ssoroleId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// UpdateMspSsoRole takes context, mspId, ssoroleId, body as parameters and
// returns an models.ApiResponse with models.SsoRoleMsp data and
// an error if there was an issue with the request or response.
// Update an MSP SSO role definition, including its display name and granted MSP privileges.
func (m *MSPsSSORoles) UpdateMspSsoRole(
	ctx context.Context,
	mspId uuid.UUID,
	ssoroleId uuid.UUID,
	body *models.SsoRoleMsp) (
	models.ApiResponse[models.SsoRoleMsp],
	error) {
	req := m.prepareRequest(ctx, "PUT", "/api/v1/msps/%v/ssoroles/%v")
	req.AppendTemplateParams(mspId, ssoroleId)
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

	var result models.SsoRoleMsp
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SsoRoleMsp](decoder)
	return models.NewApiResponse(result, resp), err
}
