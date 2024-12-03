
# Extra Route 6 Next Qualified Properties

*This model accepts additional fields of type interface{}.*

## Structure

`ExtraRoute6NextQualifiedProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metric` | `models.Optional[int]` | Optional | - |
| `Preference` | `models.Optional[int]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "metric": 172,
  "preference": 4,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

