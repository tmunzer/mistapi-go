
# Mxedge Tunterm Switch Config

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermSwitchConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortVlanId` | `*int` | Optional | - |
| `VlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_vlan_id": 66,
  "vlan_ids": [
    "String5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

