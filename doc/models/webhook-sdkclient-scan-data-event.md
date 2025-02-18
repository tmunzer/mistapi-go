
# Webhook Sdkclient Scan Data Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookSdkclientScanDataEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConnectionAp` | `string` | Required | MAC address of the AP the client is connected to |
| `ConnectionBand` | `string` | Required | 5GHz or 2.4GHz band, of the BSSID the client is connected to |
| `ConnectionBssid` | `string` | Required | BSSID of the AP the client is connected to |
| `ConnectionChannel` | `int` | Required | Channel of the band the client is connected to |
| `ConnectionRssi` | `float64` | Required | RSSI of the client’s connection to the AP/BSSID |
| `LastSeen` | `float64` | Required | Time client last seen with scan data |
| `Mac` | `string` | Required | Client's MAC Address |
| `ScanData` | [`[]models.WebhookSdkclientScanDataEventScanDataItem`](../../doc/models/webhook-sdkclient-scan-data-event-scan-data-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `uuid.UUID` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "connection_ap": "connection_ap4",
  "connection_band": "connection_band8",
  "connection_bssid": "connection_bssid4",
  "connection_channel": 154,
  "connection_rssi": 55.12,
  "last_seen": 116.26,
  "mac": "mac0",
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
    },
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
    },
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
```

