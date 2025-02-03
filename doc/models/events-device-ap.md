
# Events Device Ap

AP events

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 21.42,
  "ap": "ap0",
  "apfw": "apfw0",
  "count": 140,
  "device_type": "device_type6",
  "mac": "mac8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

