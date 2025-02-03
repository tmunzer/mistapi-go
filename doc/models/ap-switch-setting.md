
# Ap Switch Setting

*This model accepts additional fields of type interface{}.*

## Structure

`ApSwitchSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EnableVlan` | `*bool` | Optional | - |
| `PortVlanId` | `*int` | Optional | Native VLAN id, optional<br>**Constraints**: `>= 1`, `<= 4094` |
| `VlanIds` | `[]int` | Optional | List of VLAN ids this |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enable_vlan": false,
  "port_vlan_id": 112,
  "vlan_ids": [
    0,
    1,
    2
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

