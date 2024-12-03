
# Webhook Device Events

device event webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookDeviceEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookDeviceEventsEvent`](../../doc/models/webhook-device-events-event.md) | Required | list of events<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"device_events"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "device_name": "device_name8",
      "device_type": "switch",
      "ev_type": "notice",
      "mac": "mac4",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 130,
      "type": "type0",
      "ap": "ap6",
      "ap_name": "ap_name8",
      "audit_id": "00001d3a-0000-0000-0000-000000000000",
      "reason": "reason6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "device_events",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

