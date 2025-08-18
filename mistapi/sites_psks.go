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

// SitesPsks represents a controller struct.
type SitesPsks struct {
	baseController
}

// NewSitesPsks creates a new instance of SitesPsks.
// It takes a baseController as a parameter and returns a pointer to the SitesPsks.
func NewSitesPsks(baseController baseController) *SitesPsks {
	sitesPsks := SitesPsks{baseController: baseController}
	return &sitesPsks
}

// ListSitePsks takes context, siteId, ssid, role, name, limit, page as parameters and
// returns an models.ApiResponse with []models.Psk data and
// an error if there was an issue with the request or response.
// Get List of Site PSKs
func (s *SitesPsks) ListSitePsks(
	ctx context.Context,
	siteId uuid.UUID,
	ssid *string,
	role *string,
	name *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Psk],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/psks")
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
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if role != nil {
		req.QueryParam("role", *role)
	}
	if name != nil {
		req.QueryParam("name", *name)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSitePsk takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Psk data and
// an error if there was an issue with the request or response.
// Create Site PSK
// When `usage`==`macs`, corresponding "macs" field will hold a list consisting of client mac addresses (["xx:xx:xx:xx:xx",...]) or mac patterns(["xx:xx:*","xx*",...]) or both (["xx:xx:xx:xx:xx:xx", "xx:*", ...]). This list is capped at 5000
func (s *SitesPsks) CreateSitePsk(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.Psk) (
	models.ApiResponse[models.Psk],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/psks")
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

	var result models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteMultiplePsks takes context, siteId, body as parameters and
// returns an models.ApiResponse with []models.Psk data and
// an error if there was an issue with the request or response.
// Update multiple PSKs
func (s *SitesPsks) UpdateSiteMultiplePsks(
	ctx context.Context,
	siteId uuid.UUID,
	body []models.Psk) (
	models.ApiResponse[[]models.Psk],
	error) {
	req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/psks")
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

	var result []models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// ImportSitePsks takes context, siteId, file as parameters and
// returns an models.ApiResponse with []models.Psk data and
// an error if there was an issue with the request or response.
// Import PSK from CSV file or JSON
// ## CSV File Format
// ```csv
// PSK Import CSV File Format:
// name,ssid,passphrase,usage,vlan_id,mac
// Common,warehouse,foryoureyesonly,single,35,a31425f31278
// Justin,reception,visible,multi,1002
// ```
func (s *SitesPsks) ImportSitePsks(
	ctx context.Context,
	siteId uuid.UUID,
	file *models.FileWrapper) (
	models.ApiResponse[[]models.Psk],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/psks/import")
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
	formFields := []https.FormParam{}
	if file != nil {
		fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
		formFields = append(formFields, fileParam)
	}
	req.FormData(formFields)

	var result []models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSitePsk takes context, siteId, pskId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site PSK
func (s *SitesPsks) DeleteSitePsk(
	ctx context.Context,
	siteId uuid.UUID,
	pskId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/psks/%v")
	req.AppendTemplateParams(siteId, pskId)
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

// GetSitePsk takes context, siteId, pskId as parameters and
// returns an models.ApiResponse with models.Psk data and
// an error if there was an issue with the request or response.
// Get Site PSK Details
func (s *SitesPsks) GetSitePsk(
	ctx context.Context,
	siteId uuid.UUID,
	pskId uuid.UUID) (
	models.ApiResponse[models.Psk],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/psks/%v")
	req.AppendTemplateParams(siteId, pskId)
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

	var result models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSitePsk takes context, siteId, pskId, body as parameters and
// returns an models.ApiResponse with models.Psk data and
// an error if there was an issue with the request or response.
// Update Site PSK
func (s *SitesPsks) UpdateSitePsk(
	ctx context.Context,
	siteId uuid.UUID,
	pskId uuid.UUID,
	body *models.Psk) (
	models.ApiResponse[models.Psk],
	error) {
	req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/psks/%v")
	req.AppendTemplateParams(siteId, pskId)
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

	var result models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}
