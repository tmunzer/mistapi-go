
# Utils Mac Table

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsMacTable`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MacAddress` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `VlanId` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac_address": "f8c1165c6400",
  "port_id": "ge-0/0/0.0",
  "vlan_id": "ge-0/0/0.0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

