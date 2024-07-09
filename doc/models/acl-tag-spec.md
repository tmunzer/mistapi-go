
# Acl Tag Spec

## Structure

`AclTagSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | matched dst port, "0" means any<br>**Default**: `"80"` |
| `Protocol` | `*string` | Optional | `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`. `protocol_number` is between 1-254<br>**Default**: `"any"` |

## Example (as JSON)

```json
{
  "port_range": "80",
  "protocol": "any"
}
```

