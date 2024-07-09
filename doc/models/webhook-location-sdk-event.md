
# Webhook Location Sdk Event

## Structure

`WebhookLocationSdkEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*int` | Optional | - |
| `Type` | `*string` | Optional | **Default**: `"sdk"` |
| `X` | `*float64` | Optional | x, in meter |
| `Y` | `*float64` | Optional | y, in meter |

## Example (as JSON)

```json
{
  "id": "de87bf9d-183f-e383-cc68-6ba43947d403",
  "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
  "name": "optional",
  "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "timestamp": 1461220784,
  "type": "sdk",
  "x": 13.5,
  "y": 3.2
}
```

