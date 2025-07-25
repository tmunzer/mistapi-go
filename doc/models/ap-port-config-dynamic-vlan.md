
# Ap Port Config Dynamic Vlan

Optional dynamic vlan

*This model accepts additional fields of type interface{}.*

## Structure

`ApPortConfigDynamicVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultVlanId` | `*int` | Optional | **Constraints**: `>= 1`, `<= 4094` |
| `Enabled` | `*bool` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Vlans` | `map[string]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "default_vlan_id": 999,
  "vlans": {
    "1-10": null,
    "user": null
  },
  "enabled": false,
  "type": "type2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

