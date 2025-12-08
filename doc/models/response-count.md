
# Response Count

## Structure

`ResponseCount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Distinct` | `string` | Required | - |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Results` | [`[]models.CountResult`](../../doc/models/count-result.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "distinct": "distinct8",
  "end": 0,
  "limit": 170,
  "results": [
    {
      "count": 226,
      "exampleAdditionalProperty": "count_result_additionalProperties2"
    }
  ],
  "start": 214,
  "total": 8
}
```

