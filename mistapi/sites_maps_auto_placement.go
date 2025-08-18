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

// SitesMapsAutoPlacement represents a controller struct.
type SitesMapsAutoPlacement struct {
	baseController
}

// NewSitesMapsAutoPlacement creates a new instance of SitesMapsAutoPlacement.
// It takes a baseController as a parameter and returns a pointer to the SitesMapsAutoPlacement.
func NewSitesMapsAutoPlacement(baseController baseController) *SitesMapsAutoPlacement {
	sitesMapsAutoPlacement := SitesMapsAutoPlacement{baseController: baseController}
	return &sitesMapsAutoPlacement
}

// DeleteSiteApAutoOrientation takes context, mapId, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API is called to force stop auto placement for a given map
func (s *SitesMapsAutoPlacement) DeleteSiteApAutoOrientation(
	ctx context.Context,
	mapId uuid.UUID,
	siteId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/maps/%v/auto_orient")
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
		"400": {Message: "Autoplacement was not triggered"},
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

// GetSiteApAutoOrientation takes context, mapId, siteId as parameters and
// returns an models.ApiResponse with models.ResponseAutoOrientationInfo data and
// an error if there was an issue with the request or response.
// This API is called to view the current status of auto orient for a given map.
func (s *SitesMapsAutoPlacement) GetSiteApAutoOrientation(
	ctx context.Context,
	mapId uuid.UUID,
	siteId uuid.UUID) (
	models.ApiResponse[models.ResponseAutoOrientationInfo],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/maps/%v/auto_orient")
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
		"400": {Message: "Autoplacement was not triggered"},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	var result models.ResponseAutoOrientationInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoOrientationInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// StartSiteApAutoOrientation takes context, mapId, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseAutoOrientation data and
// an error if there was an issue with the request or response.
// This API is called to trigger a map for auto orient. For auto orient feature to work, BLE data needs to be collected from the APs on the map. This precess is not disruptive unlike FTM collection. Repeated POST requests to this endpoint while a map is still running will be rejected.
// `force_collection` is set to `false` by default. If `force_collection`==`false`, the API attempts to start orientation with existing data. If no data exists, the API attempts to start collecting orientation data. If `force_collection`==`true`, the API attempts to start collecting orientation data.
// Providing a list of device macs is optional. If provided, auto orientation suggestions will be made only for the specified devices. If no list is provided, all APs associated with the map are considered by default.
func (s *SitesMapsAutoPlacement) StartSiteApAutoOrientation(
	ctx context.Context,
	mapId uuid.UUID,
	siteId uuid.UUID,
	body *models.AutoOrient) (
	models.ApiResponse[models.ResponseAutoOrientation],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/maps/%v/auto_orient")
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
		"400": {Message: "Bad Request", Unmarshaller: errors.NewResponseDetailString},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ResponseAutoOrientation
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoOrientation](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteApAutoplacement takes context, siteId, mapId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API is called to force stop auto placement for a given map
func (s *SitesMapsAutoPlacement) DeleteSiteApAutoplacement(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/maps/%v/auto_placement")
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
		"400": {Message: "Autoplacement was not triggered"},
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

// GetSiteApAutoPlacement takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with models.ResponseAutoPlacementInfo data and
// an error if there was an issue with the request or response.
// This API is called to view the current status of auto placement for a given map.
// #### Status Descriptions
// | Status | Description |
// | --- | --- |
// | `pending` | Autoplacement has not been requested for this map |
// | `inprogress` | Autoplacement is currently processing |
// | `done` | The autoplacement process has completed |
// | `data_needed` | Additional position data is required for autoplacement. Users should verify the requested anchor APs have a position on the map |
// | `invalid_model` | Autoplacement is not supported on the model of the APs on the map |
// | `invalid_version` | Autoplacement is not supported with the APs current firmware version |
// | `error` | There was an error in the autoplacement process |
func (s *SitesMapsAutoPlacement) GetSiteApAutoPlacement(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	models.ApiResponse[models.ResponseAutoPlacementInfo],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/maps/%v/auto_placement")
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

	var result models.ResponseAutoPlacementInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoPlacementInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// RunSiteApAutoplacement takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with models.ResponseAutoplacement data and
// an error if there was an issue with the request or response.
// This API is called to trigger auto placement for a map. For the auto placement feature to work, RTT-FTM data needs to be collected from the APs on the map.
// This scan is disruptive, and users must be notified of service disruption during the auto placement process. Repeated POST requests to this endpoint while a map is still running will be rejected.
// `force_collection` is set to `false` by default. If `force_collection` is set to `false`, the API attempts to start localization with existing data. If no data exists, the API attempts to start orchestration.
// If `force_collection` is set to `true`, the API attempts to start orchestration.
// Providing a list of devices is optional. If provided, autoplacement suggestions will be made only for the specified devices. If no list is provided, all APs associated with the map are considered by default.
func (s *SitesMapsAutoPlacement) RunSiteApAutoplacement(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.AutoPlacement) (
	models.ApiResponse[models.ResponseAutoplacement],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/maps/%v/auto_placement")
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

	var result models.ResponseAutoplacement
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoplacement](decoder)
	return models.NewApiResponse(result, resp), err
}

// ClearSiteApAutoOrient takes context, siteId, mapId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API is used to destroy the autoorientations of a map or subset of APs on a map.
func (s *SitesMapsAutoPlacement) ClearSiteApAutoOrient(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.MacAddresses) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/maps/%v/clear_auto_orient",
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

// ClearSiteApAutoplacement takes context, siteId, mapId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API is used to destroy the cached autoplacement locations of a map or subset of APs on a map.
func (s *SitesMapsAutoPlacement) ClearSiteApAutoplacement(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.MacAddresses) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/maps/%v/clear_autoplacement",
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

// ConfirmSiteApLocalizationData takes context, siteId, mapId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API is used to accept or reject the cached autoplacement and auto-orientation values of a map or subset of APs on a map. Any APs that have autoplacement values are stored in cache for up to 7 days while awaiting acceptance or rejection.
// ```
// Accepting the autoplacement values overwrites the existing X, Y, and orientation of the accepted APs with their cached autoplacement values.
// Rejecting the autoplacement values causes the APs to retain their current X, Y, and orientation.
// ```
// Once a decision (accept or reject) is made, or the 7-day time-to-live (TTL) expires, the cached values are deleted.
func (s *SitesMapsAutoPlacement) ConfirmSiteApLocalizationData(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.UseAutoApValues) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/maps/%v/use_auto_ap_values",
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
		"400": {Message: "Map does not exist or belong to specified site / Invalid localization service. Expected [placement, orientation]"},
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
