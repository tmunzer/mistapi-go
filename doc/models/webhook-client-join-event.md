
# Webhook Client Join Event

## Structure

`WebhookClientJoinEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | MAC address of the AP the client connected to |
| `ApName` | `string` | Required | user-friendly name of the AP the client connected to. |
| `Band` | `string` | Required | 5GHz or 2.4GHz band |
| `Bssid` | `string` | Required | - |
| `Connect` | `int` | Required | Time when the user connects |
| `ConnectFloat` | `float64` | Required | floating point connect timestamp with millisecond precision |
| `Mac` | `string` | Required | Client's MAC Address |
| `OrgId` | `uuid.UUID` | Required | - |
| `Rssi` | `float64` | Required | RSSI when the client associated |
| `SiteId` | `uuid.UUID` | Required | - |
| `SiteName` | `string` | Required | - |
| `Ssid` | `string` | Required | ESSID |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Version` | `float64` | Required | schema version of this message |
| `WlanId` | `uuid.UUID` | Required | - |

## Example (as JSON)

```json
{
  "ap": "ap8",
  "ap_name": "ap_name4",
  "band": "band8",
  "bssid": "bssid0",
  "connect": 34,
  "connect_float": 43.08,
  "mac": "mac0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rssi": 226.88,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "site_name": "site_name2",
  "ssid": "ssid6",
  "timestamp": 100.84,
  "version": 184.42,
  "wlan_id": "0000068e-0000-0000-0000-000000000000"
}
```

