
# Aggregate Route

*This model accepts additional fields of type interface{}.*

## Structure

`AggregateRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | **Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 4294967295` |
| `Preference` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 4294967295` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "discard": false,
  "metric": 124,
  "preference": 52,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

