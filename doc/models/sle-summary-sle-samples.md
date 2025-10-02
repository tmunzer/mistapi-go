
# Sle Summary Sle Samples

*This model accepts additional fields of type interface{}.*

## Structure

`SleSummarySleSamples`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | - |
| `Total` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | - |
| `Value` | [`[]models.NumberOrNull`](../../doc/models/containers/number-or-null.md) | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": [
    79.29,
    79.3,
    79.31
  ],
  "total": [
    249.39,
    249.38
  ],
  "value": [
    178.72,
    178.73
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

