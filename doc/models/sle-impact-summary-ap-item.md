
# Sle Impact Summary Ap Item

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactSummaryApItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Degraded` | `float64` | Required | - |
| `Duration` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac8",
  "degraded": 172.76,
  "duration": 45.82,
  "name": "name6",
  "total": 55.24,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

