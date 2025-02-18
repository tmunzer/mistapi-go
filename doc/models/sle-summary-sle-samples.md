
# Sle Summary Sle Samples

*This model accepts additional fields of type interface{}.*

## Structure

`SleSummarySleSamples`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `[]float64` | Required | - |
| `Total` | `[]float64` | Required | - |
| `Value` | `[]float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": [
    92.75,
    92.76,
    92.77
  ],
  "total": [
    6.85,
    6.84
  ],
  "value": [
    192.18,
    192.19
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

