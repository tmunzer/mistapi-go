
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
  "ap": "ap2",
  "band": "band4",
  "client_manufacture": "client_manufacture4",
  "connect": 190,
  "disconnect": 140,
  "duration": 229.78,
  "mac": "mac6",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "ssid0",
  "timestamp": 128.0,
  "wlan_id": "0000225a-0000-0000-0000-000000000000",
  "for_site": false,
  "tags": [
    "tags7",
    "tags8",
    "tags9"
  ]
}
```

