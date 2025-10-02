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

// SitesStatsMxEdges represents a controller struct.
type SitesStatsMxEdges struct {
	baseController
}

// NewSitesStatsMxEdges creates a new instance of SitesStatsMxEdges.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsMxEdges.
func NewSitesStatsMxEdges(baseController baseController) *SitesStatsMxEdges {
	sitesStatsMxEdges := SitesStatsMxEdges{baseController: baseController}
	return &sitesStatsMxEdges
}

// ListSiteMxEdgesStats takes context, siteId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsMxedge data and
// an error if there was an issue with the request or response.
// Get List of Site MxEdges Stats
func (s *SitesStatsMxEdges) ListSiteMxEdgesStats(
	ctx context.Context,
	siteId uuid.UUID,
	start *string,
	end *string,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.StatsMxedge],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/mxedges")
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

	var result []models.StatsMxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsMxedge](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteMxEdgeStats takes context, siteId, mxedgeId, start, end, duration as parameters and
// returns an models.ApiResponse with models.StatsMxedge data and
// an error if there was an issue with the request or response.
// Get One Site MxEdge Stats
func (s *SitesStatsMxEdges) GetSiteMxEdgeStats(
	ctx context.Context,
	siteId uuid.UUID,
	mxedgeId uuid.UUID,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.StatsMxedge],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/mxedges/%v")
	req.AppendTemplateParams(siteId, mxedgeId)
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
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result models.StatsMxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsMxedge](decoder)
	return models.NewApiResponse(result, resp), err
}
