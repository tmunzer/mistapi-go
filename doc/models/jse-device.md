
# Jse Device

*This model accepts additional fields of type interface{}.*

## Structure

`JseDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtIp` | `*string` | Optional | When available |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "last_seen": 1470417522.0,
  "ext_ip": "ext_ip6",
  "mac": "mac0",
  "model": "model6",
  "serial": "serial6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

