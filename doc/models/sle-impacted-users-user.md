
# Sle Impacted Users User

SLE impact row for a user

## Structure

`SleImpactedUsersUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | MAC address of the AP associated with this impacted user |
| `ApName` | `*string` | Optional | Display name of the AP associated with this impacted user |
| `Degraded` | `*float64` | Optional | Portion of the SLE total that was degraded for this user |
| `DeviceOs` | `*string` | Optional | Client device OS for this impacted user |
| `DeviceType` | `*string` | Optional | Client device type for this impacted user |
| `Duration` | `*float64` | Optional | Observation time represented by this user impact row |
| `Mac` | `*string` | Optional | Client MAC address for this impacted user |
| `Name` | `*string` | Optional | Display name for the user impact row |
| `Ssid` | `*string` | Optional | Wireless network SSID used by this impacted user |
| `Total` | `*float64` | Optional | Overall SLE total measured for this user impact row |
| `WlanId` | `*string` | Optional | Identifier of the WLAN used by this impacted user |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedUsersUser := models.SleImpactedUsersUser{
        ApMac:                models.ToPointer("ap_mac2"),
        ApName:               models.ToPointer("ap_name2"),
        Degraded:             models.ToPointer(float64(227.5)),
        DeviceOs:             models.ToPointer("device_os0"),
        DeviceType:           models.ToPointer("device_type0"),
    }

}
```

