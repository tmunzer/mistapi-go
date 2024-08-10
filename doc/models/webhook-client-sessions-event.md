
# Webhook Client Sessions Event

## Structure

`WebhookClientSessionsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | mac address of the AP the client roamed or disconnected from |
| `ApName` | `string` | Required | user-friendly name of the AP the client roamed or disconnected from. |
| `Band` | `string` | Required | 5GHz or 2.4GHz band |
| `Bssid` | `string` | Required | - |
| `ClientFamily` | `string` | Required | device family E.g. “Mac”, “iPhone”, “Apple watch” |
| `ClientManufacture` | `string` | Required | device manufacturer E.g. “Apple” |
| `ClientModel` | `string` | Required | device model E.g. “8+”, “XS” |
| `ClientOs` | `string` | Required | device operating system E.g. “Mojave”, “Windows 10”, “Linux” |
| `Connect` | `int` | Required | time when the user connects |
| `ConnectFloat` | `float64` | Required | floating point connect timestamp with millisecond precision |
| `Disconnect` | `int` | Required | time when the user disconnects |
| `DisconnectFloat` | `float64` | Required | floating point disconnect timestamp with millisecond precision |
| `Duration` | `int` | Required | the duration of the roamed or complete session indicated by termination_reason field. |
| `Mac` | `string` | Required | the client’s mac |
| `NextAp` | `string` | Required | the AP the client has roamed to. |
| `OrgId` | `uuid.UUID` | Required | - |
| `Rssi` | `float64` | Required | latest average RSSI before the user disconnects |
| `SiteId` | `uuid.UUID` | Required | - |
| `SiteName` | `string` | Required | - |
| `Ssid` | `string` | Required | - |
| `TerminationReason` | `int` | Required | 1 disassociate - when the client disassociates. 2 inactive - when the client is timeout. 3 roamed - when the client is roamed between APs |
| `Timestamp` | `float64` | Required | - |
| `Version` | `float64` | Required | schema version of this message |
| `WlanId` | `uuid.UUID` | Required | - |

## Example (as JSON)

```json
{
  "ap": "ap4",
  "ap_name": "ap_name6",
  "band": "band0",
  "bssid": "bssid2",
  "client_family": "client_family2",
  "client_manufacture": "client_manufacture8",
  "client_model": "client_model6",
  "client_os": "client_os6",
  "connect": 146,
  "connect_float": 5.8,
  "disconnect": 96,
  "disconnect_float": 155.58,
  "duration": 150,
  "mac": "mac2",
  "next_ap": "next_ap6",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rssi": 189.6,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "site_name": "site_name0",
  "ssid": "ssid6",
  "termination_reason": 24,
  "timestamp": 63.56,
  "version": 147.14,
  "wlan_id": "00000dde-0000-0000-0000-000000000000"
}
```

