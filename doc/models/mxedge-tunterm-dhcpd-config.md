
# Mxedge Tunterm Dhcpd Config

Global and per-VLAN. Property key is the VLAN ID

*This model accepts additional fields of type [models.MxedgeTuntermDhcpdConfigProperty](../../doc/models/mxedge-tunterm-dhcpd-config-property.md).*

## Structure

`MxedgeTuntermDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Servers` | `[]string` | Optional | List of DHCP servers; required if `type`==`relay` |
| `Type` | [`*models.MxedgeTuntermDhcpdTypeEnum`](../../doc/models/mxedge-tunterm-dhcpd-type-enum.md) | Optional | enum: `relay`<br><br>**Default**: `"relay"` |
| `AdditionalProperties` | [`map[string]models.MxedgeTuntermDhcpdConfigProperty`](../../doc/models/mxedge-tunterm-dhcpd-config-property.md) | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "relay",
  "servers": [
    "servers9",
    "servers0",
    "servers1"
  ],
  "exampleAdditionalProperty": {
    "enabled": false,
    "servers": [
      "servers9",
      "servers0",
      "servers1"
    ],
    "type": "relay",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

