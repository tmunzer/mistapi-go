package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// SamplesWebhooks represents a controller struct.
type SamplesWebhooks struct {
    baseController
}

// NewSamplesWebhooks creates a new instance of SamplesWebhooks.
// It takes a baseController as a parameter and returns a pointer to the SamplesWebhooks.
func NewSamplesWebhooks(baseController baseController) *SamplesWebhooks {
    samplesWebhooks := SamplesWebhooks{baseController: baseController}
    return &samplesWebhooks
}

// Alarms takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `alarm` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) Alarms(
    ctx context.Context,
    body *models.WebhookAlarms) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_alarm_")
    
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

// AssetRaw takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `asset_raw` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
// **will be deprecated after 06/30/2024**
func (s *SamplesWebhooks) AssetRaw(
    ctx context.Context,
    body *models.WebhookAssetRaw) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_asset_raw_")
    
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

// Audits takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `audit` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) Audits(
    ctx context.Context,
    body *models.WebhookAudits) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_audit_")
    
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

// ClientInfo takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `client-info` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) ClientInfo(
    ctx context.Context,
    body *models.WebhookClientInfo) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_client_info_")
    
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

// ClientJoin takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `client_join` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) ClientJoin(
    ctx context.Context,
    body *models.WebhookClientJoin) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_client_join_")
    
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

// ClientLatency takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `client-latency` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) ClientLatency(
    ctx context.Context,
    body *models.WebhookClientLatency) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_client_latency_")
    
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

// ClientSessions takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `client_sessions` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) ClientSessions(
    ctx context.Context,
    body *models.WebhookClientSessions) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_client_sessions_")
    
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

// DeviceEvents takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `device_events` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) DeviceEvents(
    ctx context.Context,
    body *models.WebhookDeviceEvents) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_device_events_")
    
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

// DeviceUpDown takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `device_updowns` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) DeviceUpDown(
    ctx context.Context,
    body *models.WebhookDeviceUpdowns) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_device_updowns_")
    
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

// DiscoveredRawRssi takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `discovered-raw-rssi` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) DiscoveredRawRssi(
    ctx context.Context,
    body *models.WebhookDiscoveredRawRssi) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_discovered_raw_rssi_")
    
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

// Location takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `location` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) Location(
    ctx context.Context,
    body *models.WebhookLocation) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_location_")
    
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

// LocationAsset takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `location_asset` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) LocationAsset(
    ctx context.Context,
    body *models.WebhookLocationAsset) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_location_asset_")
    
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

// LocationCentrak takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `location_centrak` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) LocationCentrak(
    ctx context.Context,
    body *models.WebhookLocationCentrak) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_location_centrak_")
    
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

// LocationClient takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `location_client` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) LocationClient(
    ctx context.Context,
    body *models.WebhookLocationClient) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_location_client_")
    
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

// LocationSdk takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `location_sdk` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) LocationSdk(
    ctx context.Context,
    body *models.WebhookLocationSdk) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_location_sdk_")
    
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

// LocationUnclient takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `location_unclient` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) LocationUnclient(
    ctx context.Context,
    body *models.WebhookLocationUnclient) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_location_unclient_")
    
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

// NacAccounting takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `nac-accounting` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) NacAccounting(
    ctx context.Context,
    body *models.WebhookNacAccounting) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_nac_accounting_")
    
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

// NacEvents takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Example Delivery of nac_events
func (s *SamplesWebhooks) NacEvents(
    ctx context.Context,
    body *models.WebhookNacEvents) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_nac_events_")
    
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

// OccupancyAlerts takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `occupancy_alerts` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) OccupancyAlerts(
    ctx context.Context,
    body *models.WebhookOccupancyAlerts) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_occupancy_alerts_")
    
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

// Ping takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `ping` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) Ping(
    ctx context.Context,
    body *models.WebhookPing) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_ping_")
    
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

// SdkclientScanData takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `sdkclient_scan_data` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) SdkclientScanData(
    ctx context.Context,
    body *models.WebhookSdkclientScanData) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_sdkclient_scan_data")
    
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

// SiteSle takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `site_sle` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) SiteSle(
    ctx context.Context,
    body *models.WebhookSiteSle) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_site_sle_")
    
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

// Zone takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Webhook sample for `zone` topic
// **Note**: The server host will be your own server FQDN where the Mist Cloud is sending the webhook messages
func (s *SamplesWebhooks) Zone(
    ctx context.Context,
    body *models.WebhookZone) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/webhook_example/_zone_")
    
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
