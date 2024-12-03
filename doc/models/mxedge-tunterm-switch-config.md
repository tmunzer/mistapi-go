
# Mxedge Tunterm Switch Config

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermSwitchConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortVlanId` | `*int` | Optional | - |
| `VlanIds` | [`[]models.MxedgeTuntermSwitchConfigVlanId`](../../doc/models/containers/mxedge-tunterm-switch-config-vlan-id.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_vlan_id": 66,
  "vlan_ids": [
    252
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

