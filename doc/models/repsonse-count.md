
# Repsonse Count

*This model accepts additional fields of type interface{}.*

## Structure

`RepsonseCount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Distinct` | `string` | Required | - |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Results` | [`[]models.CountResult`](../../doc/models/count-result.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "distinct": "distinct6",
  "end": 212,
  "limit": 214,
  "results": [
    {
      "count": 226,
      "exampleAdditionalProperty": "count_result_additionalProperties2"
    }
  ],
  "start": 170,
  "total": 52,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

