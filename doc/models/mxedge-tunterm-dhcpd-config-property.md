
# Mxedge Tunterm Dhcpd Config Property

## Structure

`MxedgeTuntermDhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Servers` | `[]string` | Optional | list of DHCP servers; required if `type`==`relay` |
| `Type` | [`*models.MxedgeTuntermDhcpdConfigTypeEnum`](../../doc/models/mxedge-tunterm-dhcpd-config-type-enum.md) | Optional | enum: `relay`<br>**Default**: `"relay"` |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "relay",
  "servers": [
    "servers7",
    "servers8",
    "servers9"
  ]
}
```

