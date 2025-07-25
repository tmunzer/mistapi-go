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
// Ping test from the AP to confirm ‘reachability’ of the Radius server. Utilize Juniper EX switch(to which an AP is connected to) radius test capabilities to get details on the Radius Server ‘availability’ .
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
      "/api/v1/sites/%v/devices/%v/check_radius_server",
    )
    req.AppendTemplateParams(siteId, deviceId)
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
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceSyntheticTest takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.SynthetictestInfo data and
// an error if there was an issue with the request or response.
// Get Device Synthetic Test
func (s *SitesSyntheticTests) GetSiteDeviceSyntheticTest(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.SynthetictestInfo],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/%v/synthetic_test")
    req.AppendTemplateParams(siteId, deviceId)
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
        "400": {Message: "Device not online / Device not supported / Already in progress"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    var result models.SynthetictestInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SynthetictestInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// TriggerSiteDeviceSyntheticTest takes context, siteId, deviceId, body as parameters and
// returns an *Response and
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
      "/api/v1/sites/%v/devices/%v/synthetic_test",
    )
    req.AppendTemplateParams(siteId, deviceId)
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
        "400": {Message: "Device not online / Device not supported / Already in progress"},
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

// TriggerSiteSyntheticTest takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseSynthetictest data and
// an error if there was an issue with the request or response.
// Trigger Synthetic Testing
func (s *SitesSyntheticTests) TriggerSiteSyntheticTest(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Synthetictest) (
    models.ApiResponse[models.ResponseSynthetictest],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/synthetic_test")
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
    
    var result models.ResponseSynthetictest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSynthetictest](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteSyntheticTest takes context, siteId, mac, portId, vlanId, by, reason, mType, protocol, tenant as parameters and
// returns an models.ApiResponse with models.ResponseSynthetictestSearch data and
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
    mType *models.SynthetictestTypeEnum,
    protocol *models.SynthetictestProtocolEnum,
    tenant *string) (
    models.ApiResponse[models.ResponseSynthetictestSearch],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/synthetic_test/search")
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
    if protocol != nil {
        req.QueryParam("protocol", *protocol)
    }
    if tenant != nil {
        req.QueryParam("tenant", *tenant)
    }
    
    var result models.ResponseSynthetictestSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSynthetictestSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
