package mistapigo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// SitesSyntheticTests represents a controller struct.
type SitesSyntheticTests struct {
	baseController
}

// NewSitesSyntheticTests creates a new instance of SitesSyntheticTests.
// It takes a baseController as a parameter and returns a pointer to the SitesSyntheticTests.
func NewSitesSyntheticTests(baseController baseController) *SitesSyntheticTests {
	sitesSyntheticTests := SitesSyntheticTests{baseController: baseController}
	return &sitesSyntheticTests
}

// StartSiteSwitchRadiusSyntheticTest takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Ping test from the AP to confirm ‘reachability’ of the Radius server. Utilize Juniper EX switch(to which an AP is connected to) radius test capabilities to get details on the Radius Server ‘availability’.
func (s *SitesSyntheticTests) StartSiteSwitchRadiusSyntheticTest(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	body *models.SynthetictestRadiusServer) (
	models.ApiResponse[models.WebsocketSession],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v/check_radius_server", siteId, deviceId),
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

	var result models.WebsocketSession
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceSyntheticTest takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Get Device Synthetic Test
func (s *SitesSyntheticTests) GetSiteDeviceSyntheticTest(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v/synthetic_test", siteId, deviceId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Device not online / Device not supported / Already in progress"},
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

// TriggerSiteDeviceSyntheticTest takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Trigger Device Synthetic Test
func (s *SitesSyntheticTests) TriggerSiteDeviceSyntheticTest(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	body *models.SynthetictestDevice) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v/synthetic_test", siteId, deviceId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Device not online / Device not supported / Already in progress"},
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

// GetSiteSyntheticTestStatus takes context, siteId as parameters and
// returns an models.ApiResponse with models.SynthetictestInfo data and
// an error if there was an issue with the request or response.
// Get Synthetic Testing Status
func (s *SitesSyntheticTests) GetSiteSyntheticTestStatus(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[models.SynthetictestInfo],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/synthetic_test", siteId),
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

	var result models.SynthetictestInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SynthetictestInfo](decoder)
	return models.NewApiResponse(result, resp), err
}

// TriggerSiteSyntheticTest takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ReponseSynthetictest data and
// an error if there was an issue with the request or response.
// Trigger Synthetic Testing
func (s *SitesSyntheticTests) TriggerSiteSyntheticTest(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.Synthetictest) (
	models.ApiResponse[models.ReponseSynthetictest],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/synthetic_test", siteId),
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

	var result models.ReponseSynthetictest
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReponseSynthetictest](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteSyntheticTest takes context, siteId, mac, portId, vlanId, by, reason, mType as parameters and
// returns an models.ApiResponse with models.ReponseSynthetictestSearch data and
// an error if there was an issue with the request or response.
// Search Site Synthetic Testing
func (s *SitesSyntheticTests) SearchSiteSyntheticTest(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	portId *string,
	vlanId *string,
	by *string,
	reason *string,
	mType *models.SynthetictestTypeEnum) (
	models.ApiResponse[models.ReponseSynthetictestSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/synthetic_test/search", siteId),
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
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if vlanId != nil {
		req.QueryParam("vlan_id", *vlanId)
	}
	if by != nil {
		req.QueryParam("by", *by)
	}
	if reason != nil {
		req.QueryParam("reason", *reason)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}

	var result models.ReponseSynthetictestSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReponseSynthetictestSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
