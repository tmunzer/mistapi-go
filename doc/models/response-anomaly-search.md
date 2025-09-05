
# Response Anomaly Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAnomalySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Page` | `int` | Required | - |
| `Results` | [`[]models.Anomaly`](../../doc/models/anomaly.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1711035686,
  "limit": 10,
  "page": 1,
  "results": [
    {
      "events": [
        "events4"
      ],
      "since": 133.06,
      "sle_baseline": 108.54,
      "sle_deviation": 23.0,
      "timestamp": 2.64,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 1710949286,
  "total": 232,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

