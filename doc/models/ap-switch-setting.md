
# Ap Switch Setting

## Structure

`ApSwitchSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | `[]int` | Optional | additional VLAN IDs, only valid in mesh base mode |
| `EnableVlan` | `*bool` | Optional | - |
| `PortVlanId` | `*int` | Optional | native VLAN id, optional<br>**Constraints**: `>= 1`, `<= 4094` |
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
  "port_vlan_id": 246,
  "vlan_ids": [
    254,
    255
  ]
}
```

