
# Wxlan Tag Spec

## Structure

`WxlanTagSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | matched dst port, "0" means any<br>**Default**: `"0"` |
| `Protocol` | `*string` | Optional | tcp / udp / icmp / gre / any / ":protocol_number", `protocol_number` is between 1-254<br>**Default**: `"any"` |
| `Subnets` | `[]string` | Optional | matched dst subnet |

## Example (as JSON)

```json
{
  "port_range": "0",
  "protocol": "any",
  "subnets": [
    "subnets3"
  ]
}
```

