
# Response Device Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Counts` | [`*models.ResponseDeviceUpgradeCounts`](../../doc/models/response-device-upgrade-counts.md) | Optional | - |
| `CurrentPhase` | `*int` | Optional | current canary or rrm phase in progress |
| `EnableP2p` | `*bool` | Optional | whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | whether to force upgrade when requested version is same as running version |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `MaxFailurePercentage` | `*int` | Optional | percentage of failures allowed |
| `MaxFailures` | `[]int` | Optional | number of failures allowed within a canary phase or serial rollout |
| `RebootAt` | `*int` | Optional | reboot start time in epoch |
| `StartTime` | `*float64` | Optional | firmware download start time in epoch |
| `Status` | [`*models.DeviceUpgradeStatusEnum`](../../doc/models/device-upgrade-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading` |
| `Strategy` | [`*models.DeviceUpgradeStrategyEnum`](../../doc/models/device-upgrade-strategy-enum.md) | Optional | For APs only. enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | version to upgrade to<br>**Constraints**: *Minimum Length*: `1` |
| `UpgradePlan` | `*interface{}` | Optional | a dictionary of rrm phase number to devices part of that phase |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "strategy": "big_bang",
  "counts": {
    "download_requested": [
      "download_requested2",
      "download_requested3",
      "download_requested4"
    ],
    "downloaded": [
      "downloaded0"
    ],
    "failed": [
      "failed2",
      "failed3",
      "failed4"
    ],
    "reboot_in_progress": [
      "reboot_in_progress9",
      "reboot_in_progress0"
    ],
    "rebooted": [
      "rebooted7",
      "rebooted8"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "current_phase": 168,
  "enable_p2p": false,
  "force": false,
  "max_failure_percentage": 34,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

