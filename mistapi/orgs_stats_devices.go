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

// OrgsStatsDevices represents a controller struct.
type OrgsStatsDevices struct {
	baseController
}

// NewOrgsStatsDevices creates a new instance of OrgsStatsDevices.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsDevices.
func NewOrgsStatsDevices(baseController baseController) *OrgsStatsDevices {
	orgsStatsDevices := OrgsStatsDevices{baseController: baseController}
	return &orgsStatsDevices
}

// ListOrgDevicesStats takes context, orgId, mType, status, siteId, mac, evpntopoId, evpnUnused, fields, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsDevice data and
// an error if there was an issue with the request or response.
// Get List of Org Devices stats
// This API renders some high-level device stats, pagination is assumed and returned in response header (as the response is an array)
func (o *OrgsStatsDevices) ListOrgDevicesStats(
	ctx context.Context,
	orgId uuid.UUID,
	mType *models.DeviceTypeWithAllEnum,
	status *models.DeviceStatusEnum,
	siteId *string,
	mac *string,
	evpntopoId *string,
	evpnUnused *string,
	fields *string,
	start *int,
	end *int,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.StatsDevice],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/devices")
	req.AppendTemplateParams(orgId)
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if evpntopoId != nil {
		req.QueryParam("evpntopo_id", *evpntopoId)
	}
	if evpnUnused != nil {
		req.QueryParam("evpn_unused", *evpnUnused)
	}
	if fields != nil {
		req.QueryParam("fields", *fields)
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

	var result []models.StatsDevice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsDevice](decoder)
	return models.NewApiResponse(result, resp), err
}
