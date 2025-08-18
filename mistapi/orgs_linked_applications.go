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

// OrgsLinkedApplications represents a controller struct.
type OrgsLinkedApplications struct {
	baseController
}

// NewOrgsLinkedApplications creates a new instance of OrgsLinkedApplications.
// It takes a baseController as a parameter and returns a pointer to the OrgsLinkedApplications.
func NewOrgsLinkedApplications(baseController baseController) *OrgsLinkedApplications {
	orgsLinkedApplications := OrgsLinkedApplications{baseController: baseController}
	return &orgsLinkedApplications
}

// GetOrgOauthAppLinkedStatus takes context, orgId, appName, forward as parameters and
// returns an models.ApiResponse with models.AccountOauthInfo data and
// an error if there was an issue with the request or response.
// Get Org Level OAuth Application Linked Status
func (o *OrgsLinkedApplications) GetOrgOauthAppLinkedStatus(
	ctx context.Context,
	orgId uuid.UUID,
	appName models.OauthAppNameEnum,
	forward string) (
	models.ApiResponse[models.AccountOauthInfo],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/%v/link_accounts")
	req.AppendTemplateParams(orgId, appName)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.QueryParam("forward", forward)

	var result models.AccountOauthInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountOauthInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// AddOrgOauthAppAccounts takes context, orgId, appName, body as parameters and
// returns an models.ApiResponse with models.AccountOauthInfo data and
// an error if there was an issue with the request or response.
// Add Jamf, VMware Authorization With Mist Portal
func (o *OrgsLinkedApplications) AddOrgOauthAppAccounts(
	ctx context.Context,
	orgId uuid.UUID,
	appName models.OauthAppNameEnum,
	body *models.AccountOauthAdd) (
	models.ApiResponse[models.AccountOauthInfo],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/%v/link_accounts")
	req.AppendTemplateParams(orgId, appName)
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
		"400": {Message: "Unsuccessful"},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.AccountOauthInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountOauthInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgOauthAppAccount takes context, orgId, appName, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Update Zoom, Teams, Intune Authorization.
// Request Payload, These Field And Values Will Be Specific To Each Of The Third Party Apps Accounts.
func (o *OrgsLinkedApplications) UpdateOrgOauthAppAccount(
	ctx context.Context,
	orgId uuid.UUID,
	appName models.OauthAppNameEnum,
	body *models.AccountOauthConfig) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/setting/%v/link_accounts")
	req.AppendTemplateParams(orgId, appName)
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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

// DeleteOrgOauthAppAuthorization takes context, orgId, appName, accountId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Level OAuth Application Authorization With Mist Portal
func (o *OrgsLinkedApplications) DeleteOrgOauthAppAuthorization(
	ctx context.Context,
	orgId uuid.UUID,
	appName models.OauthAppNameEnum,
	accountId string) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		"/api/v1/orgs/%v/setting/%v/link_accounts/%v",
	)
	req.AppendTemplateParams(orgId, appName, accountId)
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
		"400": {Message: "Unsuccessful"},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
