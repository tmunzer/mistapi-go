
# Response Site Device Upgrade

## Structure

`ResponseSiteDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Counts` | [`*models.ResponseSiteDeviceUpgradeCounts`](../../doc/models/response-site-device-upgrade-counts.md) | Optional | - |
| `CurrentPhase` | `*int` | Optional | current canary or rrm phase in progress |
| `EnableP2p` | `*bool` | Optional | whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | whether to force upgrade when requested version is same as running version |
| `Id` | `uuid.UUID` | Required | unique id for the upgrade |
| `MaxFailurePercentage` | `*int` | Optional | percentage of failures allowed |
| `MaxFailures` | `[]int` | Optional | number of failures allowed within a canary phase or serial rollout |
| `RebootAt` | `*int` | Optional | reboot start time in epoch |
| `StartTime` | `*float64` | Optional | firmware download start time in epoch |
| `Status` | [`*models.DeviceUpgradeStatusEnum`](../../doc/models/device-upgrade-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading` |
| `Strategy` | [`*models.DeviceUpgradeStrategyEnum`](../../doc/models/device-upgrade-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | version to upgrade to<br>**Constraints**: *Minimum Length*: `1` |
| `UpgradePlan` | `*interface{}` | Optional | a dictionary of rrm phase number to devices part of that phase |

## Example (as JSON)

```json
{
  "id": "00000af6-0000-0000-0000-000000000000",
  "strategy": "big_bang",
  "counts": {
    "download_requested": 138,
    "downloaded": 70,
    "failed": 166,
    "reboot_in_progress": 88,
    "rebooted": 76
  },
  "current_phase": 156,
  "enable_p2p": false,
  "force": false,
  "max_failure_percentage": 22
}
```

