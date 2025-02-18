
# Upgrade Devices Targets

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeDevicesTargets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownloadRequested` | `[]string` | Optional | List of devices MAC Addresses which cloud has requested to download firmware |
| `Downloaded` | `[]string` | Optional | List of devices MAC Addresses which have the firmware downloaded |
| `Downloading` | `[]string` | Optional | List of devices MAC Addresses which are currently downloading the firmware |
| `Failed` | `[]string` | Optional | List of devices MAC Addresses which have failed to upgrade |
| `RebootInProgress` | `[]string` | Optional | List of devices MAC Addresses which are rebooting |
| `Rebooted` | `[]string` | Optional | List of devices MAC Addresses which have rebooted successfully |
| `Scheduled` | `[]string` | Optional | List of devices MAC Addresses which cloud has scheduled an upgrade for |
| `Skipped` | `[]string` | Optional | List of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade |
| `Total` | `*int` | Optional | Count of devices part of this upgrade |
| `Upgraded` | `[]string` | Optional | Count of devices which have upgraded successfully |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "download_requested": [
    "download_requested6",
    "download_requested5"
  ],
  "downloaded": [
    "downloaded0",
    "downloaded1"
  ],
  "downloading": [
    "downloading6",
    "downloading7",
    "downloading8"
  ],
  "failed": [
    "failed6",
    "failed5"
  ],
  "reboot_in_progress": [
    "reboot_in_progress3"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

