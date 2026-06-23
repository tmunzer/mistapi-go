
# Upgrade Site Devices Counts

Device counts grouped by status for a site upgrade job

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

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeSiteDevicesCounts := models.UpgradeSiteDevicesCounts{
        DownloadRequested:    models.ToPointer(70),
        Downloaded:           models.ToPointer(254),
        Failed:               models.ToPointer(98),
        RebootInProgress:     models.ToPointer(236),
        Rebooted:             models.ToPointer(144),
    }

}
```

