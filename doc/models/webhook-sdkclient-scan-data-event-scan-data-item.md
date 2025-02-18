
# Webhook Sdkclient Scan Data Event Scan Data Item

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookSdkclientScanDataEventScanDataItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | MAC address of the AP associated with the BSSID scanned |
| `Band` | [`models.ScanDataItemBandEnum`](../../doc/models/scan-data-item-band-enum.md) | Required | 5GHz or 2.4GHz band, associated with the BSSID scanned. enum: `2.4`, `5`<br>**Constraints**: *Minimum Length*: `1` |
| `Bssid` | `string` | Required | BSSID found during client’s background scan for wifi |
| `Channel` | `int` | Required | Channel of the band found in the scan |
| `Rssi` | `float64` | Required | Client's RSSI relative to the BSSID scanned |
| `Ssid` | `string` | Required | ESSID containing the BSSID scanned |
| `Timestamp` | `float64` | Required | Time the scan of the particular BSSID occurred |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "ap0",
  "band": "2.4",
  "bssid": "bssid8",
  "channel": 26,
  "rssi": 10.96,
  "ssid": "ssid2",
  "timestamp": 140.92,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

