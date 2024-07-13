package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
)

// OrgsAssets represents a controller struct.
type OrgsAssets struct {
	baseController
}

// NewOrgsAssets creates a new instance of OrgsAssets.
// It takes a baseController as a parameter and returns a pointer to the OrgsAssets.
func NewOrgsAssets(baseController baseController) *OrgsAssets {
	orgsAssets := OrgsAssets{baseController: baseController}
	return &orgsAssets
}

// ListOrgAssets takes context, orgId, page, limit as parameters and
// returns an models.ApiResponse with []models.Asset data and
// an error if there was an issue with the request or response.
// Get List of Org Assets
func (o *OrgsAssets) ListOrgAssets(
	ctx context.Context,
	orgId uuid.UUID,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.Asset],
	error) {
	req := o.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/orgs/%v/assets", orgId))
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

// CreateOrgAsset takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Create Org Asset
func (o *OrgsAssets) CreateOrgAsset(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Asset) (
	models.ApiResponse[models.Asset],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/assets", orgId),
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

// ImportOrgAssets takes context, orgId, file as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Impert Org Assets.
// It can be done via a CSV file or a JSON payload.
// #### CSV File Format
// ```csv
// name,mac
// "asset_name",5c5b53010101
// ```
func (o *OrgsAssets) ImportOrgAssets(
	ctx context.Context,
	orgId uuid.UUID,
	file *models.FileWrapper) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/assets/import", orgId),
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "multipart/form_data")
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

// DeleteOrgAsset takes context, orgId, assetId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org Asset
func (o *OrgsAssets) DeleteOrgAsset(
	ctx context.Context,
	orgId uuid.UUID,
	assetId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/orgs/%v/assets/%v", orgId, assetId),
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

// GetOrgAsset takes context, orgId, assetId as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Get Org Asset Details
func (o *OrgsAssets) GetOrgAsset(
	ctx context.Context,
	orgId uuid.UUID,
	assetId uuid.UUID) (
	models.ApiResponse[models.Asset],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/assets/%v", orgId, assetId),
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

// UpdateOrgAsset takes context, orgId, assetId, body as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Update Org Asset
func (o *OrgsAssets) UpdateOrgAsset(
	ctx context.Context,
	orgId uuid.UUID,
	assetId uuid.UUID,
	body *models.Asset) (
	models.ApiResponse[models.Asset],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/orgs/%v/assets/%v", orgId, assetId),
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

// ListOrgAssetsStats takes context, orgId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.StatsAsset data and
// an error if there was an issue with the request or response.
// Get List of Org Assets Stats
func (o *OrgsAssets) ListOrgAssetsStats(
	ctx context.Context,
	orgId uuid.UUID,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[[]models.StatsAsset],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/stats/assets", orgId),
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

// CountOrgAssetsByDistanceField takes context, orgId, distinct as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Org Assets
func (o *OrgsAssets) CountOrgAssetsByDistanceField(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgAssetCountDistinctEnum) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/stats/assets/count", orgId),
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

// SearchOrgAssets takes context, orgId, siteId, mac, deviceName, name, mapId, ibeaconUuid, ibeaconMajor, ibeaconMinor, eddystoneUidNamespace, eddystoneUidInstance, eddystoneUrl, apMac, beam, rssi, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseStatsAssets data and
// an error if there was an issue with the request or response.
// Search for Org Assets
func (o *OrgsAssets) SearchOrgAssets(
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
	duration *string) (
	models.ApiResponse[models.ResponseStatsAssets],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/stats/assets/search", orgId),
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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

	var result models.ResponseStatsAssets
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseStatsAssets](decoder)
	return models.NewApiResponse(result, resp), err
}
