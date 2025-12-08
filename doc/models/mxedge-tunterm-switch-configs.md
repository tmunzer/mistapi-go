
# Mxedge Tunterm Switch Configs

If custom vlan settings are desired

*This model accepts additional fields of type [models.MxedgeTuntermSwitchConfig](../../doc/models/mxedge-tunterm-switch-config.md).*

## Structure

`MxedgeTuntermSwitchConfigs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `AdditionalProperties` | [`map[string]models.MxedgeTuntermSwitchConfig`](../../doc/models/mxedge-tunterm-switch-config.md) | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "exampleAdditionalProperty": {
    "port_vlan_id": 176,
    "vlan_ids": [
      "String7",
      "String8"
    ]
  }
}
```

