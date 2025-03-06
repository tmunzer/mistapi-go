
# Webhook Sdkclient Scan Data

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookSdkclientScanData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookSdkclientScanDataEvent`](../../doc/models/webhook-sdkclient-scan-data-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `sdkclient_scan_data`<br>**Value**: `"sdkclient_scan_data"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "connection_ap": "connection_ap0",
      "connection_band": "connection_band2",
      "connection_bssid": "connection_bssid8",
      "connection_channel": 46,
      "connection_rssi": 33.56,
      "last_seen": 1470417522.0,
      "mac": "mac4",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "scan_data": [
        {
          "ap": "ap6",
          "band": "2.4",
          "bssid": "bssid2",
          "channel": 72,
          "rssi": 228.1,
          "ssid": "ssid4",
          "timestamp": 102.06,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "sdkclient_scan_data",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

