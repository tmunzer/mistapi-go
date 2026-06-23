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
)

// OrgsReports represents a controller struct.
type OrgsReports struct {
	baseController
}

// NewOrgsReports creates a new instance of OrgsReports.
// It takes a baseController as a parameter and returns a pointer to the OrgsReports.
func NewOrgsReports(baseController baseController) *OrgsReports {
	orgsReports := OrgsReports{baseController: baseController}
	return &orgsReports
}

// DisableOrgE911Report takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgE911Report data and
// an error if there was an issue with the request or response.
// Disable automatic E911 AP BSSID report generation for the organization.
func (o *OrgsReports) DisableOrgE911Report(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.OrgE911Report],
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/exports/e911_report")
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

	var result models.OrgE911Report
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgE911Report](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgE911Report takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgE911Report data and
// an error if there was an issue with the request or response.
// Get the status of E911 AP BSSID reports and download URL if available.
func (o *OrgsReports) GetOrgE911Report(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.OrgE911Report],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/exports/e911_report")
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

	var result models.OrgE911Report
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgE911Report](decoder)
	return models.NewApiResponse(result, resp), err
}

// EnableOrgE911Report takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgE911Report data and
// an error if there was an issue with the request or response.
// Enable automatic E911 AP BSSID report generation for the organization. Reports will be generated immediately and then every 24 hours.
func (o *OrgsReports) EnableOrgE911Report(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.OrgE911Report],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/exports/e911_report")
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

	var result models.OrgE911Report
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgE911Report](decoder)
	return models.NewApiResponse(result, resp), err
}
