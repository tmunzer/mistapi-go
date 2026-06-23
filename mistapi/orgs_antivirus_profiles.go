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

// OrgsAntivirusProfiles represents a controller struct.
type OrgsAntivirusProfiles struct {
	baseController
}

// NewOrgsAntivirusProfiles creates a new instance of OrgsAntivirusProfiles.
// It takes a baseController as a parameter and returns a pointer to the OrgsAntivirusProfiles.
func NewOrgsAntivirusProfiles(baseController baseController) *OrgsAntivirusProfiles {
	orgsAntivirusProfiles := OrgsAntivirusProfiles{baseController: baseController}
	return &orgsAntivirusProfiles
}

// ListOrgAntivirusProfiles takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Avprofile data and
// an error if there was an issue with the request or response.
// List organization antivirus scanning profiles, including inspected protocols, scan limits, whitelist settings, and fallback actions.
func (o *OrgsAntivirusProfiles) ListOrgAntivirusProfiles(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Avprofile],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/avprofiles")
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Avprofile
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Avprofile](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgAntivirusProfile takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Avprofile data and
// an error if there was an issue with the request or response.
// Create an organization antivirus scanning profile with inspected protocols, maximum file size, whitelist settings, and fallback action.
func (o *OrgsAntivirusProfiles) CreateOrgAntivirusProfile(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Avprofile) (
	models.ApiResponse[models.Avprofile],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/avprofiles")
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

	var result models.Avprofile
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Avprofile](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgAntivirusProfile takes context, orgId, avprofileId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization antivirus scanning profile by profile ID.
func (o *OrgsAntivirusProfiles) DeleteOrgAntivirusProfile(
	ctx context.Context,
	orgId uuid.UUID,
	avprofileId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/avprofiles/%v")
	req.AppendTemplateParams(orgId, avprofileId)
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

// GetOrgAntivirusProfile takes context, orgId, avprofileId as parameters and
// returns an models.ApiResponse with models.Avprofile data and
// an error if there was an issue with the request or response.
// Return one organization antivirus scanning profile, including inspected protocols, scan limits, whitelist settings, and fallback action.
func (o *OrgsAntivirusProfiles) GetOrgAntivirusProfile(
	ctx context.Context,
	orgId uuid.UUID,
	avprofileId uuid.UUID) (
	models.ApiResponse[models.Avprofile],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/avprofiles/%v")
	req.AppendTemplateParams(orgId, avprofileId)
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

	var result models.Avprofile
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Avprofile](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgAntivirusProfile takes context, orgId, avprofileId, body as parameters and
// returns an models.ApiResponse with models.Avprofile data and
// an error if there was an issue with the request or response.
// Update an organization antivirus scanning profile's inspected protocols, maximum file size, whitelist settings, or fallback action.
func (o *OrgsAntivirusProfiles) UpdateOrgAntivirusProfile(
	ctx context.Context,
	orgId uuid.UUID,
	avprofileId uuid.UUID,
	body *models.Avprofile) (
	models.ApiResponse[models.Avprofile],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/avprofiles/%v")
	req.AppendTemplateParams(orgId, avprofileId)
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

	var result models.Avprofile
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Avprofile](decoder)
	return models.NewApiResponse(result, resp), err
}
