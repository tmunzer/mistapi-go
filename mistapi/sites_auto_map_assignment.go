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

// SitesAutoMapAssignment represents a controller struct.
type SitesAutoMapAssignment struct {
	baseController
}

// NewSitesAutoMapAssignment creates a new instance of SitesAutoMapAssignment.
// It takes a baseController as a parameter and returns a pointer to the SitesAutoMapAssignment.
func NewSitesAutoMapAssignment(baseController baseController) *SitesAutoMapAssignment {
	sitesAutoMapAssignment := SitesAutoMapAssignment{baseController: baseController}
	return &sitesAutoMapAssignment
}

// ApplySiteAutoMapAssignment takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseAutoMapAssignmentApply data and
// an error if there was an issue with the request or response.
// Apply (accept) auto map assignment results for a site. Devices are associated with their assigned maps. Omit `map_ids` or provide an empty list to accept all pending assignments; provide specific `map_ids` for a partial accept.
func (s *SitesAutoMapAssignment) ApplySiteAutoMapAssignment(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.AutoMapAssignmentRequest) (
	models.ApiResponse[models.ResponseAutoMapAssignmentApply],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/apply_auto_map_assignment",
	)
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

	var result models.ResponseAutoMapAssignmentApply
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoMapAssignmentApply](decoder)
	return models.NewApiResponse(result, resp), err
}

// CancelSiteAutoMapAssignment takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Cancel an in-progress auto map assignment operation for the site. Validates that auto map assignment is currently running, notifies all APs to fetch new configuration, and sends a cancel command to the orchestration service.
func (s *SitesAutoMapAssignment) CancelSiteAutoMapAssignment(
	ctx context.Context,
	siteId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/auto_map_assignment")
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
		"400": {Message: "Auto map assignment not in progress"},
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

// GetSiteAutoMapAssignmentStatus takes context, siteId as parameters and
// returns an models.ApiResponse with models.ResponseAutoMapAssignmentInfo data and
// an error if there was an issue with the request or response.
// Get the current status of auto map assignment for the site.
func (s *SitesAutoMapAssignment) GetSiteAutoMapAssignmentStatus(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.ResponseAutoMapAssignmentInfo],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/auto_map_assignment")
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

	var result models.ResponseAutoMapAssignmentInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoMapAssignmentInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// StartSiteAutoMapAssignment takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseAutoMapAssignment data and
// an error if there was an issue with the request or response.
// Start the auto map assignment process for a site. The service automatically assigns APs to maps based on BLE ranging data and requires at least 3 APs with compatible firmware and model support for BLE.
// Repeated POST requests while a site assignment is still running will be rejected.
func (s *SitesAutoMapAssignment) StartSiteAutoMapAssignment(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.AutoMapAssignment) (
	models.ApiResponse[models.ResponseAutoMapAssignment],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/auto_map_assignment")
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
		"400": {Message: "Auto map assignment already in progress"},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ResponseAutoMapAssignment
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoMapAssignment](decoder)
	return models.NewApiResponse(result, resp), err
}

// ClearSiteAutoMapAssignment takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseAutoMapAssignmentClear data and
// an error if there was an issue with the request or response.
// Clear (reject) auto map assignment results for a site without applying them. The cached assignment results are cleared. Omit `map_ids` or provide an empty list to reject all pending assignments; provide specific `map_ids` for a partial reject.
func (s *SitesAutoMapAssignment) ClearSiteAutoMapAssignment(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.AutoMapAssignmentRequest) (
	models.ApiResponse[models.ResponseAutoMapAssignmentClear],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/clear_auto_map_assignment",
	)
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

	var result models.ResponseAutoMapAssignmentClear
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoMapAssignmentClear](decoder)
	return models.NewApiResponse(result, resp), err
}
