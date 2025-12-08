
# Ap Port Config Dynamic Vlan

Optional dynamic vlan

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
    "1-10": null,
    "user": null
  },
  "enabled": false,
  "type": "type2"
}
```

