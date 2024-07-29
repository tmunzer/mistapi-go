
# Webhook Sdkclient Scan Data

## Structure

`WebhookSdkclientScanData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookSdkclientScanDataEvent`](../../doc/models/webhook-sdkclient-scan-data-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `sdkclient_scan_data`<br>**Default**: `"sdkclient_scan_data"` |

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
      "last_seen": 94.7,
      "mac": "mac4",
      "scan_data": [
        {
          "ap": "ap6",
          "band": "2.4",
          "bssid": "bssid2",
          "channel": 72,
          "rssi": 228.1,
          "ssid": "ssid4",
          "timestamp": 102.06
        }
      ],
      "site_id": "0000245a-0000-0000-0000-000000000000"
    }
  ],
  "topic": "sdkclient_scan_data"
}
```

