
# Tunterm Dhcpd Config

DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID

## Structure

`TuntermDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Servers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Type` | [`*models.TuntermDhcpdTypeEnum`](../../doc/models/tunterm-dhcpd-type-enum.md) | Optional | enum: `relay`<br>**Default**: `"relay"` |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "relay",
  "servers": [
    "servers1",
    "servers2"
  ]
}
```

