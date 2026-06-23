
# Upgrade Devices Targets

Read-only device target lists grouped by upgrade status

## Structure

`UpgradeDevicesTargets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownloadRequested` | `[]string` | Optional | List of device MAC addresses that the cloud requested to download firmware |
| `Downloaded` | `[]string` | Optional | List of device MAC addresses that have downloaded the firmware |
| `Downloading` | `[]string` | Optional | List of device MAC addresses that are currently downloading the firmware |
| `Failed` | `[]string` | Optional | List of device MAC addresses that failed to upgrade |
| `RebootInProgress` | `[]string` | Optional | List of device MAC addresses for devices that are rebooting |
| `Rebooted` | `[]string` | Optional | List of device MAC addresses that rebooted successfully |
| `Scheduled` | `[]string` | Optional | List of device MAC addresses that the cloud scheduled an upgrade for |
| `Skipped` | `[]string` | Optional | List of device MAC addresses that skipped upgrade because requested version was same as running version. Use force to always upgrade |
| `Total` | `*int` | Optional | Count of devices part of this upgrade |
| `Upgraded` | `[]string` | Optional | List of device MAC addresses that upgraded successfully |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeDevicesTargets := models.UpgradeDevicesTargets{
        DownloadRequested:    []string{
            "download_requested6",
            "download_requested5",
            "download_requested4",
        },
        Downloaded:           []string{
            "downloaded0",
        },
        Downloading:          []string{
            "downloading6",
            "downloading7",
        },
        Failed:               []string{
            "failed6",
            "failed5",
            "failed4",
        },
        RebootInProgress:     []string{
            "reboot_in_progress3",
            "reboot_in_progress4",
            "reboot_in_progress5",
        },
    }

}
```

