
# Webhook Client Join

client join webhook sample

## Structure

`WebhookClientJoin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientJoinEvent`](../../doc/models/webhook-client-join-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"client-join"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "ap": "ap6",
      "ap_name": "ap_name8",
      "band": "band2",
      "bssid": "bssid4",
      "connect": 64,
      "connect_float": 130.42,
      "mac": "mac4",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "rssi": 58.22,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "site_name": "site_name2",
      "ssid": "ssid8",
      "timestamp": 188.18,
      "version": 15.76,
      "wlan_id": "0000177c-0000-0000-0000-000000000000"
    }
  ],
  "topic": "client-join"
}
```

