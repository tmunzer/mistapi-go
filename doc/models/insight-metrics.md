
# Insight Metrics

*This model accepts additional fields of type interface{}.*

## Structure

`InsightMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Interval` | `int` | Required | - |
| `Results` | `[]interface{}` | Required | Results depends on the `metric`<br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 16,
  "interval": 182,
  "results": [
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "start": 230,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

