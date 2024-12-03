
# Mxedge Tunterm Dhcpd Config Property

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermDhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Servers` | `[]string` | Optional | list of DHCP servers; required if `type`==`relay` |
| `Type` | [`*models.MxedgeTuntermDhcpdConfigTypeEnum`](../../doc/models/mxedge-tunterm-dhcpd-config-type-enum.md) | Optional | enum: `relay`<br>**Default**: `"relay"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "relay",
  "servers": [
    "servers3",
    "servers4"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

