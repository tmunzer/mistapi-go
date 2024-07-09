
# Events Device Ap

ap events

## Structure

`EventsDeviceAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `Apfw` | `*string` | Optional | - |
| `Count` | `*int` | Optional | - |
| `DeviceType` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Text` | `*string` | Optional | - |
| `Timestamp` | `float64` | Required | - |
| `Type` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 9.8,
  "ap": "ap2",
  "apfw": "apfw2",
  "count": 22,
  "device_type": "device_type8",
  "mac": "mac6"
}
```

