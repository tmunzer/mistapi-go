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

// OrgsWebhooks represents a controller struct.
type OrgsWebhooks struct {
	baseController
}

// NewOrgsWebhooks creates a new instance of OrgsWebhooks.
// It takes a baseController as a parameter and returns a pointer to the OrgsWebhooks.
func NewOrgsWebhooks(baseController baseController) *OrgsWebhooks {
	orgsWebhooks := OrgsWebhooks{baseController: baseController}
	return &orgsWebhooks
}

// ListOrgWebhooks takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Webhook data and
// an error if there was an issue with the request or response.
// Get List of Org Webhooks
func (o *OrgsWebhooks) ListOrgWebhooks(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Webhook],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/webhooks")
	req.AppendTemplateParams(orgId)
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

// CreateOrgWebhook takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Webhook data and
// an error if there was an issue with the request or response.
// **N.B**. For org webhooks, only alarms/audits/client-info/client-join/client-sessions/device_events/device-updowns/mxedge_events Infrastructure topics are supported.
// Webhook defines a webhook, modeled after [github\u2019s model](https://developer.github.com/webhooks/).
// There is two types of webhooks:
// * webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-be01e694-7253-4195-8563-78e2a745e114))
// * raw data webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-e2d5d5f8-4bdb-4efc-93e4-90f4b33d0b2b))
// ##### Webhooks
// Webhooks can be configured at the org level (subset of topics only) and at the site level. It is possible to have multiple topics in the same webhook configuration and/or to have multiple webhooks configured at the same time.
func (o *OrgsWebhooks) CreateOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Webhook) (
	models.ApiResponse[models.Webhook],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/webhooks")
	req.AppendTemplateParams(orgId)
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

// DeleteOrgWebhook takes context, orgId, webhookId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Webhook
func (o *OrgsWebhooks) DeleteOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/webhooks/%v")
	req.AppendTemplateParams(orgId, webhookId)
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

// GetOrgWebhook takes context, orgId, webhookId as parameters and
// returns an models.ApiResponse with models.Webhook data and
// an error if there was an issue with the request or response.
// Get Org Webhook Details
func (o *OrgsWebhooks) GetOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID) (
	models.ApiResponse[models.Webhook],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/webhooks/%v")
	req.AppendTemplateParams(orgId, webhookId)
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

// UpdateOrgWebhook takes context, orgId, webhookId, body as parameters and
// returns an models.ApiResponse with models.Webhook data and
// an error if there was an issue with the request or response.
// Update Org Webhook
func (o *OrgsWebhooks) UpdateOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID,
	body *models.Webhook) (
	models.ApiResponse[models.Webhook],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/webhooks/%v")
	req.AppendTemplateParams(orgId, webhookId)
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

// CountOrgWebhooksDeliveries takes context, orgId, webhookId, mError, statusCode, status, topic, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Org Webhooks deliveries
// Topics Supported:
// - alarms
// - audits
// - device-updowns
// - occupancy-alerts
// - ping
func (o *OrgsWebhooks) CountOrgWebhooksDeliveries(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID,
	mError *string,
	statusCode *int,
	status *models.WebhookDeliveryStatusEnum,
	topic *models.WebhookDeliveryTopicEnum,
	distinct *models.WebhookDeliveryDistinctEnum,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/webhooks/%v/events/count")
	req.AppendTemplateParams(orgId, webhookId)
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

// SearchOrgWebhooksDeliveries takes context, orgId, webhookId, mError, statusCode, status, topic, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.SearchWebhookDelivery data and
// an error if there was an issue with the request or response.
// Search Org Webhooks deliveries
// Topics Supported:
// - alarms
// - audits
// - device-updowns
// - occupancy-alerts
// - ping
func (o *OrgsWebhooks) SearchOrgWebhooksDeliveries(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID,
	mError *string,
	statusCode *int,
	status *models.WebhookDeliveryStatusEnum,
	topic *models.WebhookDeliveryTopicEnum,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.SearchWebhookDelivery],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/webhooks/%v/events/search")
	req.AppendTemplateParams(orgId, webhookId)
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
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.SearchWebhookDelivery
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SearchWebhookDelivery](decoder)
	return models.NewApiResponse(result, resp), err
}

// PingOrgWebhook takes context, orgId, webhookId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send a Ping event to the webhook
func (o *OrgsWebhooks) PingOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/webhooks/%v/ping")
	req.AppendTemplateParams(orgId, webhookId)
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
