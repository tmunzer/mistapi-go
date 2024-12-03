
# Mxedge Tunterm Igmp Snooping Config

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermIgmpSnoopingConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Querier` | [`*models.MxedgeTuntermIgmpSnoopingQuerier`](../../doc/models/mxedge-tunterm-igmp-snooping-querier.md) | Optional | - |
| `VlanIds` | `[]int` | Optional | the list of vlans on which tunterm performs IGMP snooping |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "querier": {
    "max_response_time": 136,
    "mtu": 120,
    "query_interval": 156,
    "robustness": 80,
    "version": 0,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "vlan_ids": [
    232
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

