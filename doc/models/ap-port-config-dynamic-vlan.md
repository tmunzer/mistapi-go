
# Ap Port Config Dynamic Vlan

optional dynamic vlan

## Structure

`ApPortConfigDynamicVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultVlanId` | `*int` | Optional | **Constraints**: `>= 1`, `<= 4094` |
| `Enabled` | `*bool` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Vlans` | `map[string]string` | Optional | - |

## Example (as JSON)

```json
{
  "default_vlan_id": 999,
  "vlans": {
    "1-10": "vlans5",
    "user": "vlans6"
  },
  "enabled": false,
  "type": "type0"
}
```

