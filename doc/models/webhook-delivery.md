
# Webhook Delivery

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookDelivery`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `*string` | Optional | Error message, if there is one |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ReqHeaders` | `*string` | Optional | HTTP request headers |
| `ReqPayload` | `*string` | Optional | HTTP request payload |
| `ReqUrl` | `*string` | Optional | HTTP request URL |
| `RespBody` | `*string` | Optional | HTTP response body |
| `RespHeaders` | `*string` | Optional | HTTP response headers |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | [`*models.WebhookDeliveryStatusEnum`](../../doc/models/webhook-delivery-status-enum.md) | Optional | webhook delivery status. enum: `failure`, `success` |
| `StatusCode` | `*int` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Topic` | [`*models.WebhookDeliveryTopicEnum`](../../doc/models/webhook-delivery-topic-enum.md) | Optional | webhook topic. enum: `alarms`, `audits`, `device-updowns`, `occupancy-alerts`, `ping` |
| `WebhookId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "req_headers": "{\\\"Content-Type\\\":[\\\"application/json\\\"],\\\"User-Agent\\\":[\\\"Mist-webhook\\\"]}",
  "req_payload": "{\\\"topic\\\":\\\"audits\\\",\\\"events\\\":[{\\\"admin_name\\\":\\\"John Doe john.doe@juniper.net\\\",\\\"after\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": null, \\\\\"power_max\\\\\": null, \\\\\"power\\\\\": 10, \\\\\"preamble\\\\\": \\\\\"short\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"before\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": 8, \\\\\"power_max\\\\\": 18, \\\\\"power\\\\\": null, \\\\\"preamble\\\\\": \\\\\"long\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"id\\\":\\\"737909a2-04ff-4aeb-b9da-cc924e74a4dd\\\",\\\"message\\\":\\\"Update Site Settings\\\",\\\"org_id\\\":\\\"fc7e2967-e7ef-41e6-b007-1217713de05a\\\",\\\"site_id\\\":\\\"256c3a35-9cb7-436e-bc6d-314972645d95\\\",\\\"site_name\\\":\\\"Test Site\\\",\\\"src_ip\\\":\\\"1.2.3.4\\\",\\\"timestamp\\\":1685956576.923601,\\\"user_agent\\\":\\\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36\\\"}]}",
  "req_url": "http://example.com",
  "resp_body": "Ok",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "status": "failure",
  "status_code": 200,
  "timestamp": 1687962508.583656,
  "topic": "audits",
  "webhook_id": "7a11b901-f719-4c91-8aef-deb8699a6364",
  "error": "error0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

