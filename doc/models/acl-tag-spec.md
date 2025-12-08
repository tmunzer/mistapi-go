
# Acl Tag Spec

## Structure

`AclTagSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Matched dst port, "0" means any<br><br>**Default**: `"0"` |
| `Protocol` | `*string` | Optional | `tcp` / `udp` / `icmp` / `icmp6` / `gre` / `any` / `:protocol_number`, `protocol_number` is between 1-254, default is `any` `protocol_number` is between 1-254<br><br>**Default**: `"any"` |

## Example (as JSON)

```json
{
  "port_range": "0",
  "protocol": "any"
}
```

