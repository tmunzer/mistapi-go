
# Stats Cluster Config

## Structure

`StatsClusterConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Configuration` | `*string` | Optional | - |
| `ControlLinkInfo` | [`*models.StatsClusterConfigControlLinkInfo`](../../doc/models/stats-cluster-config-control-link-info.md) | Optional | - |
| `EthernetConnection` | [`[]models.StatsClusterConfigEthernetConnectionItem`](../../doc/models/stats-cluster-config-ethernet-connection-item.md) | Optional | - |
| `FabricLinkInfo` | [`*models.StatsClusterConfigFabricLinkInfo`](../../doc/models/stats-cluster-config-fabric-link-info.md) | Optional | - |
| `LastStatusChangeReason` | `*string` | Optional | - |
| `Operational` | `*string` | Optional | - |
| `PrimaryNodeHealth` | `*string` | Optional | - |
| `RedundancyGroupInformation` | [`[]models.StatsClusterConfigRedundancyGroupInfoItem`](../../doc/models/stats-cluster-config-redundancy-group-info-item.md) | Optional | - |
| `SecondaryNodeHealth` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "configuration": "configuration4",
  "control_link_info": {
    "name": "name0",
    "status": "status2"
  },
  "ethernet_connection": [
    {
      "name": "name2",
      "status": "status6"
    },
    {
      "name": "name2",
      "status": "status6"
    },
    {
      "name": "name2",
      "status": "status6"
    }
  ],
  "fabric_link_info": {
    "DataPlaneNotifiedStatus": "DataPlaneNotifiedStatus8",
    "Interface": [
      "Interface0",
      "Interface1",
      "Interface2"
    ],
    "InternalStatus": "InternalStatus2",
    "State": "State4",
    "Status": "Status0"
  },
  "last_status_change_reason": "last_status_change_reason2"
}
```

