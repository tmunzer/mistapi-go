
# Sle Classifier Samples

*This model accepts additional fields of type interface{}.*

## Structure

`SleClassifierSamples`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `[]float64` | Required | - |
| `Duration` | `[]float64` | Required | - |
| `Total` | `[]float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": [
    180.93,
    180.94
  ],
  "duration": [
    246.52
  ],
  "total": [
    33.07,
    33.06
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

