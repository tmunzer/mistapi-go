
# Wirelss Client Session

## Structure

`WirelssClientSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | - |
| `Band` | `string` | Required | - |
| `ClientManufacture` | `*string` | Required | - |
| `Connect` | `int` | Required | - |
| `Disconnect` | `int` | Required | - |
| `Duration` | `float64` | Required | - |
| `ForSite` | `*bool` | Optional | - |
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
  "connect": 210,
  "disconnect": 160,
  "duration": 250.46,
  "mac": "mac4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "ssid2",
  "timestamp": 148.68,
  "wlan_id": "0000193e-0000-0000-0000-000000000000",
  "for_site": false,
  "tags": [
    "tags5",
    "tags6"
  ]
}
```

