
# Webhook Location Sdk Event

## Structure

`WebhookLocationSdkEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `MapId` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Type` | `*string` | Optional | **Default**: `"sdk"` |
| `X` | `*float64` | Optional | x, in meter |
| `Y` | `*float64` | Optional | y, in meter |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
  "name": "optional",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "sdk",
  "x": 13.5,
  "y": 3.2,
  "timestamp": 92.28
}
```

