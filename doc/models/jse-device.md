
# Jse Device

*This model accepts additional fields of type interface{}.*

## Structure

`JseDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtIp` | `*string` | Optional | When available |
| `LastSeen` | `*float64` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ext_ip": "ext_ip6",
  "last_seen": 77.36,
  "mac": "mac0",
  "model": "model6",
  "serial": "serial6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

