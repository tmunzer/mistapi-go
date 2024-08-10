
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
  "connect": 100.32,
  "disconnect": 23.02,
  "duration": 253.96,
  "mac": "mac4",
  "org_id": "00000504-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "ssid2",
  "timestamp": 103.82,
  "wlan_id": "000004bc-0000-0000-0000-000000000000",
  "tags": [
    "tags5",
    "tags6",
    "tags7"
  ]
}
```

