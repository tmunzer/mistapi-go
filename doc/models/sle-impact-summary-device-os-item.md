
# Sle Impact Summary Device Os Item

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactSummaryDeviceOsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `float64` | Required | - |
| `DeviceOs` | `string` | Required | - |
| `Duration` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 192.96,
  "device_os": "device_os6",
  "duration": 66.02,
  "name": "name6",
  "total": 220.96,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

