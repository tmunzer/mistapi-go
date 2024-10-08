
# Repsonse Count

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

## Example (as JSON)

```json
{
  "distinct": "distinct6",
  "end": 212,
  "limit": 214,
  "results": [
    {
      "count": 226
    }
  ],
  "start": 170,
  "total": 52
}
```

