
# Upgrade Org Device Targets

## Structure

`UpgradeOrgDeviceTargets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownloadRequested` | `[]string` | Optional | list of devices MAC Addresses which cloud has requested to download firmware |
| `Downloaded` | `[]string` | Optional | list of devices MAC Addresses which have the firmware downloaded |
| `Failed` | `[]string` | Optional | list of devices MAC Addresses which have failed to upgrade |
| `RebootInProgress` | `[]string` | Optional | list of devices MAC Addresses which are rebooting |
| `Rebooted` | `[]string` | Optional | list of devices MAC Addresses which have rebooted successfully |
| `Scheduled` | `[]string` | Optional | list of devices MAC Addresses which cloud has scheduled an upgrade for |
| `Skipped` | `[]string` | Optional | list of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade |
| `Total` | `*int` | Optional | count of devices part of this upgrade |
| `Upgraded` | `[]string` | Optional | count of devices which have upgraded successfully |

## Example (as JSON)

```json
{
  "download_requested": [
    "download_requested0",
    "download_requested1"
  ],
  "downloaded": [
    "downloaded4",
    "downloaded5"
  ],
  "failed": [
    "failed0",
    "failed1"
  ],
  "reboot_in_progress": [
    "reboot_in_progress3"
  ],
  "rebooted": [
    "rebooted9",
    "rebooted0",
    "rebooted1"
  ]
}
```
