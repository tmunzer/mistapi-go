
# Protect Re Custom

custom acls

## Structure

`ProtectReCustom`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | matched dst port, "0" means any<br>**Default**: `"0"` |
| `Protocol` | [`*models.ProtectReCustomProtocolEnum`](../../doc/models/protect-re-custom-protocol-enum.md) | Optional | enum: `any`, `icmp`, `tcp`, `udp`<br>**Default**: `"any"` |
| `Subnet` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "port_range": "80,1035-1040",
  "protocol": "any",
  "subnet": [
    "subnet1"
  ]
}
```

