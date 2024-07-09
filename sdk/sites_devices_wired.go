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

// SitesDevicesWired represents a controller struct.
type SitesDevicesWired struct {
	baseController
}

// NewSitesDevicesWired creates a new instance of SitesDevicesWired.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesWired.
func NewSitesDevicesWired(baseController baseController) *SitesDevicesWired {
	sitesDevicesWired := SitesDevicesWired{baseController: baseController}
	return &sitesDevicesWired
}

// DeleteSiteLocalSwitchPortConfig takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Sometimes HelpDesk Admin needs to change port configs
func (s *SitesDevicesWired) DeleteSiteLocalSwitchPortConfig(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v/local_port_config", siteId, deviceId),
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

// UpdateSiteLocalSwitchPortConfig takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Sometimes HelpDesk Admin needs to change port configs
func (s *SitesDevicesWired) UpdateSiteLocalSwitchPortConfig(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	body *models.JunosPortConfig) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v/local_port_config", siteId, deviceId),
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

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// GetSiteSwitchesMetrics takes context, siteId, mType, scope, switchMac as parameters and
// returns an models.ApiResponse with models.ResponseSwitchMetrics data and
// an error if there was an issue with the request or response.
// Get version compliance metrics for managed or monitored switches
func (s *SitesDevicesWired) GetSiteSwitchesMetrics(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.SwitchMetricTypeEnum,
	scope *models.SwitchMetricScopeEnum,
	switchMac *string) (
	models.ApiResponse[models.ResponseSwitchMetrics],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/switches/metrics", siteId),
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if scope != nil {
		req.QueryParam("scope", *scope)
	}
	if switchMac != nil {
		req.QueryParam("switch_mac", *switchMac)
	}

	var result models.ResponseSwitchMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSwitchMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}
