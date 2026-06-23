
# Webhook Delivery

Record of a webhook delivery attempt and the destination response

## Structure

`WebhookDelivery`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `*string` | Optional | Message returned for a failed delivery attempt, when available |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `ReqHeaders` | `*string` | Optional | HTTP headers sent with the webhook delivery request |
| `ReqPayload` | `*string` | Optional | JSON payload sent in the webhook delivery request |
| `ReqUrl` | `*string` | Optional | Destination URL used for the webhook delivery request |
| `RespBody` | `*string` | Optional | Response body returned by the webhook destination |
| `RespHeaders` | `*string` | Optional | Response headers returned by the webhook destination |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Status` | [`*models.WebhookDeliveryStatusEnum`](../../doc/models/webhook-delivery-status-enum.md) | Optional | webhook delivery status. enum: `failure`, `success` |
| `StatusCode` | `*int` | Optional | HTTP status code returned by the webhook destination |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Topic` | [`*models.WebhookDeliveryTopicEnum`](../../doc/models/webhook-delivery-topic-enum.md) | Optional | webhook topic. enum: `alarms`, `audits`, `device-updowns`, `occupancy-alerts`, `ping` |
| `WebhookId` | `*uuid.UUID` | Optional | Unique identifier of the configured webhook used for this delivery attempt |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookDelivery := models.WebhookDelivery{
        Error:                models.ToPointer("error4"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        ReqHeaders:           models.ToPointer("{\\\"Content-Type\\\":[\\\"application/json\\\"],\\\"User-Agent\\\":[\\\"Mist-webhook\\\"]}"),
        ReqPayload:           models.ToPointer("{\\\"topic\\\":\\\"audits\\\",\\\"events\\\":[{\\\"admin_name\\\":\\\"John Doe john.doe@juniper.net\\\",\\\"after\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": null, \\\\\"power_max\\\\\": null, \\\\\"power\\\\\": 10, \\\\\"preamble\\\\\": \\\\\"short\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"before\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": 8, \\\\\"power_max\\\\\": 18, \\\\\"power\\\\\": null, \\\\\"preamble\\\\\": \\\\\"long\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"id\\\":\\\"737909a2-04ff-4aeb-b9da-cc924e74a4dd\\\",\\\"message\\\":\\\"Update Site Settings\\\",\\\"org_id\\\":\\\"fc7e2967-e7ef-41e6-b007-1217713de05a\\\",\\\"site_id\\\":\\\"256c3a35-9cb7-436e-bc6d-314972645d95\\\",\\\"site_name\\\":\\\"Test Site\\\",\\\"src_ip\\\":\\\"1.2.3.4\\\",\\\"timestamp\\\":1685956576.923601,\\\"user_agent\\\":\\\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36\\\"}]}"),
        ReqUrl:               models.ToPointer("https://example.com"),
        RespBody:             models.ToPointer("Ok"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Status:               models.ToPointer(models.WebhookDeliveryStatusEnum_FAILURE),
        StatusCode:           models.ToPointer(200),
        Topic:                models.ToPointer(models.WebhookDeliveryTopicEnum_AUDITS),
        WebhookId:            models.ToPointer(uuid.MustParse("7a11b901-f719-4c91-8aef-deb8699a6364")),
    }

}
```

