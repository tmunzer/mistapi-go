
# Extra Route Next Qualified Properties

*This model accepts additional fields of type interface{}.*

## Structure

`ExtraRouteNextQualifiedProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metric` | `models.Optional[int]` | Optional | - |
| `Preference` | `models.Optional[int]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "metric": 0,
  "preference": 176,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

