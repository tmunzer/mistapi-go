
# Events Skyatp

SkyATP events

## Structure

`EventsSkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `string` | Required | - |
| `ForSite` | `*bool` | Optional | - |
| `Ip` | `string` | Required | - |
| `Mac` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `ThreatLevel` | `int` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `Type` | `string` | Required | - |

## Example (as JSON)

```json
{
  "device_mac": "device_mac8",
  "ip": "ip8",
  "mac": "mac8",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "threat_level": 248,
  "timestamp": 72.72,
  "type": "type6",
  "for_site": false
}
```

