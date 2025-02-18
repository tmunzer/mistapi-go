
# Stats Device Other Vendor Specific Port

*This model accepts additional fields of type interface{}.*

## Structure

`StatsDeviceOtherVendorSpecificPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BytesIn` | `*int` | Optional | - |
| `BytesOut` | `*int` | Optional | - |
| `HealthCategory` | `*string` | Optional | - |
| `HealthScore` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Mode` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `State` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Uptime` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "bytes_in": 190,
  "bytes_out": 182,
  "health_category": "health_category2",
  "health_score": 184,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

