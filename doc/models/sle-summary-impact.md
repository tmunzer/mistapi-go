
# Sle Summary Impact

*This model accepts additional fields of type interface{}.*

## Structure

`SleSummaryImpact`

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
  "num_aps": 9.92,
  "num_users": 245.16,
  "total_aps": 228.44,
  "total_users": 19.54,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

