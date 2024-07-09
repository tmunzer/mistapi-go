
# Extra Route 6 Properties

## Structure

`ExtraRoute6Properties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | this takes precedence<br>**Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 4294967295` |
| `NextQualified` | [`map[string]models.ExtraRoute6PropertiesNextQualifiedProperties`](../../doc/models/extra-route-6-properties-next-qualified-properties.md) | Optional | - |
| `NoResolve` | `*bool` | Optional | **Default**: `false` |
| `Preference` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 4294967295` |
| `Via` | `*string` | Optional | next-hop IP Address |

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

