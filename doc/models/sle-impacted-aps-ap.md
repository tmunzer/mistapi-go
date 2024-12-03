
# Sle Impacted Aps Ap

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedApsAp`

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
  "ap_mac": "ap_mac2",
  "degraded": 10.9,
  "duration": 139.96,
  "name": "name0",
  "total": 217.1,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

