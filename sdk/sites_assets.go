package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// SitesAssets represents a controller struct.
type SitesAssets struct {
	baseController
}

// NewSitesAssets creates a new instance of SitesAssets.
// It takes a baseController as a parameter and returns a pointer to the SitesAssets.
func NewSitesAssets(baseController baseController) *SitesAssets {
	sitesAssets := SitesAssets{baseController: baseController}
	return &sitesAssets
}

// ListSiteAssets takes context, siteId, page, limit as parameters and
// returns an models.ApiResponse with []models.Asset data and
// an error if there was an issue with the request or response.
// Get List of Site Assets
func (s *SitesAssets) ListSiteAssets(
	ctx context.Context,
	siteId uuid.UUID,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.Asset],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/assets", siteId),
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

	var result []models.Asset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Asset](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSiteAsset takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Create Site Asset
func (s *SitesAssets) CreateSiteAsset(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.Asset) (
	models.ApiResponse[models.Asset],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/assets", siteId),
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Asset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Asset](decoder)
	return models.NewApiResponse(result, resp), err
}

// ImportSiteAssets takes context, siteId, upsert, file as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Impert Site Assets.
// It can be done via a CSV file or a JSON payload.
// ## CSV File Format
// ```csv
// name,mac
// "asset_name",5c5b53010101
// ```
func (s *SitesAssets) ImportSiteAssets(
	ctx context.Context,
	siteId uuid.UUID,
	upsert *models.ImportSiteAssetsUpsertEnum,
	file *models.FileWrapper) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/assets/import", siteId),
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
	req.Header("Content-Type", "multipart/form_data")
	if upsert != nil {
		req.QueryParam("upsert", *upsert)
	}
	formFields := []https.FormParam{}
	if file != nil {
		fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
		formFields = append(formFields, fileParam)
	}
	req.FormData(formFields)

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// DeleteSiteAsset takes context, siteId, assetId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Site Asset
func (s *SitesAssets) DeleteSiteAsset(
	ctx context.Context,
	siteId uuid.UUID,
	assetId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/sites/%v/assets/%v", siteId, assetId),
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

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// GetSiteAsset takes context, siteId, assetId as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Get Site Asset Details
func (s *SitesAssets) GetSiteAsset(
	ctx context.Context,
	siteId uuid.UUID,
	assetId uuid.UUID) (
	models.ApiResponse[models.Asset],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/assets/%v", siteId, assetId),
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

	var result models.Asset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Asset](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteAsset takes context, siteId, assetId, body as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Update Site Asset
func (s *SitesAssets) UpdateSiteAsset(
	ctx context.Context,
	siteId uuid.UUID,
	assetId uuid.UUID,
	body *models.Asset) (
	models.ApiResponse[models.Asset],
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/sites/%v/assets/%v", siteId, assetId),
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Asset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Asset](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteAssetsStats takes context, siteId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.StatsAsset data and
// an error if there was an issue with the request or response.
// Get List of Site Assets Stats
func (s *SitesAssets) ListSiteAssetsStats(
	ctx context.Context,
	siteId uuid.UUID,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[[]models.StatsAsset],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/assets", siteId),
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
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result []models.StatsAsset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsAsset](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteAssetStats takes context, siteId, start, end, duration as parameters and
// returns an models.ApiResponse with models.StatsAsset data and
// an error if there was an issue with the request or response.
// Get Site Asset Details
func (s *SitesAssets) GetSiteAssetStats(
	ctx context.Context,
	siteId uuid.UUID,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.StatsAsset],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/assets/asset_id", siteId),
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
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result models.StatsAsset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsAsset](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteAssets takes context, siteId, distinct as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Asset by distinct field
func (s *SitesAssets) CountSiteAssets(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteAssetsCountDistinctEnum) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/assets/count", siteId),
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteAssets takes context, siteId, mac, mapId, ibeaconUuid, ibeaconMajor, ibeaconMinor, eddystoneUidNamespace, eddystoneUidInstance, eddystoneUrl, deviceName, by, name, apMac, beam, rssi, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseStatsAssets data and
// an error if there was an issue with the request or response.
// Assets Search
func (s *SitesAssets) SearchSiteAssets(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	mapId *uuid.UUID,
	ibeaconUuid *uuid.UUID,
	ibeaconMajor *int,
	ibeaconMinor *int,
	eddystoneUidNamespace *string,
	eddystoneUidInstance *string,
	eddystoneUrl *string,
	deviceName *string,
	by *string,
	name *string,
	apMac *string,
	beam *string,
	rssi *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseStatsAssets],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/assets/search", siteId),
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
	if mac != nil {
		req.QueryParam("mac", *mac)
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
	if deviceName != nil {
		req.QueryParam("device_name", *deviceName)
	}
	if by != nil {
		req.QueryParam("by", *by)
	}
	if name != nil {
		req.QueryParam("name", *name)
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

	var result models.ResponseStatsAssets
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseStatsAssets](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteDiscoveredAssets takes context, siteId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.Asset data and
// an error if there was an issue with the request or response.
// Get List of Site Discovered BLE Assets that doesn’t match any of the Asset / Assetfilters
func (s *SitesAssets) ListSiteDiscoveredAssets(
	ctx context.Context,
	siteId uuid.UUID,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[[]models.Asset],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/discovered_assets", siteId),
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
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result []models.Asset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Asset](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteAssetsOfInterest takes context, siteId, duration, start, end, page, limit as parameters and
// returns an models.ApiResponse with []models.AssetOfInterest data and
// an error if there was an issue with the request or response.
// Get a list of BLE beacons that matches Asset or AssetFilter
func (s *SitesAssets) GetSiteAssetsOfInterest(
	ctx context.Context,
	siteId uuid.UUID,
	duration *string,
	start *int,
	end *int,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.AssetOfInterest],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/filtered_assets", siteId),
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
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result []models.AssetOfInterest
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.AssetOfInterest](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteDiscoveredAssetByMap takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with []models.StatsAsset data and
// an error if there was an issue with the request or response.
// Get a list of BLE beacons that we discovered (whether they’re defined as assets or not)
func (s *SitesAssets) GetSiteDiscoveredAssetByMap(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	models.ApiResponse[[]models.StatsAsset],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/maps/%v/discovered_assets", siteId, mapId),
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

	var result []models.StatsAsset
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsAsset](decoder)
	return models.NewApiResponse(result, resp), err
}
