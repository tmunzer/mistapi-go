
# Mxedge Tunterm Igmp Snooping Config

IGMP snooping settings for Mist Tunnel VLANs

## Structure

`MxedgeTuntermIgmpSnoopingConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | [`*models.MxedgeTuntermIgmpSnoopingConfigEnabled`](../../doc/models/containers/mxedge-tunterm-igmp-snooping-config-enabled.md) | Optional | This is a container for any-of cases. |
| `Querier` | [`*models.MxedgeTuntermIgmpSnoopingQuerier`](../../doc/models/mxedge-tunterm-igmp-snooping-querier.md) | Optional | IGMP querier settings for tunnel termination |
| `VlanIds` | [`*models.MxedgeTuntermIgmpSnoopingConfigVlanIds`](../../doc/models/containers/mxedge-tunterm-igmp-snooping-config-vlan-ids.md) | Optional | List of VLAN IDs where tunnel termination performs IGMP snooping |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermIgmpSnoopingConfig := models.MxedgeTuntermIgmpSnoopingConfig{
        Enabled:              models.ToPointer(models.MxedgeTuntermIgmpSnoopingConfigEnabledContainer.FromBoolean(true)),
        Querier:              models.ToPointer(models.MxedgeTuntermIgmpSnoopingQuerier{
            MaxResponseTime:      models.ToPointer(136),
            Mtu:                  models.ToPointer(120),
            QueryInterval:        models.ToPointer(156),
            Robustness:           models.ToPointer(7),
            Version:              models.ToPointer(0),
        }),
        VlanIds:              models.ToPointer(models.MxedgeTuntermIgmpSnoopingConfigVlanIdsContainer.FromArrayOfNumber([]int{
            112,
            111,
        })),
    }

}
```

