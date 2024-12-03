
# Upgrade Org Device Targets

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "download_requested": [
    "download_requested0",
    "download_requested1",
    "download_requested2"
  ],
  "downloaded": [
    "downloaded4"
  ],
  "failed": [
    "failed0",
    "failed1",
    "failed2"
  ],
  "reboot_in_progress": [
    "reboot_in_progress3",
    "reboot_in_progress2",
    "reboot_in_progress1"
  ],
  "rebooted": [
    "rebooted9",
    "rebooted0"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

