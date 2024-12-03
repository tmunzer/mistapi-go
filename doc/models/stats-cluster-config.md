
# Stats Cluster Config

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "configuration": "configuration4",
  "control_link_info": {
    "name": "name0",
    "status": "status2",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "ethernet_connection": [
    {
      "name": "name2",
      "status": "status6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "name": "name2",
      "status": "status6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "name": "name2",
      "status": "status6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
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
    "Status": "Status0",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "last_status_change_reason": "last_status_change_reason2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

