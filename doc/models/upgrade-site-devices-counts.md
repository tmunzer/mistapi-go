
# Upgrade Site Devices Counts

## Structure

`UpgradeSiteDevicesCounts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownloadRequested` | `*int` | Optional | Count of devices which cloud has requested to download firmware |
| `Downloaded` | `*int` | Optional | Count of ap's which have the firmware downloaded |
| `Failed` | `*int` | Optional | Count of devices which have failed to upgrade |
| `RebootInProgress` | `*int` | Optional | Count of devices which are rebooting |
| `Rebooted` | `*int` | Optional | Count of devices which have rebooted successfully |
| `Scheduled` | `*int` | Optional | Count of devices which cloud has scheduled an upgrade for |
| `Skipped` | `*int` | Optional | Count of devices which skipped upgrade since requested version was same as running version. Use force to always upgrade |
| `Total` | `*int` | Optional | Count of devices part of this upgrade |
| `Upgraded` | `*int` | Optional | Count of devices which have upgraded successfully |

## Example (as JSON)

```json
{
  "download_requested": 212,
  "downloaded": 112,
  "failed": 240,
  "reboot_in_progress": 94,
  "rebooted": 2
}
```

