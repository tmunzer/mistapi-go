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

// OrgsAlarmTemplates represents a controller struct.
type OrgsAlarmTemplates struct {
	baseController
}

// NewOrgsAlarmTemplates creates a new instance of OrgsAlarmTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsAlarmTemplates.
func NewOrgsAlarmTemplates(baseController baseController) *OrgsAlarmTemplates {
	orgsAlarmTemplates := OrgsAlarmTemplates{baseController: baseController}
	return &orgsAlarmTemplates
}

// ListOrgAlarmTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// List alarm templates defined for the organization, including default delivery settings and per-alarm rule configuration.
func (o *OrgsAlarmTemplates) ListOrgAlarmTemplates(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.AlarmTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/alarmtemplates")
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

	var result []models.AlarmTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.AlarmTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgAlarmTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// Create an organization alarm template that defines default delivery settings and per-alarm rule overrides.
// Available rules can be found in [List Alarm Definitions]($e/Events%20Definitions/listAlarmDefinitions)
// The `delivery` object is only required when it differs from the template delivery settings.
// To assign an Alarm template to a site, use the [Update Site]($e/Sites/updateSiteInfo) endpoint and specify the Alarm template ID in the `alarmtemplate_id` field of the request body.
func (o *OrgsAlarmTemplates) CreateOrgAlarmTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.AlarmTemplate) (
	models.ApiResponse[models.AlarmTemplate],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/alarmtemplates")
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

	var result models.AlarmTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AlarmTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// UnsuppressOrgSuppressedAlarms takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Remove alarm suppression entries currently configured for this organization.
func (o *OrgsAlarmTemplates) UnsuppressOrgSuppressedAlarms(
	ctx context.Context,
	orgId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/alarmtemplates/suppress")
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

// ListOrgSuppressedAlarms takes context, orgId, scope as parameters and
// returns an models.ApiResponse with models.ResponseOrgSuppressAlarm data and
// an error if there was an issue with the request or response.
// List alarm suppression entries currently configured for this organization, optionally filtered by organization-wide or site-specific scope.
func (o *OrgsAlarmTemplates) ListOrgSuppressedAlarms(
	ctx context.Context,
	orgId uuid.UUID,
	scope *models.SuppressedAlarmScopeEnum) (
	models.ApiResponse[models.ResponseOrgSuppressAlarm],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/alarmtemplates/suppress")
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
	if scope != nil {
		req.QueryParam("scope", *scope)
	}

	var result models.ResponseOrgSuppressAlarm
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseOrgSuppressAlarm](decoder)
	return models.NewApiResponse(result, resp), err
}

// SuppressOrgAlarm takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Create or schedule an alarm suppression window for the organization or selected sites, for example during planned maintenance.
func (o *OrgsAlarmTemplates) SuppressOrgAlarm(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.SuppressedAlarm) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/alarmtemplates/suppress")
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// DeleteOrgAlarmTemplate takes context, orgId, alarmtemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization alarm template by template ID from this organization.
func (o *OrgsAlarmTemplates) DeleteOrgAlarmTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	alarmtemplateId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/alarmtemplates/%v")
	req.AppendTemplateParams(orgId, alarmtemplateId)
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

// GetOrgAlarmTemplate takes context, orgId, alarmtemplateId as parameters and
// returns an models.ApiResponse with models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// Return one organization alarm template, including default delivery settings and per-alarm rule overrides.
func (o *OrgsAlarmTemplates) GetOrgAlarmTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	alarmtemplateId uuid.UUID) (
	models.ApiResponse[models.AlarmTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/alarmtemplates/%v")
	req.AppendTemplateParams(orgId, alarmtemplateId)
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

	var result models.AlarmTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AlarmTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgAlarmTemplate takes context, orgId, alarmtemplateId, body as parameters and
// returns an models.ApiResponse with models.AlarmTemplate data and
// an error if there was an issue with the request or response.
// Update an organization alarm template's default delivery settings or per-alarm rule overrides.
func (o *OrgsAlarmTemplates) UpdateOrgAlarmTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	alarmtemplateId uuid.UUID,
	body *models.AlarmTemplate) (
	models.ApiResponse[models.AlarmTemplate],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/alarmtemplates/%v")
	req.AppendTemplateParams(orgId, alarmtemplateId)
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

	var result models.AlarmTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AlarmTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}
