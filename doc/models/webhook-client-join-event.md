
# Webhook Client Join Event

## Structure

`WebhookClientJoinEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | mac address of the AP the client connected to |
| `ApName` | `string` | Required | user-friendly name of the AP the client connected to. |
| `Band` | `string` | Required | 5GHz or 2.4GHz band |
| `Bssid` | `string` | Required | - |
| `Connect` | `int` | Required | time when the user connects |
| `ConnectFloat` | `float64` | Required | floating point connect timestamp with millisecond precision |
| `Mac` | `string` | Required | the client’s mac |
| `OrgId` | `uuid.UUID` | Required | - |
| `Rssi` | `float64` | Required | RSSI when the client associated |
| `SiteId` | `uuid.UUID` | Required | - |
| `SiteName` | `string` | Required | - |
| `Ssid` | `string` | Required | ESSID |
| `Timestamp` | `float64` | Required | - |
| `Version` | `float64` | Required | schema version of this message |
| `WlanId` | `uuid.UUID` | Required | - |

## Example (as JSON)

```json
{
  "ap": "ap8",
  "ap_name": "ap_name6",
  "band": "band8",
  "bssid": "bssid0",
  "connect": 8,
  "connect_float": 250.18,
  "mac": "mac0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rssi": 177.98,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "site_name": "site_name2",
  "ssid": "ssid6",
  "timestamp": 51.94,
  "version": 135.52,
  "wlan_id": "00001a84-0000-0000-0000-000000000000"
}
```

