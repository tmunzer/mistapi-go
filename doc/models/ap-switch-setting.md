
# Ap Switch Setting

*This model accepts additional fields of type interface{}.*

## Structure

`ApSwitchSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EnableVlan` | `*bool` | Optional | - |
| `PortVlanId` | [`*models.ApSwitchSettingPortVlanId`](../../doc/models/containers/ap-switch-setting-port-vlan-id.md) | Optional | Native VLAN id, optional |
| `VlanIds` | `[]int` | Optional | List of VLAN ids<br>**Constraints**: `>= 1`, `<= 4094` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enable_vlan": false,
  "port_vlan_id": 34,
  "vlan_ids": [
    1,
    1,
    2
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

