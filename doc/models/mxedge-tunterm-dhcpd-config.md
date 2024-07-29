
# Mxedge Tunterm Dhcpd Config

global and per-VLAN. Property key is the VLAN ID

## Structure

`MxedgeTuntermDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Servers` | `[]string` | Optional | list of DHCP servers; required if `type`==`relay` |
| `Type` | [`*models.MxedgeTuntermDhcpdTypeEnum`](../../doc/models/mxedge-tunterm-dhcpd-type-enum.md) | Optional | enum: `relay`<br>**Default**: `"relay"` |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "relay",
  "servers": [
    "servers7"
  ]
}
```

