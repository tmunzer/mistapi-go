
# Upgrade Site Devices

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeSiteDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CanaryPhases` | `[]int` | Optional | Only if `strategy`==`canary`. Phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. default is [1, 10, 50, 100] |
| `DeviceIds` | `[]uuid.UUID` | Optional | id's of devices which will be selected for upgrade |
| `EnableP2p` | `*bool` | Optional | For APs only. Whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | `force`==`true` will force upgrade when requested version is same as running version<br><br>**Default**: `false` |
| `MaxFailurePercentage` | `*int` | Optional | If `strategy`!=`big_bang`. percentage of failures allowed across the entire upgrade<br><br>**Default**: `5`<br><br>**Constraints**: `>= 0`, `<= 100` |
| `MaxFailures` | `[]int` | Optional | If `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used |
| `Models` | `[]string` | Optional | Models which will be selected for upgrade |
| `P2pClusterSize` | `*int` | Optional | For APs only and if `enable_p2p`==`true`.<br><br>**Default**: `10`<br><br>**Constraints**: `>= 0` |
| `P2pParallelism` | `*int` | Optional | For APs only and if `enable_p2p`==`true`. Number of parallel p2p download batches to create |
| `Reboot` | `*bool` | Optional | For Switches and Gateways only (APs are automatically rebooted). Reboot device immediately after upgrade is completed<br><br>**Default**: `false` |
| `RebootAt` | `*int` | Optional | For Switches and Gateways only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time` |
| `RrmFirstBatchPercentage` | `*int` | Optional | For APs only and if `strategy`==`rrm`. Percentage of APs that need to be present in the first RRM batch |
| `RrmMaxBatchPercentage` | `*int` | Optional | For APs only and if `strategy`==`rrm`. Max percentage of APs that need to be present in each RRM batch |
| `RrmMeshUpgrade` | [`*models.UpgradeDeviceRrmMeshUpgradeEnum`](../../doc/models/upgrade-device-rrm-mesh-upgrade-enum.md) | Optional | For APs only and if `strategy`==`rrm`. Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade. enum: `parallel`, `sequential`<br><br>**Default**: `"sequential"` |
| `RrmNodeOrder` | [`*models.UpgradeDeviceRrmNodeOrderEnum`](../../doc/models/upgrade-device-rrm-node-order-enum.md) | Optional | For APs only and if `strategy`==`rrm`. Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`<br><br>**Default**: `"fringe_to_center"` |
| `RrmSlowRamp` | `*bool` | Optional | For APs only and if `strategy`==`rrm`. True will make rrm batch sizes slowly ramp up |
| `Rules` | `[]map[string]string` | Optional | Rules used to identify devices which will be selected for upgrade. Device will be selected as long as it satisfies any one rule  <br>Property key defines the type of matching, value is the string to match. e.g:<br><br>* `match_name`: Device name must match the property value<br>* `match_name[0:3]`: Device name must match the first 3 letters of the property value<br>* `match_name[2:6]`: Device name must match the property value from the 2nd to the 6th letter<br>* `match_model`: Device model must match the property value<br>* `match_model[1:3]`: Device model must match the property value from the 1nd to the 3rd letter |
| `Snapshot` | `*bool` | Optional | For Junos devices only. Perform recovery snapshot after device is rebooted<br><br>**Default**: `false` |
| `StartTime` | `*int` | Optional | Upgrade start time in epoch seconds, default is now |
| `Strategy` | [`*models.UpgradeDeviceStrategyEnum`](../../doc/models/upgrade-device-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)<br><br>**Default**: `"big_bang"` |
| `Version` | `*string` | Optional | Specific version / stable, default is to use the latest available version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "canary_phases": null,
  "force": false,
  "max_failure_percentage": 5,
  "p2p_cluster_size": 0,
  "reboot": false,
  "reboot_at": 1624399840,
  "rrm_first_batch_percentage": 2,
  "rrm_max_batch_percentage": 10,
  "rrm_mesh_upgrade": "sequential",
  "rrm_node_order": "fringe_to_center",
  "rules": [
    {
      "match_model": "AP43",
      "match_name[2:8]": "access"
    },
    {
      "match_model": "AP45"
    }
  ],
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

