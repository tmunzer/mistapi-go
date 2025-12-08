
# Aggregate Route

## Structure

`AggregateRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | **Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 4294967295` |
| `Preference` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 4294967295` |

## Example (as JSON)

```json
{
  "discard": false,
  "metric": 124,
  "preference": 52
}
```

