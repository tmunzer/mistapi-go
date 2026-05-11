
# Mxedge Tunterm Igmp Snooping Config

## Structure

`MxedgeTuntermIgmpSnoopingConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | [`*models.MxedgeTuntermIgmpSnoopingConfigEnabled`](../../doc/models/containers/mxedge-tunterm-igmp-snooping-config-enabled.md) | Optional | This is a container for any-of cases. |
| `Querier` | [`*models.MxedgeTuntermIgmpSnoopingQuerier`](../../doc/models/mxedge-tunterm-igmp-snooping-querier.md) | Optional | - |
| `VlanIds` | [`*models.MxedgeTuntermIgmpSnoopingConfigVlanIds`](../../doc/models/containers/mxedge-tunterm-igmp-snooping-config-vlan-ids.md) | Optional | List of vlans on which tunterm performs IGMP snooping |

## Example (as JSON)

```json
{
  "enabled": true,
  "querier": {
    "max_response_time": 136,
    "mtu": 120,
    "query_interval": 156,
    "robustness": 7,
    "version": 0
  },
  "vlan_ids": [
    16,
    17
  ]
}
```

