
# Utils Tunterm Bounce Port

## Structure

`UtilsTuntermBouncePort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `HoldTime` | `*int` | Optional | in milli seconds, hold time between multiple port bounces |
| `Ports` | `[]string` | Required | list of ports to bounce |

## Example (as JSON)

```json
{
  "hold_time": 210,
  "ports": [
    "ports9",
    "ports0"
  ]
}
```

