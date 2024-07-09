
# Ap Switch Setting

## Structure

`ApSwitchSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | `[]int` | Optional | additional VLAN IDs, only valid in mesh base mode |
| `EnableVlan` | `*bool` | Optional | - |
| `PortVlanId` | [`*models.ApSwitchSettingPortVlanId`](../../doc/models/containers/ap-switch-setting-port-vlan-id.md) | Optional | This is a container for one-of cases. |
| `VlanIds` | `[]int` | Optional | list of VLAN ids this |

## Example (as JSON)

```json
{
  "additional_vlan_ids": [
    27,
    28,
    29
  ],
  "enable_vlan": false,
  "port_vlan_id": 168,
  "vlan_ids": [
    254,
    255
  ]
}
```

