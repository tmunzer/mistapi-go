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

// SitesSpectrumAnalysis represents a controller struct.
type SitesSpectrumAnalysis struct {
	baseController
}

// NewSitesSpectrumAnalysis creates a new instance of SitesSpectrumAnalysis.
// It takes a baseController as a parameter and returns a pointer to the SitesSpectrumAnalysis.
func NewSitesSpectrumAnalysis(baseController baseController) *SitesSpectrumAnalysis {
	sitesSpectrumAnalysis := SitesSpectrumAnalysis{baseController: baseController}
	return &sitesSpectrumAnalysis
}

// GetSiteRunningSpectrumAnalysis takes context, siteId as parameters and
// returns an models.ApiResponse with models.ResponseRunningSpectrumAnalysis data and
// an error if there was an issue with the request or response.
// Get the running spectrum analysis for a site
func (s *SitesSpectrumAnalysis) GetSiteRunningSpectrumAnalysis(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.ResponseRunningSpectrumAnalysis],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/analyze_spectrum")
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

	var result models.ResponseRunningSpectrumAnalysis
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseRunningSpectrumAnalysis](decoder)
	return models.NewApiResponse(result, resp), err
}

// InitiateSiteAnalyzeSpectrum takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Initiate a spectrum analysis for a site
// The output will be available through websocket. As there can be multiple command
// issued against the same device at the same time and the output all goes through
// the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json { "subscribe": "/sites/{site_id}/analyze_spectrum" } ```
// #### Example output from ws stream
// ```json
// {
// "event": "data",
// "channel": "/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/analyze_spectrum",
// "data": {
// "session": "session_id",
// "fft_samples": [
// {
// "frequency": 2437.0,
// "rssi / signal ?": -93
// },
// ...
// ],
// "channel_usage": [
// {
// "channel": 36,
// "noise": -78,
// "wifi": 0.13,
// "non_wifi": 0.08
// },
// ...
// ]
// }
// }
// ```
func (s *SitesSpectrumAnalysis) InitiateSiteAnalyzeSpectrum(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.SpectrumAnalysis) (
	models.ApiResponse[models.WebsocketSession],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/analyze_spectrum")
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

	var result models.WebsocketSession
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteSpectrumAnalysis takes context, siteId, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponsePastSpectrumAnalysis data and
// an error if there was an issue with the request or response.
// List the past spectrum analysis for a site
func (s *SitesSpectrumAnalysis) ListSiteSpectrumAnalysis(
	ctx context.Context,
	siteId uuid.UUID,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponsePastSpectrumAnalysis],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/analyze_spectrum")
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result models.ResponsePastSpectrumAnalysis
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponsePastSpectrumAnalysis](decoder)
	return models.NewApiResponse(result, resp), err
}
