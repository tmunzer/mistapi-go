
# Sle Classifier Impact

*This model accepts additional fields of type interface{}.*

## Structure

`SleClassifierImpact`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `float64` | Required | - |
| `NumUsers` | `float64` | Required | - |
| `TotalAps` | `float64` | Required | - |
| `TotalUsers` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_aps": 7.66,
  "num_users": 6.74,
  "total_aps": 246.02,
  "total_users": 1.96,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

