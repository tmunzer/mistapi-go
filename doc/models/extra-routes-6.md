
# Extra Routes 6

Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")

## Structure

`ExtraRoutes6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "2a02:1234:420a:10c9::/64": {
    "via": "2a02:1234:200a::100"
  },
  "via": "via6"
}
```

