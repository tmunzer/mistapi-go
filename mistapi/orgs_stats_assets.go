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

// OrgsStatsAssets represents a controller struct.
type OrgsStatsAssets struct {
	baseController
}

// NewOrgsStatsAssets creates a new instance of OrgsStatsAssets.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsAssets.
func NewOrgsStatsAssets(baseController baseController) *OrgsStatsAssets {
	orgsStatsAssets := OrgsStatsAssets{baseController: baseController}
	return &orgsStatsAssets
}

// ListOrgAssetsStats takes context, orgId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsAsset data and
// an error if there was an issue with the request or response.
// Get List of Org Assets Stats
func (o *OrgsStatsAssets) ListOrgAssetsStats(
	ctx context.Context,
	orgId uuid.UUID,
	start *int,
	end *int,
	duration *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.StatsAsset],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/assets")
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

	var result []models.StatsAsset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsAsset](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgAssetsByDistanceField takes context, orgId, distinct, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Org Assets
func (o *OrgsStatsAssets) CountOrgAssetsByDistanceField(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgAssetCountDistinctEnum,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/assets/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
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

// SearchOrgAssets takes context, orgId, siteId, mac, deviceName, name, mapId, ibeaconUuid, ibeaconMajor, ibeaconMinor, eddystoneUidNamespace, eddystoneUidInstance, eddystoneUrl, apMac, beam, rssi, limit, start, end, duration, sort as parameters and
// returns an models.ApiResponse with models.ResponseStatsAssets data and
// an error if there was an issue with the request or response.
// Search for Org Assets
func (o *OrgsStatsAssets) SearchOrgAssets(
	ctx context.Context,
	orgId uuid.UUID,
	siteId *string,
	mac *string,
	deviceName *string,
	name *string,
	mapId *string,
	ibeaconUuid *string,
	ibeaconMajor *string,
	ibeaconMinor *string,
	eddystoneUidNamespace *string,
	eddystoneUidInstance *string,
	eddystoneUrl *string,
	apMac *string,
	beam *int,
	rssi *int,
	limit *int,
	start *int,
	end *int,
	duration *string,
	sort *string) (
	models.ApiResponse[models.ResponseStatsAssets],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/assets/search")
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
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if deviceName != nil {
		req.QueryParam("device_name", *deviceName)
	}
	if name != nil {
		req.QueryParam("name", *name)
	}
	if mapId != nil {
		req.QueryParam("map_id", *mapId)
	}
	if ibeaconUuid != nil {
		req.QueryParam("ibeacon_uuid", *ibeaconUuid)
	}
	if ibeaconMajor != nil {
		req.QueryParam("ibeacon_major", *ibeaconMajor)
	}
	if ibeaconMinor != nil {
		req.QueryParam("ibeacon_minor", *ibeaconMinor)
	}
	if eddystoneUidNamespace != nil {
		req.QueryParam("eddystone_uid_namespace", *eddystoneUidNamespace)
	}
	if eddystoneUidInstance != nil {
		req.QueryParam("eddystone_uid_instance", *eddystoneUidInstance)
	}
	if eddystoneUrl != nil {
		req.QueryParam("eddystone_url", *eddystoneUrl)
	}
	if apMac != nil {
		req.QueryParam("ap_mac", *apMac)
	}
	if beam != nil {
		req.QueryParam("beam", *beam)
	}
	if rssi != nil {
		req.QueryParam("rssi", *rssi)
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
	if sort != nil {
		req.QueryParam("sort", *sort)
	}

	var result models.ResponseStatsAssets
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseStatsAssets](decoder)
	return models.NewApiResponse(result, resp), err
}
