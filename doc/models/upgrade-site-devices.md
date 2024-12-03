
# Upgrade Site Devices

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeSiteDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CanaryPhases` | `[]int` | Optional | phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. default is [1, 10, 50, 100] |
| `DeviceIds` | `[]uuid.UUID` | Optional | - |
| `EnableP2p` | `*bool` | Optional | whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | true will force upgrade when requested version is same as running version<br>**Default**: `false` |
| `MaxFailurePercentage` | `*float64` | Optional | percentage of failures allowed across the entire upgrade(not applicable for `big_bang`)<br>**Default**: `5`<br>**Constraints**: `>= 0`, `<= 100` |
| `MaxFailures` | `[]int` | Optional | number of failures allowed within each phase(applicable for `canary` or `rrm`). Will be used if provided, else max_failure_percentage will be used |
| `Models` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `P2pClusterSize` | `*int` | Optional | **Default**: `10`<br>**Constraints**: `>= 0` |
| `P2pParallelism` | `*int` | Optional | number of parallel p2p download batches to creat |
| `Reboot` | `*bool` | Optional | Reboot device immediately after upgrade is completed (Available on Junos OS devices)<br>**Default**: `false` |
| `RebootAt` | `*float64` | Optional | reboot start time in epoch seconds, default is `start_time` |
| `RrmFirstBatchPercentage` | `*int` | Optional | percentage of AP’s that need to be present in the first rrm batch |
| `RrmMaxBatchPercentage` | `*int` | Optional | max percentage of AP’s that need to be present in each rrm batch |
| `RrmMeshUpgrade` | `*string` | Optional | sequential or parallel (default parallel). Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade |
| `RrmNodeOrder` | [`*models.UpgradeSiteDevicesRrmNodeOrderEnum`](../../doc/models/upgrade-site-devices-rrm-node-order-enum.md) | Optional | Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`<br>**Default**: `"fringe_to_center"` |
| `RrmSlowRamp` | `*bool` | Optional | true will make rrm batch sizes slowly ramp up |
| `Snapshot` | `*bool` | Optional | Perform recovery snapshot after device is rebooted (Available on Junos OS devices)<br>**Default**: `false` |
| `StartTime` | `*float64` | Optional | upgrade start time in epoch seconds, default is now |
| `Strategy` | [`*models.DeviceUpgradeStrategyEnum`](../../doc/models/device-upgrade-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `Version` | `*string` | Optional | specific version / stable<br>**Default**: `"latest"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "canary_phases": [
    135,
    136
  ],
  "force": false,
  "max_failure_percentage": 5.0,
  "max_failures": [
    1,
    1,
    5,
    5
  ],
  "p2p_cluster_size": 0,
  "reboot": false,
  "reboot_at": 1624399840,
  "rrm_first_batch_percentage": 2,
  "rrm_max_batch_percentage": 10,
  "rrm_node_order": "fringe_to_center",
  "snapshot": false,
  "start_time": 1624399840,
  "strategy": "big_bang",
  "version": "3.1.5",
  "device_ids": [
    "000019b7-0000-0000-0000-000000000000"
  ],
  "enable_p2p": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

