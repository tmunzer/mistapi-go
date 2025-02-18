
# Sle Impact Summary Device Type Item

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactSummaryDeviceTypeItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `float64` | Required | - |
| `DeviceType` | `string` | Required | - |
| `Duration` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 230.14,
  "device_type": "device_type6",
  "duration": 103.2,
  "name": "name4",
  "total": 253.86,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

