
# Ap Switch Setting

## Structure

`ApSwitchSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EnableVlan` | `*bool` | Optional | - |
| `PortVlanId` | `*int` | Optional | native VLAN id, optional<br>**Constraints**: `>= 1`, `<= 4094` |
| `VlanIds` | `[]int` | Optional | list of VLAN ids this |

## Example (as JSON)

```json
{
  "enable_vlan": false,
  "port_vlan_id": 112,
  "vlan_ids": [
    0,
    1,
    2
  ]
}
```

