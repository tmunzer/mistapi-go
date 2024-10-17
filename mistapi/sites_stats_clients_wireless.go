package mistapi

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesStatsClientsWireless represents a controller struct.
type SitesStatsClientsWireless struct {
	baseController
}

// NewSitesStatsClientsWireless creates a new instance of SitesStatsClientsWireless.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsClientsWireless.
func NewSitesStatsClientsWireless(baseController baseController) *SitesStatsClientsWireless {
	sitesStatsClientsWireless := SitesStatsClientsWireless{baseController: baseController}
	return &sitesStatsClientsWireless
}

// ListSiteWirelessClientsStats takes context, siteId, wired, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.StatsClient data and
// an error if there was an issue with the request or response.
// Get List of Site All Clients Stats Details
func (s *SitesStatsClientsWireless) ListSiteWirelessClientsStats(
	ctx context.Context,
	siteId uuid.UUID,
	wired *bool,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[[]models.StatsClient],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/clients", siteId),
	)
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
	if wired != nil {
		req.QueryParam("wired", *wired)
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

	var result []models.StatsClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteWirelessClientStats takes context, siteId, clientMac, wired as parameters and
// returns an models.ApiResponse with []models.StatsClient data and
// an error if there was an issue with the request or response.
// Get Site Client Stats Details
func (s *SitesStatsClientsWireless) GetSiteWirelessClientStats(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string,
	wired *bool) (
	models.ApiResponse[[]models.StatsClient],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/clients/%v", siteId, clientMac),
	)
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
	if wired != nil {
		req.QueryParam("wired", *wired)
	}

	var result []models.StatsClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteWirelessClientsStatsByMap takes context, siteId, mapId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsWirelessClient data and
// an error if there was an issue with the request or response.
// Get Site Clients Stats By Map
func (s *SitesStatsClientsWireless) GetSiteWirelessClientsStatsByMap(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	start *int,
	end *int,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.StatsWirelessClient],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/maps/%v/clients", siteId, mapId),
	)
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

	var result []models.StatsWirelessClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsWirelessClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteUnconnectedClientStats takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with []models.StatsUnconnectedClient data and
// an error if there was an issue with the request or response.
// Get List of Site Unconnected Client Location
func (s *SitesStatsClientsWireless) ListSiteUnconnectedClientStats(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	models.ApiResponse[[]models.StatsUnconnectedClient],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/maps/%v/unconnected_clients", siteId, mapId),
	)
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

	var result []models.StatsUnconnectedClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsUnconnectedClient](decoder)
	return models.NewApiResponse(result, resp), err
}
