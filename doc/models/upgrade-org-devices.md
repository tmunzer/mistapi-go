
# Upgrade Org Devices

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeOrgDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CanaryPhases` | `[]int` | Optional | For APs only and if `strategy`==`canary`. Phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. |
| `EnableP2p` | `*bool` | Optional | for APs only. whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | true will force upgrade when requested version is same as running version<br>**Default**: `false` |
| `MaxFailurePercentage` | `*float64` | Optional | for APs only and if `strategy`!=`big_bang`. percentage of failures allowed across the entire upgrade<br>**Default**: `5`<br>**Constraints**: `>= 0`, `<= 100` |
| `MaxFailures` | `[]int` | Optional | For APs only and if `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used |
| `Models` | `[]string` | Optional | models which will be selected for upgrade |
| `P2pClusterSize` | `*int` | Optional | For APs only and if `enable_p2p`==`true`.<br>**Default**: `10`<br>**Constraints**: `>= 0` |
| `P2pParallelism` | `*int` | Optional | For APs only and if `enable_p2p`==`true`. Number of parallel p2p download batches to create |
| `Reboot` | `*bool` | Optional | For Junos devices only (APs are automatically rebooted). Reboot device immediately after upgrade is completed<br>**Default**: `false` |
| `RebootAt` | `*float64` | Optional | For Junos devices only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time` |
| `RrmFirstBatchPercentage` | `*int` | Optional | For APs only and if `strategy`==`rrm`. Percentage of APs that need to be present in the first RRM batch |
| `RrmMaxBatchPercentage` | `*int` | Optional | For APs only and if `strategy`==`rrm`. Max percentage of APs that need to be present in each RRM batch |
| `RrmMeshUpgrade` | [`*models.DeviceUpgradeRrmMeshUpgradeEnum`](../../doc/models/device-upgrade-rrm-mesh-upgrade-enum.md) | Optional | For APs only and if `strategy`==`rrm`. Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade. enum: `parallel`, `sequential`<br>**Default**: `"sequential"` |
| `RrmNodeOrder` | [`*models.DeviceUpgradeRrmNodeOrderEnum`](../../doc/models/device-upgrade-rrm-node-order-enum.md) | Optional | For APs only and if `strategy`==`rrm`. Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`<br>**Default**: `"fringe_to_center"` |
| `RrmSlowRamp` | `*bool` | Optional | For APs only and if `strategy`==`rrm`. True will make rrm batch sizes slowly ramp up |
| `Rules` | `[]string` | Optional | rules used to identify devices which will be selected for upgrade. Device will be selected as long as it satisfies any one rule |
| `SiteIds` | `[]uuid.UUID` | Optional | - |
| `Snapshot` | `*bool` | Optional | For Junos devices only. Perform recovery snapshot after device is rebooted<br>**Default**: `false` |
| `StartTime` | `*float64` | Optional | upgrade start time in epoch seconds, default is now |
| `Strategy` | [`*models.DeviceUpgradeStrategyEnum`](../../doc/models/device-upgrade-strategy-enum.md) | Optional | For APs only. enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `Version` | `*string` | Optional | specific version / stable, default is to use the lastest available version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "canary_phases": [
    139,
    138
  ],
  "force": false,
  "max_failure_percentage": 5.0,
  "p2p_cluster_size": 0,
  "reboot": false,
  "reboot_at": 1624399840,
  "rrm_first_batch_percentage": 2,
  "rrm_max_batch_percentage": 10,
  "rrm_mesh_upgrade": "sequential",
  "rrm_node_order": "fringe_to_center",
  "snapshot": false,
  "start_time": 1624399840,
  "strategy": "big_bang",
  "version": "3.1.5",
  "enable_p2p": false,
  "max_failures": [
    253,
    254
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

