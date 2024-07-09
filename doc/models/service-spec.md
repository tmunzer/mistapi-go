
# Service Spec

## Structure

`ServiceSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | - |
| `Protocol` | `*string` | Optional | `https`/ `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`.<br>`protocol_number` is between 1-254<br>**Default**: `"any"` |

## Example (as JSON)

```json
{
  "port_range": "8080,8443",
  "protocol": "tcp"
}
```

