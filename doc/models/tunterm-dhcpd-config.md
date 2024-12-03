
# Tunterm Dhcpd Config

DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID

*This model accepts additional fields of type [models.TuntermDhcpdConfigProperty](../../doc/models/tunterm-dhcpd-config-property.md).*

## Structure

`TuntermDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Servers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Type` | [`*models.TuntermDhcpdTypeEnum`](../../doc/models/tunterm-dhcpd-type-enum.md) | Optional | enum: `relay`<br>**Default**: `"relay"` |
| `AdditionalProperties` | [`map[string]models.TuntermDhcpdConfigProperty`](../../doc/models/tunterm-dhcpd-config-property.md) | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "relay",
  "servers": [
    "servers5",
    "servers6",
    "servers7"
  ],
  "exampleAdditionalProperty": {
    "enabled": false,
    "servers": [
      "servers7"
    ],
    "type": "relay",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

