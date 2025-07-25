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

// SitesWebhooks represents a controller struct.
type SitesWebhooks struct {
    baseController
}

// NewSitesWebhooks creates a new instance of SitesWebhooks.
// It takes a baseController as a parameter and returns a pointer to the SitesWebhooks.
func NewSitesWebhooks(baseController baseController) *SitesWebhooks {
    sitesWebhooks := SitesWebhooks{baseController: baseController}
    return &sitesWebhooks
}

// ListSiteWebhooks takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Webhook data and
// an error if there was an issue with the request or response.
// Get List of Site Webhooks
func (s *SitesWebhooks) ListSiteWebhooks(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Webhook],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/webhooks")
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
    
    var result []models.Webhook
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Webhook](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteWebhook takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Webhook data and
// an error if there was an issue with the request or response.
// Webhook defines a webhook, modeled after [github\u2019s model](https://developer.github.com/webhooks/).
// There is two types of webhooks:
// * webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-be01e694-7253-4195-8563-78e2a745e114))        
// * raw data webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-e2d5d5f8-4bdb-4efc-93e4-90f4b33d0b2b))
// ##### Webhooks
// Webhooks can be configured at the org level (subset of topics only) and at the site level. It is possible to have multiple topics in the same webhook configuration and/or to have multiple webhooks configured at the same time.
// ##### Client Raw Data Webhooks
// Raw data webhooks are a special subset of webhooks that provide insight into raw data packets emitted by a client, identified by their advertising MAC address (assets, discovered BLE, connected Wi-Fi, unconnected Wi-Fi). The data that client raw data webhooks encompasses are reporting AP information, RSSI Data, and any special packets/telemetry packets that the client may emit. Note that client raw webhooks are the raw data coming from the client and do not contain the X,Y location data of the client. In order to get the location data for a client please see our location webhooks. Clients can be identified uniquely across these client raw data topics and location webhook topic using MAC address as the Unique identifier (client identifier).
// ###### Client Raw Data Webhooks Topics
// Topics that correspond to client raw data for different client types. 
// * `asset-raw-rssi` - Raw data from packets emitted by named and filtered assets 
// * `discovered-raw-rssi` - Raw data from packets emitted by passive BLE devices 
// * `wifi-conn-raw` - Raw data from packets emitted by connected devices 
// * `wifi-unconn-raw` - Raw data from packets emitted by unconnected devices (passive)
// ### Asset Filtering for Client Raw Data Webhooks
// The `asset-raw-rssi` webhook topic supports filtering of raw data by incorporating asset filters in the webhook payload.  
// The filter topic allows multiple Webhooks to receive a subset of the a`asset-raw-rssi` data by assigning asset filters to a given webhook. The `asset-raw-rssi` filter topic is filtered-asset-rssi.
// A webhook assigned to a filter topic can take a list of AssetFilter IDs, which act as inclusive filters to determine which named asset and filtered asset data is sent to the assigned filter topic. Filters can be applied to multiple webhooks, and the same data can be sent to multiple filter topics.
// ### Rules for Configuring Client Raw Data Webhooks
// 1. Only four instances of a webhook object can contain a specific filter topic. - A site-level entry will override an org-level entry for the same client raw data webhook topic.
// 2. An assigned asset filter must exist and belong to the same site as the webhook it is assigned to.
func (s *SitesWebhooks) CreateSiteWebhook(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Webhook) (
    models.ApiResponse[models.Webhook],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/webhooks")
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
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400Webhook},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.Webhook
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Webhook](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWebhook takes context, siteId, webhookId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Webhook
func (s *SitesWebhooks) DeleteSiteWebhook(
    ctx context.Context,
    siteId uuid.UUID,
    webhookId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/webhooks/%v")
    req.AppendTemplateParams(siteId, webhookId)
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

// GetSiteWebhook takes context, siteId, webhookId as parameters and
// returns an models.ApiResponse with models.Webhook data and
// an error if there was an issue with the request or response.
// Get Site Webhook Details
func (s *SitesWebhooks) GetSiteWebhook(
    ctx context.Context,
    siteId uuid.UUID,
    webhookId uuid.UUID) (
    models.ApiResponse[models.Webhook],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/webhooks/%v")
    req.AppendTemplateParams(siteId, webhookId)
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
    
    var result models.Webhook
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Webhook](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteWebhook takes context, siteId, webhookId, body as parameters and
// returns an models.ApiResponse with models.Webhook data and
// an error if there was an issue with the request or response.
// Update Site Webhook
func (s *SitesWebhooks) UpdateSiteWebhook(
    ctx context.Context,
    siteId uuid.UUID,
    webhookId uuid.UUID,
    body *models.Webhook) (
    models.ApiResponse[models.Webhook],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/webhooks/%v")
    req.AppendTemplateParams(siteId, webhookId)
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
    
    var result models.Webhook
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Webhook](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountSiteWebhooksDeliveries takes context, siteId, webhookId, mError, statusCode, status, topic, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Site Webhooks deliveries
// Topics Supported:
// - alarms
// - audits
// - device-updowns
// - occupancy-alerts
// - ping
func (s *SitesWebhooks) CountSiteWebhooksDeliveries(
    ctx context.Context,
    siteId uuid.UUID,
    webhookId uuid.UUID,
    mError *string,
    statusCode *int,
    status *models.WebhookDeliveryStatusEnum,
    topic *models.WebhookDeliveryTopicEnum,
    distinct *models.WebhookDeliveryDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/webhooks/%v/events/count")
    req.AppendTemplateParams(siteId, webhookId)
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
    if mError != nil {
        req.QueryParam("error", *mError)
    }
    if statusCode != nil {
        req.QueryParam("status_code", *statusCode)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if topic != nil {
        req.QueryParam("topic", *topic)
    }
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteWebhooksDeliveries takes context, siteId, webhookId, mError, statusCode, status, topic, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.SearchWebhookDelivery data and
// an error if there was an issue with the request or response.
// Search Site Webhooks deliveries
// Topics Supported:
// - alarms
// - audits
// - device-updowns
// - occupancy-alerts
// - ping
func (s *SitesWebhooks) SearchSiteWebhooksDeliveries(
    ctx context.Context,
    siteId uuid.UUID,
    webhookId uuid.UUID,
    mError *string,
    statusCode *int,
    status *models.WebhookDeliveryStatusEnum,
    topic *models.WebhookDeliveryTopicEnum,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.SearchWebhookDelivery],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/webhooks/%v/events/search")
    req.AppendTemplateParams(siteId, webhookId)
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
    if mError != nil {
        req.QueryParam("error", *mError)
    }
    if statusCode != nil {
        req.QueryParam("status_code", *statusCode)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if topic != nil {
        req.QueryParam("topic", *topic)
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.SearchWebhookDelivery
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SearchWebhookDelivery](decoder)
    return models.NewApiResponse(result, resp), err
}

// PingSiteWebhook takes context, siteId, webhookId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send a Ping event to the webhook
func (s *SitesWebhooks) PingSiteWebhook(
    ctx context.Context,
    siteId uuid.UUID,
    webhookId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/webhooks/%v/ping")
    req.AppendTemplateParams(siteId, webhookId)
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
