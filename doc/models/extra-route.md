
# Extra Route

## Structure

`ExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | This takes precedence<br><br>**Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 2147483647` |
| `NextQualified` | [`map[string]models.ExtraRouteNextQualifiedProperties`](../../doc/models/extra-route-next-qualified-properties.md) | Optional | - |
| `NoResolve` | `*bool` | Optional | **Default**: `false` |
| `Preference` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 2147483647` |
| `Via` | [`*models.NextHopVia`](../../doc/models/containers/next-hop-via.md) | Optional | Next-hop IP Address. Can be a single IP address or an array of IP addresses for ECMP (Equal-Cost Multi-Path) load balancing across multiple next-hops. |

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

