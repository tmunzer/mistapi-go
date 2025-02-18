
# Mxedge Tunterm Multicast Config

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mdns` | [`*models.MxedgeTuntermMulticastMdns`](../../doc/models/mxedge-tunterm-multicast-mdns.md) | Optional | - |
| `Ssdp` | [`*models.MxedgeTuntermMulticastSsdp`](../../doc/models/mxedge-tunterm-multicast-ssdp.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mdns": {
    "enabled": false,
    "vlan_ids": [
      "vlan_ids4",
      "vlan_ids5"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "ssdp": {
    "enabled": false,
    "vlan_ids": [
      "vlan_ids2",
      "vlan_ids3",
      "vlan_ids4"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

