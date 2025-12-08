
# Extra Route 6

## Structure

`ExtraRoute6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | This takes precedence<br><br>**Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 2147483647` |
| `NextQualified` | [`map[string]models.ExtraRoute6NextQualifiedProperties`](../../doc/models/extra-route-6-next-qualified-properties.md) | Optional | - |
| `NoResolve` | `*bool` | Optional | **Default**: `false` |
| `Preference` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 2147483647` |
| `Via` | `*string` | Optional | Next-hop IP Address |

## Example (as JSON)

```json
{
  "discard": false,
  "next_qualified": {
    "2a02:1234:200a::100": {
      "metric": 170,
      "preference": 40
    }
  },
  "no_resolve": false,
  "preference": 30,
  "via": "2a02:1234:200a::100",
  "metric": 10
}
```

