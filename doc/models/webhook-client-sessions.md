
# Webhook Client Sessions

Client session webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookClientSessions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientSessionsEvent`](../../doc/models/webhook-client-sessions-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"client-sessions"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "ap": "ap6",
      "ap_name": "ap_name8",
      "band": "band2",
      "bssid": "bssid4",
      "client_family": "client_family4",
      "client_manufacture": "client_manufacture6",
      "client_model": "client_model4",
      "client_os": "client_os8",
      "connect": 64,
      "connect_float": 130.42,
      "disconnect": 14,
      "disconnect_float": 24.2,
      "duration": 68,
      "mac": "mac4",
      "next_ap": "next_ap8",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "rssi": 58.22,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "site_name": "site_name2",
      "ssid": "ssid8",
      "termination_reason": 198,
      "timestamp": 188.18,
      "version": 15.76,
      "wlan_id": "0000177c-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "client-sessions",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

