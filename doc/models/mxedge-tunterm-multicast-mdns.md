
# Mxedge Tunterm Multicast Mdns

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermMulticastMdns`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `VlanIds` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "vlan_ids": [
    "vlan_ids4",
    "vlan_ids5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

