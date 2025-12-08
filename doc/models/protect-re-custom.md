
# Protect Re Custom

Custom acls

## Structure

`ProtectReCustom`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Matched dst port, "0" means any<br><br>**Default**: `"0"` |
| `Protocol` | [`*models.ProtectReCustomProtocolEnum`](../../doc/models/protect-re-custom-protocol-enum.md) | Optional | enum: `any`, `icmp`, `tcp`, `udp`<br><br>**Default**: `"any"` |
| `Subnets` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "port_range": "80,1035-1040",
  "protocol": "any",
  "subnets": [
    "subnets5",
    "subnets4"
  ]
}
```

