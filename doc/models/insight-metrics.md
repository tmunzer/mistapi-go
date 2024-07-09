
# Insight Metrics

## Structure

`InsightMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Interval` | `int` | Required | - |
| `Results` | `[]interface{}` | Required | results depends on the `metric`<br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 110,
  "interval": 236,
  "results": [
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "start": 68
}
```

