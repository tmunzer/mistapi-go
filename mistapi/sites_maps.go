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
	"net/http"
)

// SitesMaps represents a controller struct.
type SitesMaps struct {
	baseController
}

// NewSitesMaps creates a new instance of SitesMaps.
// It takes a baseController as a parameter and returns a pointer to the SitesMaps.
func NewSitesMaps(baseController baseController) *SitesMaps {
	sitesMaps := SitesMaps{baseController: baseController}
	return &sitesMaps
}

// ListSiteMaps takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Map data and
// an error if there was an issue with the request or response.
// Get List of Site Maps
func (s *SitesMaps) ListSiteMaps(
	ctx context.Context,
	siteId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Map],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/maps")
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
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Map
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Map](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSiteMap takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Map data and
// an error if there was an issue with the request or response.
// Create Site Map
func (s *SitesMaps) CreateSiteMap(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.Map) (
	models.ApiResponse[models.Map],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/maps")
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

	var result models.Map
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Map](decoder)
	return models.NewApiResponse(result, resp), err
}

// ImportSiteMaps takes context, siteId, autoDeviceprofileAssignment, csv, file, json as parameters and
// returns an models.ApiResponse with models.ResponseMapImport data and
// an error if there was an issue with the request or response.
// Import data from files is a multipart POST which has an file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches.
// # Note
// This endpoint (at the site level), the AP must be already assigned to the site to be placed on the floorplan. If you want to place APs from the Org inventory, it is required to use the endpoint at the Org level [importOrgMaps](#operation/importOrgMaps)
// # CSV File Format
// ```csv
// Vendor AP name,Mist AP Mac
// US Office AP-2,5c:5b:35:00:00:02
// US Office AP-3,5c5b35000002
// ```
func (s *SitesMaps) ImportSiteMaps(
	ctx context.Context,
	siteId uuid.UUID,
	autoDeviceprofileAssignment *bool,
	csv *models.FileWrapper,
	file *models.FileWrapper,
	json *models.MapImportJson) (
	models.ApiResponse[models.ResponseMapImport],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/maps/import")
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
	formFields := []https.FormParam{}
	if autoDeviceprofileAssignment != nil {
		auto_deviceprofile_assignmentParam := https.FormParam{Key: "auto_deviceprofile_assignment", Value: *autoDeviceprofileAssignment, Headers: http.Header{}}
		formFields = append(formFields, auto_deviceprofile_assignmentParam)
	}
	if csv != nil {
		csvParam := https.FormParam{Key: "csv", Value: *csv, Headers: http.Header{}}
		formFields = append(formFields, csvParam)
	}
	if file != nil {
		fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
		formFields = append(formFields, fileParam)
	}
	if json != nil {
		jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
		formFields = append(formFields, jsonParam)
	}
	req.FormData(formFields)

	var result models.ResponseMapImport
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMapImport](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteMap takes context, siteId, mapId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Map
func (s *SitesMaps) DeleteSiteMap(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/maps/%v")
	req.AppendTemplateParams(siteId, mapId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetSiteMap takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with models.Map data and
// an error if there was an issue with the request or response.
// Get Site Map Details
func (s *SitesMaps) GetSiteMap(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	models.ApiResponse[models.Map],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/maps/%v")
	req.AppendTemplateParams(siteId, mapId)
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

	var result models.Map
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Map](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteMap takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with models.Map data and
// an error if there was an issue with the request or response.
// Update Site Map
func (s *SitesMaps) UpdateSiteMap(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.Map) (
	models.ApiResponse[models.Map],
	error) {
	req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/maps/%v")
	req.AppendTemplateParams(siteId, mapId)
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

	var result models.Map
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Map](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteMapImage takes context, siteId, mapId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Map Image
func (s *SitesMaps) DeleteSiteMapImage(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/maps/%v/image")
	req.AppendTemplateParams(siteId, mapId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// AddSiteMapImage takes context, siteId, mapId, file, json as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Add image map is a multipart POST which has an file (Image) and an optional json parameter
func (s *SitesMaps) AddSiteMapImage(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	file models.FileWrapper,
	json *string) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/maps/%v/image")
	req.AppendTemplateParams(siteId, mapId)
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
	formFields := []https.FormParam{}
	fileParam := https.FormParam{Key: "file", Value: file, Headers: http.Header{}}
	formFields = append(formFields, fileParam)
	if json != nil {
		jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
		formFields = append(formFields, jsonParam)
	}
	req.FormData(formFields)

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ReplaceSiteMapImage takes context, siteId, mapId, file, json as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Replace Map Image
// This works like an PUT where the image will be replaced. If transform is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)
func (s *SitesMaps) ReplaceSiteMapImage(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	file models.FileWrapper,
	json *models.MapSiteReplaceFileJson) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/maps/%v/replace")
	req.AppendTemplateParams(siteId, mapId)
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
	formFields := []https.FormParam{}
	fileParam := https.FormParam{Key: "file", Value: file, Headers: http.Header{}}
	formFields = append(formFields, fileParam)
	if json != nil {
		jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
		formFields = append(formFields, jsonParam)
	}
	req.FormData(formFields)

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// BulkAssignSiteApsToMap takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with models.ResponseSetDevicesMap data and
// an error if there was an issue with the request or response.
// This API can be used to assign a list of AP Macs associated with site_id to the specified map_id. Note that map_id must be associated with corresponding site_id. This API obeys the following rules
// 1. if AP is unassigned to any Map, it gets associated with map_id
// 2. Any moved APs are returned in the response
// 3. If the AP is considered a locked AP, no action will be taken
func (s *SitesMaps) BulkAssignSiteApsToMap(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.MacAddresses) (
	models.ApiResponse[models.ResponseSetDevicesMap],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/maps/%v/set_map")
	req.AppendTemplateParams(siteId, mapId)
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

	var result models.ResponseSetDevicesMap
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSetDevicesMap](decoder)
	return models.NewApiResponse(result, resp), err
}

// ImportSiteWayfindings takes context, siteId, mapId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This imports the vendor map meta data into the Map JSON. This is required by the SDK and App in order to access/render the vendor Map properly.
func (s *SitesMaps) ImportSiteWayfindings(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.WayfindingImportJson) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/maps/%v/wayfinding/import",
	)
	req.AppendTemplateParams(siteId, mapId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
