package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// UtilitiesLAN represents a controller struct.
type UtilitiesLAN struct {
    baseController
}

// NewUtilitiesLAN creates a new instance of UtilitiesLAN.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesLAN.
func NewUtilitiesLAN(baseController baseController) *UtilitiesLAN {
    utilitiesLAN := UtilitiesLAN{baseController: baseController}
    return &utilitiesLAN
}

// UpgradeSiteDevicesBios takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Upgrade Bios on Multiple Device
func (u *UtilitiesLAN) UpgradeSiteDevicesBios(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UpgradeBiosMulti) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/upgrade_bios", siteId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// UpgradeSiteDevicesFpga takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Upgrade Bios on Multiple Device
func (u *UtilitiesLAN) UpgradeSiteDevicesFpga(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UpgradeFpgaMulti) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/upgrade_fpga", siteId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// PortsBounceFromSwitch takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Port Bounce can be performed from the Switch.The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```json
// {
// "event": "data",
// "channel": "/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd",
// "data": {
// "session": "session_id",
// "raw": "Port bounce complete."
// }
// }
// ```
func (u *UtilitiesLAN) PortsBounceFromSwitch(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsBouncePort) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/bounce_port", siteId, deviceId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// CableTestFromSwitch takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// TDR can be performed from the Switch. The output will be available through websocket. As there can be multiple command issued against the same Switch at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```json
// {
// "event": "data",
// "channel": "/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd",
// "data": {
// "session": "session_id",
// "raw": "Interface TDR detail:\nTest status : Test successfully executed  ge-0/0/0\n"
// }
// }
// ```
func (u *UtilitiesLAN) CableTestFromSwitch(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsCableTests) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/cable_test", siteId, deviceId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// ClearBpduErrosFromPortsOnSwitch takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Clear bridge protocol data unit (BPDU) error condition caused by the detection of a possible bridging loop from Spanning Tree Protocol (STP) operation that renders the port unoperational.
func (u *UtilitiesLAN) ClearBpduErrosFromPortsOnSwitch(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearBpdu) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/clear_bpdu_error", siteId, deviceId),
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
        "400": {Message: "Port not specified"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// ClearAllLearnedMacsFromPortOnSwitch takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Clear all learned MAC addresses, including persistent MAC addresses, on a port.
func (u *UtilitiesLAN) ClearAllLearnedMacsFromPortOnSwitch(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearMacs) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/clear_macs", siteId, deviceId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// PollSiteSwitchStats takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API can be used to poll statistics from the Switch proactively once. After it is called, the statistics will be pushed back to the cloud within the statistics interval.
func (u *UtilitiesLAN) PollSiteSwitchStats(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/poll_stats", siteId, deviceId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// CreateSiteDeviceSnapshot takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.ResponseDeviceSnapshot data and
// an error if there was an issue with the request or response.
// Create recovery device snapshot (Available on Junos OS EX2300-, EX3400-, EX4400- devices)
func (u *UtilitiesLAN) CreateSiteDeviceSnapshot(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.ResponseDeviceSnapshot],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/snapshot", siteId, deviceId),
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
        "400": {Message: "Bad Request"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    var result models.ResponseDeviceSnapshot
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceSnapshot](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeDeviceBios takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.ResponseDeviceBiosUpgrade data and
// an error if there was an issue with the request or response.
// Upgrade device bios
func (u *UtilitiesLAN) UpgradeDeviceBios(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UpgradeBios) (
    models.ApiResponse[models.ResponseDeviceBiosUpgrade],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/upgrade_bios", siteId, deviceId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseDeviceBiosUpgrade
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceBiosUpgrade](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeDeviceFPGA takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.ResponseDeviceBiosUpgrade data and
// an error if there was an issue with the request or response.
// Upgrade device fpga
func (u *UtilitiesLAN) UpgradeDeviceFPGA(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UpgradeFpga) (
    models.ApiResponse[models.ResponseDeviceBiosUpgrade],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/upgrade_fpga", siteId, deviceId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseDeviceBiosUpgrade
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceBiosUpgrade](decoder)
    return models.NewApiResponse(result, resp), err
}
