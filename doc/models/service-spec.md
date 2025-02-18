
# Service Spec

*This model accepts additional fields of type interface{}.*

## Structure

`ServiceSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Port number, port range, or variable |
| `Protocol` | `*string` | Optional | `https`/ `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`, `protocol_number` is between 1-254<br>**Default**: `"any"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_range": "8080,8443",
  "protocol": "tcp",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

