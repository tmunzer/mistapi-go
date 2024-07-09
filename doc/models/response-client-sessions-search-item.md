
# Response Client Sessions Search Item

## Structure

`ResponseClientSessionsSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | - |
| `Band` | `string` | Required | - |
| `ClientManufacture` | `string` | Required | - |
| `Connect` | `float64` | Required | - |
| `Disconnect` | `float64` | Required | - |
| `Duration` | `float64` | Required | - |
| `Mac` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Ssid` | `string` | Required | - |
| `Tags` | `[]string` | Optional | - |
| `Timestamp` | `float64` | Required | - |
| `WlanId` | `uuid.UUID` | Required | - |

## Example (as JSON)

```json
{
  "ap": "ap4",
  "band": "band2",
  "client_manufacture": "client_manufacture6",
  "connect": 56.12,
  "disconnect": 234.82,
  "duration": 209.76,
  "mac": "mac4",
  "org_id": "000026a8-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "ssid2",
  "timestamp": 107.98,
  "wlan_id": "00000958-0000-0000-0000-000000000000",
  "tags": [
    "tags5",
    "tags6",
    "tags7"
  ]
}
```

