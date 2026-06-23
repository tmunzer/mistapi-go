
# Stats Gateway Service Status

Gateway security service installation and runtime status

## Structure

`StatsGatewayServiceStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppidInstallResult` | `*string` | Optional | Installation result reported for the AppID package |
| `AppidInstallTimestamp` | `*string` | Optional | Installation time reported for the AppID package |
| `AppidStatus` | `*string` | Optional | Operational status of the AppID service on the gateway |
| `AppidVersion` | `*int` | Optional | Installed AppID package version number |
| `EwfStatus` | `*string` | Optional | Enhanced Web Filtering service status reported by the gateway |
| `IdpInstallResult` | `*string` | Optional | Installation result reported for the intrusion detection and prevention package |
| `IdpInstallTimestamp` | `*string` | Optional | Installation time reported for the intrusion detection and prevention package |
| `IdpPolicy` | `*string` | Optional | Intrusion detection and prevention policy applied to the gateway |
| `IdpStatus` | `*string` | Optional | Intrusion detection and prevention service status reported by the gateway |
| `IdpUpdateTimestamp` | `*string` | Optional | Last update time reported for the intrusion detection and prevention package |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsGatewayServiceStatus := models.StatsGatewayServiceStatus{
        AppidInstallResult:    models.ToPointer("appid_install_result4"),
        AppidInstallTimestamp: models.ToPointer("appid_install_timestamp4"),
        AppidStatus:           models.ToPointer("appid_status0"),
        AppidVersion:          models.ToPointer(182),
        EwfStatus:             models.ToPointer("ewf_status6"),
    }

}
```

