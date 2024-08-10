
# Cluster Config Stats

## Structure

`ClusterConfigStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Configuration` | `*string` | Optional | - |
| `ControlLinkInfo` | [`*models.ClusterConfigStatsControlLinkInfo`](../../doc/models/cluster-config-stats-control-link-info.md) | Optional | - |
| `EthernetConnection` | [`[]models.ClusterConfigStatsEthernetConnectionItem`](../../doc/models/cluster-config-stats-ethernet-connection-item.md) | Optional | - |
| `FabricLinkInfo` | [`*models.ClusterConfigStatsFabricLinkInfo`](../../doc/models/cluster-config-stats-fabric-link-info.md) | Optional | - |
| `LastStatusChangeReason` | `*string` | Optional | - |
| `Operational` | `*string` | Optional | - |
| `PrimaryNodeHealth` | `*string` | Optional | - |
| `RedundancyGroupInformation` | [`[]models.ClusterConfigStatsRedundancyGroupInfoItem`](../../doc/models/cluster-config-stats-redundancy-group-info-item.md) | Optional | - |
| `SecondaryNodeHealth` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "configuration": "configuration6",
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
  "last_status_change_reason": "last_status_change_reason4"
}
```

