
# Mxedge Tunterm Igmp Snooping Config

## Structure

`MxedgeTuntermIgmpSnoopingConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Querier` | [`*models.MxedgeTuntermIgmpSnoopingQuerier`](../../doc/models/mxedge-tunterm-igmp-snooping-querier.md) | Optional | - |
| `VlanIds` | `[]int` | Optional | the list of vlans on which tunterm performs IGMP snooping |

## Example (as JSON)

```json
{
  "enabled": false,
  "querier": {
    "max_response_time": 136,
    "mtu": 120,
    "query_interval": 156,
    "robustness": 80,
    "version": 0
  },
  "vlan_ids": [
    232
  ]
}
```

