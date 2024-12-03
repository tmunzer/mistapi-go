
# Upgrade Org Devices

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeOrgDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CanaryPhases` | `[]int` | Optional | phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. |
| `EnableP2p` | `*bool` | Optional | whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | true will force upgrade when requested version is same as running version<br>**Default**: `false` |
| `MaxFailurePercentage` | `*float64` | Optional | percentage of failures allowed across the entire upgrade(not applicable for `big_bang`)<br>**Default**: `5`<br>**Constraints**: `>= 0`, `<= 100` |
| `MaxFailures` | `[]int` | Optional | number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used |
| `Models` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `P2pClusterSize` | `*int` | Optional | **Default**: `10`<br>**Constraints**: `>= 0` |
| `P2pParallelism` | `*int` | Optional | number of parallel p2p download batches to creat |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed (Available on Junos OS devices)<br>**Default**: `false` |
| `RebootAt` | `*float64` | Optional | reboot start time in epoch seconds, default is `start_time` |
| `RrmFirstBatchPercentage` | `*int` | Optional | percentage of AP’s that need to be present in the first rrm batch |
| `RrmMaxBatchPercentage` | `*int` | Optional | max percentage of AP’s that need to be present in each rrm batch |
| `RrmMeshUpgrade` | [`*models.DeviceUpgradeRrmMeshUpgradeEnum`](../../doc/models/device-upgrade-rrm-mesh-upgrade-enum.md) | Optional | Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade. enum: `parallel`, `sequential`<br>**Default**: `"sequential"` |
| `RrmNodeOrder` | [`*models.DeviceUpgradeRrmNodeOrderEnum`](../../doc/models/device-upgrade-rrm-node-order-enum.md) | Optional | Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`<br>**Default**: `"fringe_to_center"` |
| `RrmSlowRamp` | `*bool` | Optional | true will make rrm batch sizes slowly ramp up |
| `SiteIds` | `[]uuid.UUID` | Optional | - |
| `Snapshot` | `*bool` | Optional | Perform recovery snapshot after device is rebooted (Available on Junos OS devices)<br>**Default**: `false` |
| `StartTime` | `*float64` | Optional | upgrade start time in epoch seconds, default is now |
| `Strategy` | [`*models.DeviceUpgradeStrategyEnum`](../../doc/models/device-upgrade-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `Version` | `*string` | Optional | specific version / stable<br>**Default**: `"latest"` |
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

