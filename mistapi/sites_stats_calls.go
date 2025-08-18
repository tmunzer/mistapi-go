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

// SitesStatsCalls represents a controller struct.
type SitesStatsCalls struct {
	baseController
}

// NewSitesStatsCalls creates a new instance of SitesStatsCalls.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsCalls.
func NewSitesStatsCalls(baseController baseController) *SitesStatsCalls {
	sitesStatsCalls := SitesStatsCalls{baseController: baseController}
	return &sitesStatsCalls
}

// TroubleshootSiteCall takes context, siteId, clientMac, meetingId, mac, app, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.CallTroubleshoot data and
// an error if there was an issue with the request or response.
// Troubleshoot a call
func (s *SitesStatsCalls) TroubleshootSiteCall(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string,
	meetingId string,
	mac *string,
	app *string,
	start *int,
	end *int,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.CallTroubleshoot],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		"/api/v1/sites/%v/stats/calls/client/%v/troubleshoot",
	)
	req.AppendTemplateParams(siteId, clientMac)
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
	req.QueryParam("meeting_id", meetingId)
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if app != nil {
		req.QueryParam("app", *app)
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.CallTroubleshoot
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CallTroubleshoot](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteCalls takes context, siteId, distinct, rating, app, start, end, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Calls
func (s *SitesStatsCalls) CountSiteCalls(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.CountSiteCallsDistinctEnum,
	rating *int,
	app *string,
	start *int,
	end *int,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/calls/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if rating != nil {
		req.QueryParam("rating", *rating)
	}
	if app != nil {
		req.QueryParam("app", *app)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteCalls takes context, siteId, mac, app, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseStatsCalls data and
// an error if there was an issue with the request or response.
// Search Calls
func (s *SitesStatsCalls) SearchSiteCalls(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	app *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseStatsCalls],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/calls/search")
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
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if app != nil {
		req.QueryParam("app", *app)
	}
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

	var result models.ResponseStatsCalls
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseStatsCalls](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteCallsSummary takes context, siteId, apMac, app, start, end as parameters and
// returns an models.ApiResponse with models.ResponseStatsCallsSummary data and
// an error if there was an issue with the request or response.
// Summarized, aggregated stats for the site calls
func (s *SitesStatsCalls) GetSiteCallsSummary(
	ctx context.Context,
	siteId uuid.UUID,
	apMac *string,
	app *string,
	start *int,
	end *int) (
	models.ApiResponse[models.ResponseStatsCallsSummary],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/calls/summary")
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
	if apMac != nil {
		req.QueryParam("ap_mac", *apMac)
	}
	if app != nil {
		req.QueryParam("app", *app)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}

	var result models.ResponseStatsCallsSummary
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseStatsCallsSummary](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteTroubleshootCalls takes context, siteId, ap, meetingId, mac, app, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseCallTroubleshootSummary data and
// an error if there was an issue with the request or response.
// Summary of calls troubleshoot by site
func (s *SitesStatsCalls) ListSiteTroubleshootCalls(
	ctx context.Context,
	siteId uuid.UUID,
	ap *string,
	meetingId *string,
	mac *string,
	app *string,
	start *int,
	end *int,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseCallTroubleshootSummary],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/calls/troubleshoot")
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
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if meetingId != nil {
		req.QueryParam("meeting_id", *meetingId)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if app != nil {
		req.QueryParam("app", *app)
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseCallTroubleshootSummary
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCallTroubleshootSummary](decoder)
	return models.NewApiResponse(result, resp), err
}
