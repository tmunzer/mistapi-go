
# Response Site Device Upgrades Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSiteDeviceUpgradesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Counts` | [`*models.UpgradeSiteDevicesCounts`](../../doc/models/upgrade-site-devices-counts.md) | Optional | - |
| `CurrentPhase` | `*int` | Optional | Current canary or rrm phase in progress |
| `EnableP2p` | `*bool` | Optional | Whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | Whether to force upgrade when requested version is same as running version |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `MaxFailurePercentage` | `*int` | Optional | Percentage of failures allowed |
| `MaxFailures` | `[]int` | Optional | If `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used |
| `RebootAt` | `*int` | Optional | reboot start time in epoch |
| `StartTime` | `*int` | Optional | Firmware download start time in epoch |
| `Status` | [`*models.UpgradeDeviceStatusEnum`](../../doc/models/upgrade-device-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued` |
| `Strategy` | [`*models.UpgradeDeviceStrategyEnum`](../../doc/models/upgrade-device-strategy-enum.md) | Optional | For APs only. enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)<br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | Version to upgrade to<br>**Constraints**: *Minimum Length*: `1` |
| `UpgradePlan` | `*interface{}` | Optional | a dictionary of rrm phase number to devices part of that phase |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "strategy": "big_bang",
  "counts": {
    "download_requested": 138,
    "downloaded": 70,
    "failed": 166,
    "reboot_in_progress": 88,
    "rebooted": 76,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "current_phase": 230,
  "enable_p2p": false,
  "force": false,
  "max_failure_percentage": 148,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

