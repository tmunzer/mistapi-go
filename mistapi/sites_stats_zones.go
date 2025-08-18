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

// SitesStatsZones represents a controller struct.
type SitesStatsZones struct {
	baseController
}

// NewSitesStatsZones creates a new instance of SitesStatsZones.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsZones.
func NewSitesStatsZones(baseController baseController) *SitesStatsZones {
	sitesStatsZones := SitesStatsZones{baseController: baseController}
	return &sitesStatsZones
}

// ListSiteRssiZonesStats takes context, siteId as parameters and
// returns an models.ApiResponse with []models.StatsRssiZone data and
// an error if there was an issue with the request or response.
// Get List of Site RSSI Zones Stats
func (s *SitesStatsZones) ListSiteRssiZonesStats(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[[]models.StatsRssiZone],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/rssizones")
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

	var result []models.StatsRssiZone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsRssiZone](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteRssiZoneStats takes context, siteId, zoneId as parameters and
// returns an models.ApiResponse with models.StatsZoneDetails data and
// an error if there was an issue with the request or response.
// Get Detail RSSI Zone Stats
func (s *SitesStatsZones) GetSiteRssiZoneStats(
	ctx context.Context,
	siteId uuid.UUID,
	zoneId uuid.UUID) (
	models.ApiResponse[models.StatsZoneDetails],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/rssizones/%v")
	req.AppendTemplateParams(siteId, zoneId)
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

	var result models.StatsZoneDetails
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsZoneDetails](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteZonesStats takes context, siteId, mapId, minDuration as parameters and
// returns an models.ApiResponse with []models.StatsZone data and
// an error if there was an issue with the request or response.
// Get List of Site Zones Stats
func (s *SitesStatsZones) ListSiteZonesStats(
	ctx context.Context,
	siteId uuid.UUID,
	mapId *string,
	minDuration *int) (
	models.ApiResponse[[]models.StatsZone],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/zones")
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
	if mapId != nil {
		req.QueryParam("map_id", *mapId)
	}
	if minDuration != nil {
		req.QueryParam("min_duration", *minDuration)
	}

	var result []models.StatsZone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsZone](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteZoneStats takes context, siteId, zoneId as parameters and
// returns an models.ApiResponse with models.StatsZoneDetails data and
// an error if there was an issue with the request or response.
// Get Detail Zone Stats
func (s *SitesStatsZones) GetSiteZoneStats(
	ctx context.Context,
	siteId uuid.UUID,
	zoneId uuid.UUID) (
	models.ApiResponse[models.StatsZoneDetails],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/zones/%v")
	req.AppendTemplateParams(siteId, zoneId)
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

	var result models.StatsZoneDetails
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsZoneDetails](decoder)
	return models.NewApiResponse(result, resp), err
}
