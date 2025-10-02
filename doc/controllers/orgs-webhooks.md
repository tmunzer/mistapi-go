# Orgs Webhooks

```go
orgsWebhooks := client.OrgsWebhooks()
```

## Class Name

`OrgsWebhooks`

## Methods

* [Count Org Webhooks Deliveries](../../doc/controllers/orgs-webhooks.md#count-org-webhooks-deliveries)
* [Create Org Webhook](../../doc/controllers/orgs-webhooks.md#create-org-webhook)
* [Delete Org Webhook](../../doc/controllers/orgs-webhooks.md#delete-org-webhook)
* [Get Org Webhook](../../doc/controllers/orgs-webhooks.md#get-org-webhook)
* [List Org Webhooks](../../doc/controllers/orgs-webhooks.md#list-org-webhooks)
* [Ping Org Webhook](../../doc/controllers/orgs-webhooks.md#ping-org-webhook)
* [Search Org Webhooks Deliveries](../../doc/controllers/orgs-webhooks.md#search-org-webhooks-deliveries)
* [Update Org Webhook](../../doc/controllers/orgs-webhooks.md#update-org-webhook)


# Count Org Webhooks Deliveries

Count Org Webhooks deliveries

Topics Supported:

- alarms
- audits
- device-updowns
- occupancy-alerts
- ping

```go
CountOrgWebhooksDeliveries(
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
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `webhookId` | `uuid.UUID` | Template, Required | - |
| `mError` | `*string` | Query, Optional | - |
| `statusCode` | `*int` | Query, Optional | - |
| `status` | [`*models.WebhookDeliveryStatusEnum`](../../doc/models/webhook-delivery-status-enum.md) | Query, Optional | Webhook delivery status |
| `topic` | [`*models.WebhookDeliveryTopicEnum`](../../doc/models/webhook-delivery-topic-enum.md) | Query, Optional | Webhook topic |
| `distinct` | [`*models.WebhookDeliveryDistinctEnum`](../../doc/models/webhook-delivery-distinct-enum.md) | Query, Optional | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

webhookId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

status := models.WebhookDeliveryStatusEnum_FAILURE

topic := models.WebhookDeliveryTopicEnum_AUDITS

distinct := models.WebhookDeliveryDistinctEnum_WEBHOOKID

duration := "10m"

limit := 100

apiResponse, err := orgsWebhooks.CountOrgWebhooksDeliveries(ctx, orgId, webhookId, nil, nil, &status, &topic, &distinct, nil, nil, &duration, &limit)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Create Org Webhook

**N.B**. For org webhooks, only alarms/audits/client-info/client-join/client-sessions/device_events/device-updowns/mxedge_events Infrastructure topics are supported.

Webhook defines a webhook, modeled after [github\u2019s model](https://developer.github.com/webhooks/).

There is two types of webhooks:

* webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-be01e694-7253-4195-8563-78e2a745e114))
* raw data webhooks ([examples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace/folder/224925-e2d5d5f8-4bdb-4efc-93e4-90f4b33d0b2b))

##### Webhooks

Webhooks can be configured at the org level (subset of topics only) and at the site level. It is possible to have multiple topics in the same webhook configuration and/or to have multiple webhooks configured at the same time.

```go
CreateOrgWebhook(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Webhook) (
    models.ApiResponse[models.Webhook],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Webhook`](../../doc/models/webhook.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Webhook](../../doc/models/webhook.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Webhook{
    Enabled:               models.ToPointer(true),
    Headers:               models.NewOptional(models.ToPointer(map[string]string{
        "x-custom-1": "your_custom_header_value1",
        "x-custom-2": "your_custom_header_value2",
    })),
    SingleEventPerMessage: models.ToPointer(false),
    Type:                  models.ToPointer(models.WebhookTypeEnum_HTTPPOST),
    VerifyCert:            models.ToPointer(true),
}

apiResponse, err := orgsWebhooks.CreateOrgWebhook(ctx, orgId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "created_time": 0,
  "enabled": true,
  "headers": {},
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "secret": "string",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "splunk_token": "string",
  "topics": [
    "location"
  ],
  "type": "http-post",
  "url": "string",
  "verify_cert": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400WebhookException`](../../doc/models/response-http-400-webhook-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Webhook

Delete Org Webhook

```go
DeleteOrgWebhook(
    ctx context.Context,
    orgId uuid.UUID,
    webhookId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `webhookId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

webhookId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsWebhooks.DeleteOrgWebhook(ctx, orgId, webhookId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Webhook

Get Org Webhook Details

```go
GetOrgWebhook(
    ctx context.Context,
    orgId uuid.UUID,
    webhookId uuid.UUID) (
    models.ApiResponse[models.Webhook],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `webhookId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Webhook](../../doc/models/webhook.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

webhookId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWebhooks.GetOrgWebhook(ctx, orgId, webhookId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "created_time": 0,
  "enabled": true,
  "headers": {},
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "secret": "string",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "splunk_token": "string",
  "topics": [
    "location"
  ],
  "type": "http-post",
  "url": "string",
  "verify_cert": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Webhooks

Get List of Org Webhooks

```go
ListOrgWebhooks(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Webhook],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Webhook](../../doc/models/webhook.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsWebhooks.ListOrgWebhooks(ctx, orgId, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "created_time": 0,
    "enabled": true,
    "headers": {},
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "secret": "string",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "splunk_token": "string",
    "topics": [
      "location"
    ],
    "type": "http-post",
    "url": "string",
    "verify_cert": true
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Ping Org Webhook

Send a Ping event to the webhook

```go
PingOrgWebhook(
    ctx context.Context,
    orgId uuid.UUID,
    webhookId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `webhookId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

webhookId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsWebhooks.PingOrgWebhook(ctx, orgId, webhookId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Org Webhooks Deliveries

Search Org Webhooks deliveries

Topics Supported:

- alarms
- audits
- device-updowns
- occupancy-alerts
- ping

```go
SearchOrgWebhooksDeliveries(
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
    sort *string) (
    models.ApiResponse[models.SearchWebhookDelivery],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `webhookId` | `uuid.UUID` | Template, Required | - |
| `mError` | `*string` | Query, Optional | - |
| `statusCode` | `*int` | Query, Optional | - |
| `status` | [`*models.WebhookDeliveryStatusEnum`](../../doc/models/webhook-delivery-status-enum.md) | Query, Optional | Webhook delivery status |
| `topic` | [`*models.WebhookDeliveryTopicEnum`](../../doc/models/webhook-delivery-topic-enum.md) | Query, Optional | Webhook topic |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWebhookDelivery](../../doc/models/search-webhook-delivery.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

webhookId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

status := models.WebhookDeliveryStatusEnum_FAILURE

topic := models.WebhookDeliveryTopicEnum_AUDITS

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsWebhooks.SearchOrgWebhooksDeliveries(ctx, orgId, webhookId, nil, nil, &status, &topic, &limit, nil, nil, &duration, &sort)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "end": 1688035193,
  "limit": 10,
  "results": [
    {
      "error": "string",
      "id": "55b0f02f-ebf6-4ad2-8b10-200508a97581",
      "org_id": "fc7e2967-e7ef-41e6-b007-1217713de05a",
      "req_headers": "{\\\"Content-Type\\\":[\\\"application/json\\\"],\\\"User-Agent\\\":[\\\"Mist-webhook\\\"]}",
      "req_payload": "{\\\"topic\\\":\\\"audits\\\",\\\"events\\\":[{\\\"admin_name\\\":\\\"John Doe john.doe@juniper.net\\\",\\\"after\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": null, \\\\\"power_max\\\\\": null, \\\\\"power\\\\\": 10, \\\\\"preamble\\\\\": \\\\\"short\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"before\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": 8, \\\\\"power_max\\\\\": 18, \\\\\"power\\\\\": null, \\\\\"preamble\\\\\": \\\\\"long\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"id\\\":\\\"737909a2-04ff-4aeb-b9da-cc924e74a4dd\\\",\\\"message\\\":\\\"Update Site Settings\\\",\\\"org_id\\\":\\\"fc7e2967-e7ef-41e6-b007-1217713de05a\\\",\\\"site_id\\\":\\\"256c3a35-9cb7-436e-bc6d-314972645d95\\\",\\\"site_name\\\":\\\"Test Site\\\",\\\"src_ip\\\":\\\"1.2.3.4\\\",\\\"timestamp\\\":1685956576.923601,\\\"user_agent\\\":\\\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36\\\"}]}",
      "req_url": "https://example.com",
      "resp_body": "Ok",
      "resp_headers": "string",
      "site_id": "256c3a35-9cb7-436e-bc6d-314972645d95",
      "status": "success",
      "status_code": 200,
      "timestamp": 1687962508.583656,
      "topic": "audits",
      "webhook_id": "7a11b901-f719-4c91-8aef-deb8699a6364"
    }
  ],
  "start": 1687948793,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Webhook

Update Org Webhook

```go
UpdateOrgWebhook(
    ctx context.Context,
    orgId uuid.UUID,
    webhookId uuid.UUID,
    body *models.Webhook) (
    models.ApiResponse[models.Webhook],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `webhookId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Webhook`](../../doc/models/webhook.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Webhook](../../doc/models/webhook.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

webhookId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Webhook{
    Enabled:               models.ToPointer(true),
    Headers:               models.NewOptional(models.ToPointer(map[string]string{
        "x-custom-1": "your_custom_header_value1",
        "x-custom-2": "your_custom_header_value2",
    })),
    SingleEventPerMessage: models.ToPointer(false),
    Type:                  models.ToPointer(models.WebhookTypeEnum_HTTPPOST),
    VerifyCert:            models.ToPointer(true),
}

apiResponse, err := orgsWebhooks.UpdateOrgWebhook(ctx, orgId, webhookId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "created_time": 0,
  "enabled": true,
  "headers": {},
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "secret": "string",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "splunk_token": "string",
  "topics": [
    "location"
  ],
  "type": "http-post",
  "url": "string",
  "verify_cert": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

