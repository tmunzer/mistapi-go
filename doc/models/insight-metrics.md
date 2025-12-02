
# Insight Metrics

## Structure

`InsightMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Interval` | `int` | Required | - |
| `Limit` | `*int` | Optional | - |
| `Results` | `[]interface{}` | Required | Results depends on the `metric`<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 16,
  "interval": 182,
  "limit": 154,
  "results": [
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "start": 230
}
```

