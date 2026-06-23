
# Stats Cluster Config

High-availability cluster configuration and health reported by a gateway

## Structure

`StatsClusterConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Configuration` | `*string` | Optional | Configured high-availability mode for the gateway cluster, such as active-active |
| `ControlLinkInfo` | [`*models.StatsClusterConfigControlLinkInfo`](../../doc/models/stats-cluster-config-control-link-info.md) | Optional | Control link status details for a gateway cluster |
| `EthernetConnection` | [`[]models.StatsClusterConfigEthernetConnectionItem`](../../doc/models/stats-cluster-config-ethernet-connection-item.md) | Optional | Ethernet interface connection states for a gateway cluster |
| `FabricLinkInfo` | [`*models.StatsClusterConfigFabricLinkInfo`](../../doc/models/stats-cluster-config-fabric-link-info.md) | Optional | Fabric link status details for a gateway cluster |
| `LastStatusChangeReason` | `*string` | Optional | Reason reported for the most recent gateway cluster status change |
| `Operational` | `*string` | Optional | Current operational high-availability mode of the gateway cluster |
| `PrimaryNodeHealth` | `*string` | Optional | Health state reported for the primary gateway cluster node |
| `RedundancyGroupInformation` | [`[]models.StatsClusterConfigRedundancyGroupInfoItem`](../../doc/models/stats-cluster-config-redundancy-group-info-item.md) | Optional | Redundancy group monitoring entries for a gateway cluster |
| `SecondaryNodeHealth` | `*string` | Optional | Health state reported for the secondary gateway cluster node |
| `Status` | `*string` | Optional | Overall health status reported for the gateway cluster |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsClusterConfig := models.StatsClusterConfig{
        Configuration:              models.ToPointer("configuration4"),
        ControlLinkInfo:            models.ToPointer(models.StatsClusterConfigControlLinkInfo{
            Name:                 models.ToPointer("name0"),
            Status:               models.ToPointer("status2"),
        }),
        EthernetConnection:         []models.StatsClusterConfigEthernetConnectionItem{
            models.StatsClusterConfigEthernetConnectionItem{
                Name:                 models.ToPointer("name2"),
                Status:               models.ToPointer("status6"),
            },
            models.StatsClusterConfigEthernetConnectionItem{
                Name:                 models.ToPointer("name2"),
                Status:               models.ToPointer("status6"),
            },
            models.StatsClusterConfigEthernetConnectionItem{
                Name:                 models.ToPointer("name2"),
                Status:               models.ToPointer("status6"),
            },
        },
        FabricLinkInfo:             models.ToPointer(models.StatsClusterConfigFabricLinkInfo{
            DataPlaneNotifiedStatus: models.ToPointer("DataPlaneNotifiedStatus8"),
            Interface:               []string{
                "Interface0",
                "Interface1",
                "Interface2",
            },
            InternalStatus:          models.ToPointer("InternalStatus2"),
            State:                   models.ToPointer("State4"),
            Status:                  models.ToPointer("Status0"),
        }),
        LastStatusChangeReason:     models.ToPointer("last_status_change_reason2"),
    }

}
```

