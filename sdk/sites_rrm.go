package mistapigo

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// SitesRRM represents a controller struct.
type SitesRRM struct {
	baseController
}

// NewSitesRRM creates a new instance of SitesRRM.
// It takes a baseController as a parameter and returns a pointer to the SitesRRM.
func NewSitesRRM(baseController baseController) *SitesRRM {
	sitesRRM := SitesRRM{baseController: baseController}
	return &sitesRRM
}

// GetSiteCurrentChannelPlanning takes context, siteId as parameters and
// returns an models.ApiResponse with models.Rrm data and
// an error if there was an issue with the request or response.
// Get Current Channel Planning
func (s *SitesRRM) GetSiteCurrentChannelPlanning(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.Rrm],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rrm/current", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.Rrm
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Rrm](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteCurrentRrmConsiderations takes context, siteId, deviceId, band as parameters and
// returns an models.ApiResponse with models.ResponseRrmConsideration data and
// an error if there was an issue with the request or response.
// Get Current RRM Considerations for an AP on a Specific Band
func (s *SitesRRM) GetSiteCurrentRrmConsiderations(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	band models.Dot11BandEnum) (
	models.ApiResponse[models.ResponseRrmConsideration],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rrm/current/devices/%v/band/%v", siteId, deviceId, band),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.ResponseRrmConsideration
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseRrmConsideration](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteRrmEvents takes context, siteId, band, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseEventsRrm data and
// an error if there was an issue with the request or response.
// Get Site RRM Events
func (s *SitesRRM) GetSiteRrmEvents(
	ctx context.Context,
	siteId uuid.UUID,
	band *models.Dot11BandEnum,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseEventsRrm],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rrm/events", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if band != nil {
		req.QueryParam("band", *band)
	}
	if page != nil {
		req.QueryParam("page", *page)
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

	var result models.ResponseEventsRrm
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsRrm](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteCurrentRrmNeighbors takes context, siteId, band, page, limit as parameters and
// returns an models.ApiResponse with models.ResponseRrmNeighbors data and
// an error if there was an issue with the request or response.
// Get Current RRM observed neighbors
func (s *SitesRRM) GetSiteCurrentRrmNeighbors(
	ctx context.Context,
	siteId uuid.UUID,
	band models.Dot11BandEnum,
	page *int,
	limit *int) (
	models.ApiResponse[models.ResponseRrmNeighbors],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rrm/neighbors/band/%v", siteId, band),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result models.ResponseRrmNeighbors
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseRrmNeighbors](decoder)
	return models.NewApiResponse(result, resp), err
}
