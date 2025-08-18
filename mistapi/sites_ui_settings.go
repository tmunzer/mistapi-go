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

// SitesUISettings represents a controller struct.
type SitesUISettings struct {
	baseController
}

// NewSitesUISettings creates a new instance of SitesUISettings.
// It takes a baseController as a parameter and returns a pointer to the SitesUISettings.
func NewSitesUISettings(baseController baseController) *SitesUISettings {
	sitesUISettings := SitesUISettings{baseController: baseController}
	return &sitesUISettings
}

// ListSiteUiSettings takes context, siteId as parameters and
// returns an models.ApiResponse with []models.UiSettings data and
// an error if there was an issue with the request or response.
// List the Site UI settings/databoard
func (s *SitesUISettings) ListSiteUiSettings(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[[]models.UiSettings],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/uisettings")
	req.AppendTemplateParams(siteId)
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

	var result []models.UiSettings
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UiSettings](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSiteUiSettings takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.UiSettings data and
// an error if there was an issue with the request or response.
// Create a Site UI settings/databoard
func (s *SitesUISettings) CreateSiteUiSettings(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UiSettings) (
	models.ApiResponse[models.UiSettings],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/uisettings")
	req.AppendTemplateParams(siteId)
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

	var result models.UiSettings
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UiSettings](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteUiSettingDerived takes context, siteId as parameters and
// returns an models.ApiResponse with models.UiSettings data and
// an error if there was an issue with the request or response.
// Get both site UI settings(for_site=true) and org UI settings (for_site=false)
func (s *SitesUISettings) ListSiteUiSettingDerived(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.UiSettings],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/uisettings/derived")
	req.AppendTemplateParams(siteId)
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

	var result models.UiSettings
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UiSettings](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteUiSetting takes context, siteId, uisettingId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Site UI settings
func (s *SitesUISettings) DeleteSiteUiSetting(
	ctx context.Context,
	siteId uuid.UUID,
	uisettingId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/uisettings/%v")
	req.AppendTemplateParams(siteId, uisettingId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetSiteUiSetting takes context, siteId, uisettingId as parameters and
// returns an models.ApiResponse with models.UiSettings data and
// an error if there was an issue with the request or response.
// Site UI settings
func (s *SitesUISettings) GetSiteUiSetting(
	ctx context.Context,
	siteId uuid.UUID,
	uisettingId uuid.UUID) (
	models.ApiResponse[models.UiSettings],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/uisettings/%v")
	req.AppendTemplateParams(siteId, uisettingId)
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

	var result models.UiSettings
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UiSettings](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteUiSetting takes context, siteId, uisettingId, body as parameters and
// returns an models.ApiResponse with models.UiSettings data and
// an error if there was an issue with the request or response.
// Site UI settings
func (s *SitesUISettings) UpdateSiteUiSetting(
	ctx context.Context,
	siteId uuid.UUID,
	uisettingId uuid.UUID,
	body *models.UiSettings) (
	models.ApiResponse[models.UiSettings],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/uisettings/%v")
	req.AppendTemplateParams(siteId, uisettingId)
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

	var result models.UiSettings
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UiSettings](decoder)
	return models.NewApiResponse(result, resp), err
}
