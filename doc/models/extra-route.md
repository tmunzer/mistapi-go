
# Extra Route

## Structure

`ExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | this takes precedence<br>**Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 2147483647` |
| `NextQualified` | [`map[string]models.ExtraRouteNextQualifiedProperties`](../../doc/models/extra-route-next-qualified-properties.md) | Optional | - |
| `NoResolve` | `*bool` | Optional | **Default**: `false` |
| `Preference` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 2147483647` |
| `Via` | `*string` | Optional | next-hop IP Address |

## Example (as JSON)

```json
{
  "discard": false,
  "next_qualified": {
    "10.3.1.1": {
      "metric": 170,
      "preference": 40
    }
  },
  "no_resolve": false,
  "preference": 30,
  "via": "10.2.1.1",
  "metric": 168
}
```

