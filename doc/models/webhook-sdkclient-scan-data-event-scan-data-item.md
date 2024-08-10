
# Webhook Sdkclient Scan Data Event Scan Data Item

## Structure

`WebhookSdkclientScanDataEventScanDataItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | mac address of the AP associated with the BSSID scanned |
| `Band` | [`models.ScanDataItemBandEnum`](../../doc/models/scan-data-item-band-enum.md) | Required | 5GHz or 2.4GHz band, associated with the BSSID scanned. enum: `2.4`, `5`<br>**Constraints**: *Minimum Length*: `1` |
| `Bssid` | `string` | Required | BSSID found during client’s background scan for wifi |
| `Channel` | `int` | Required | channel of the band found in the scan |
| `Rssi` | `float64` | Required | client’s RSSI relative to the BSSID scanned |
| `Ssid` | `string` | Required | ESSID containing the BSSID scanned |
| `Timestamp` | `float64` | Required | time the scan of the particular BSSID occurred |

## Example (as JSON)

```json
{
  "ap": "ap0",
  "band": "2.4",
  "bssid": "bssid8",
  "channel": 26,
  "rssi": 10.96,
  "ssid": "ssid2",
  "timestamp": 140.92
}
```

