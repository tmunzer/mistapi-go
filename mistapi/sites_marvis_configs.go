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

// SitesMarvisConfigs represents a controller struct.
type SitesMarvisConfigs struct {
	baseController
}

// NewSitesMarvisConfigs creates a new instance of SitesMarvisConfigs.
// It takes a baseController as a parameter and returns a pointer to the SitesMarvisConfigs.
func NewSitesMarvisConfigs(baseController baseController) *SitesMarvisConfigs {
	sitesMarvisConfigs := SitesMarvisConfigs{baseController: baseController}
	return &sitesMarvisConfigs
}

// CountSiteMarvisConfigActions takes context, siteId, distinct, mac, mType, src, adminId, op, portId, vlanIds, reason, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Marvis Config Actions for a site by a distinct field.
func (s *SitesMarvisConfigs) CountSiteMarvisConfigActions(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *string,
	mac *string,
	mType *string,
	src *string,
	adminId *string,
	op *string,
	portId *string,
	vlanIds *int,
	reason *string,
	limit *int,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/marvis_configs/count")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if src != nil {
		req.QueryParam("src", *src)
	}
	if adminId != nil {
		req.QueryParam("admin_id", *adminId)
	}
	if op != nil {
		req.QueryParam("op", *op)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if vlanIds != nil {
		req.QueryParam("vlan_ids", *vlanIds)
	}
	if reason != nil {
		req.QueryParam("reason", *reason)
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

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteMarvisConfigActions takes context, siteId, mac, mType, src, adminId, op, portId, vlanIds, reason, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.MarvisConfigActionsSearch data and
// an error if there was an issue with the request or response.
// Search Marvis Config Actions for a site.
func (s *SitesMarvisConfigs) SearchSiteMarvisConfigActions(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	mType *string,
	src *string,
	adminId *string,
	op *string,
	portId *string,
	vlanIds *int,
	reason *string,
	limit *int,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.MarvisConfigActionsSearch],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/marvis_configs/search")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if src != nil {
		req.QueryParam("src", *src)
	}
	if adminId != nil {
		req.QueryParam("admin_id", *adminId)
	}
	if op != nil {
		req.QueryParam("op", *op)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if vlanIds != nil {
		req.QueryParam("vlan_ids", *vlanIds)
	}
	if reason != nil {
		req.QueryParam("reason", *reason)
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

	var result models.MarvisConfigActionsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MarvisConfigActionsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteMarvisConfigAction takes context, siteId, id as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete a Marvis Config Action.
func (s *SitesMarvisConfigs) DeleteSiteMarvisConfigAction(
	ctx context.Context,
	siteId uuid.UUID,
	id uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/marvis_configs/%v")
	req.AppendTemplateParams(siteId, id)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// SubmitSiteMarvisConfigFeedback takes context, siteId, id, body as parameters and
// returns an models.ApiResponse with models.MarvisConfigFeedbackResponse data and
// an error if there was an issue with the request or response.
// Submit feedback on a Marvis-injected config action (e.g. mark as invalid).
func (s *SitesMarvisConfigs) SubmitSiteMarvisConfigFeedback(
	ctx context.Context,
	siteId uuid.UUID,
	id uuid.UUID,
	body models.MarvisConfigFeedback) (
	models.ApiResponse[models.MarvisConfigFeedbackResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/marvis_configs/%v/feedback",
	)
	req.AppendTemplateParams(siteId, id)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	req.Header("Content-Type", "application/json")
	req.Json(&body)

	var result models.MarvisConfigFeedbackResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MarvisConfigFeedbackResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
