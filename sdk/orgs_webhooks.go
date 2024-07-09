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

// ListOrgWebhooks takes context, orgId, page, limit as parameters and
// returns an models.ApiResponse with []models.Webhook data and
// an error if there was an issue with the request or response.
// Get List of Org Webhooks
func (o *OrgsWebhooks) ListOrgWebhooks(
	ctx context.Context,
	orgId uuid.UUID,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.Webhook],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks", orgId),
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
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
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
// Create Org Webhook
// **N.B**. For org webhooks, only alarms/audits/client-info/client-join/client-sessions/device_events/device-updowns/mxedge_events Infrastructure topics are supported.
func (o *OrgsWebhooks) CreateOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Webhook) (
	models.ApiResponse[models.Webhook],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks", orgId),
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

	var result models.Webhook
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Webhook](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgWebhook takes context, orgId, webhookId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org Webhook
func (o *OrgsWebhooks) DeleteOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks/%v", orgId, webhookId),
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
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks/%v", orgId, webhookId),
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
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks/%v", orgId, webhookId),
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

	var result models.Webhook
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Webhook](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgWebhooksDeliveries takes context, orgId, webhookId, mError, statusCode, status, topic, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
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
	start *int,
	end *int,
	duration *string,
	limit *int) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks/%v/events/count", orgId, webhookId),
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgWebhooksDeliveries takes context, orgId, webhookId, mError, statusCode, status, topic, start, end, duration, limit as parameters and
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
	start *int,
	end *int,
	duration *string,
	limit *int) (
	models.ApiResponse[models.SearchWebhookDelivery],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks/%v/events/search", orgId, webhookId),
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

// PingOrgWebhook takes context, orgId, webhookId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// send a Ping event to the webhook
func (o *OrgsWebhooks) PingOrgWebhook(
	ctx context.Context,
	orgId uuid.UUID,
	webhookId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/webhooks/%v/ping", orgId, webhookId),
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
