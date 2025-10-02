
# Sle Classifier Samples

*This model accepts additional fields of type interface{}.*

## Structure

`SleClassifierSamples`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | - |
| `Duration` | `[]float64` | Required | - |
| `Total` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": [
    167.47,
    167.48
  ],
  "duration": [
    246.52
  ],
  "total": [
    19.61,
    19.6
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

