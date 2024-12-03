
# Utils Tunterm Bounce Port

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsTuntermBouncePort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `HoldTime` | `*int` | Optional | in milli seconds, hold time between multiple port bounces |
| `Ports` | `[]string` | Required | list of ports to bounce |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "hold_time": 210,
  "ports": [
    "ports9",
    "ports0"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

