
# Mxedge Tunterm Switch Config

## Structure

`MxedgeTuntermSwitchConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortVlanId` | `*int` | Optional | - |
| `VlanIds` | [`[]models.MxedgeTuntermSwitchConfigVlanIds`](../../doc/models/containers/mxedge-tunterm-switch-config-vlan-ids.md) | Optional | This is Array of a container for any-of cases. |

## Example (as JSON)

```json
{
  "port_vlan_id": 0,
  "vlan_ids": [
    186,
    187,
    188
  ]
}
```

