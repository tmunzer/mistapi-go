
# Wxlan Tag Spec

## Structure

`WxlanTagSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Matched destination port, "0" means any<br><br>**Default**: `"0"` |
| `Protocol` | `*string` | Optional | tcp / udp / icmp / gre / any / ":protocol_number", `protocol_number` is between 1-254<br><br>**Default**: `"any"` |
| `Subnets` | `[]string` | Optional | Matched destination subnets and/or IP Addresses |

## Example (as JSON)

```json
{
  "port_range": "0",
  "protocol": "any",
  "subnets": [
    "0.0.0.0/0"
  ]
}
```

