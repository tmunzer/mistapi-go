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

// OrgsIntegrationJSE represents a controller struct.
type OrgsIntegrationJSE struct {
	baseController
}

// NewOrgsIntegrationJSE creates a new instance of OrgsIntegrationJSE.
// It takes a baseController as a parameter and returns a pointer to the OrgsIntegrationJSE.
func NewOrgsIntegrationJSE(baseController baseController) *OrgsIntegrationJSE {
	orgsIntegrationJSE := OrgsIntegrationJSE{baseController: baseController}
	return &orgsIntegrationJSE
}

// GetOrgJseInfo takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Return the JSE organizations associated with the configured account. Use the returned organization names when selecting JSE provider options for secure edge tunnels.
func (o *OrgsIntegrationJSE) GetOrgJseInfo(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.AccountJseInfo],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/jse/info")
	req.AppendTemplateParams(orgId)
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

	var result models.AccountJseInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgJseIntegration takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Remove the JSE integration configuration from the organization.
func (o *OrgsIntegrationJSE) DeleteOrgJseIntegration(
	ctx context.Context,
	orgId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/setting/jse/setup")
	req.AppendTemplateParams(orgId)
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

// GetOrgJseIntegration takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Return the JSE integration configuration, including the cloud hostname, integration username, and associated JSE organization names.
func (o *OrgsIntegrationJSE) GetOrgJseIntegration(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.AccountJseInfo],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/jse/setup")
	req.AppendTemplateParams(orgId)
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

	var result models.AccountJseInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// SetupOrgJseIntegration takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Configure the JSE integration with the JSE cloud hostname and integration-user credentials. In JSE, use a custom role with read access to `service_location` and read-write access to site and IPsec profile APIs, then create and activate the integration user and service locations.
func (o *OrgsIntegrationJSE) SetupOrgJseIntegration(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.AccountJseConfig) (
	models.ApiResponse[models.AccountJseInfo],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/jse/setup")
	req.AppendTemplateParams(orgId)
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

	var result models.AccountJseInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
	return models.NewApiResponse(result, resp), err
}
