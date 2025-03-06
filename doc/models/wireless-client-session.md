
# Wireless Client Session

*This model accepts additional fields of type interface{}.*

## Structure

`WirelessClientSession`

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
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `WlanId` | `uuid.UUID` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "ap0",
  "band": "band6",
  "client_manufacture": "client_manufacture2",
  "connect": 174,
  "disconnect": 124,
  "duration": 173.3,
  "mac": "mac8",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "ssid2",
  "timestamp": 71.52,
  "wlan_id": "000010fa-0000-0000-0000-000000000000",
  "for_site": false,
  "tags": [
    "tags9",
    "tags0",
    "tags1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

