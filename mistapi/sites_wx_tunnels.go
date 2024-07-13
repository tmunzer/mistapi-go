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

// SitesWxTunnels represents a controller struct.
type SitesWxTunnels struct {
	baseController
}

// NewSitesWxTunnels creates a new instance of SitesWxTunnels.
// It takes a baseController as a parameter and returns a pointer to the SitesWxTunnels.
func NewSitesWxTunnels(baseController baseController) *SitesWxTunnels {
	sitesWxTunnels := SitesWxTunnels{baseController: baseController}
	return &sitesWxTunnels
}

// ListSiteWxTunnels takes context, siteId, page, limit as parameters and
// returns an models.ApiResponse with []models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Get List of Site WxLan Tunnels
func (s *SitesWxTunnels) ListSiteWxTunnels(
	ctx context.Context,
	siteId uuid.UUID,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.WxlanTunnel],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/wxtunnels", siteId),
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

	var result []models.WxlanTunnel
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.WxlanTunnel](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSiteWxTunnel takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Create Site WxLan Tunnel
func (s *SitesWxTunnels) CreateSiteWxTunnel(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.WxlanTunnel) (
	models.ApiResponse[models.WxlanTunnel],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/wxtunnels", siteId),
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

	var result models.WxlanTunnel
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WxlanTunnel](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteWxTunnel takes context, siteId, wxtunnelId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Site WxLan Tunnel
func (s *SitesWxTunnels) DeleteSiteWxTunnel(
	ctx context.Context,
	siteId uuid.UUID,
	wxtunnelId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/sites/%v/wxtunnels/%v", siteId, wxtunnelId),
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

// GetSiteWxTunnel takes context, siteId, wxtunnelId as parameters and
// returns an models.ApiResponse with models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Get Site WxLan tunnel Details
func (s *SitesWxTunnels) GetSiteWxTunnel(
	ctx context.Context,
	siteId uuid.UUID,
	wxtunnelId uuid.UUID) (
	models.ApiResponse[models.WxlanTunnel],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/wxtunnels/%v", siteId, wxtunnelId),
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

	var result models.WxlanTunnel
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WxlanTunnel](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteWxTunnel takes context, siteId, wxtunnelId, body as parameters and
// returns an models.ApiResponse with models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Update Site WxLan Tunnel
func (s *SitesWxTunnels) UpdateSiteWxTunnel(
	ctx context.Context,
	siteId uuid.UUID,
	wxtunnelId uuid.UUID,
	body *models.WxlanTunnel) (
	models.ApiResponse[models.WxlanTunnel],
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/sites/%v/wxtunnels/%v", siteId, wxtunnelId),
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

	var result models.WxlanTunnel
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WxlanTunnel](decoder)
	return models.NewApiResponse(result, resp), err
}
