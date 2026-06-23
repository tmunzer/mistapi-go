
# Stats Switch Vc Setup Info

Virtual Chassis setup request and status details reported by a switch

## Structure

`StatsSwitchVcSetupInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigType` | `*string` | Optional, Read-only | Provisioning mode used for the Virtual Chassis setup |
| `CurrentStats` | `*string` | Optional, Read-only | Status currently reported for the Virtual Chassis setup workflow |
| `ErrMissingDevIdFpc` | `*bool` | Optional, Read-only | Whether the Virtual Chassis setup is missing a device ID for an FPC member |
| `LastUpdate` | `*float64` | Optional, Read-only | Most recent update time for the Virtual Chassis setup status |
| `RequestTime` | `*float64` | Optional, Read-only | Time when the Virtual Chassis setup request was submitted |
| `RequestType` | `*string` | Optional, Read-only | Virtual Chassis setup request type |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSwitchVcSetupInfo := models.StatsSwitchVcSetupInfo{
        ConfigType:           models.ToPointer("nonprovisioned"),
        CurrentStats:         models.ToPointer("VCSETUP_WAITING"),
        ErrMissingDevIdFpc:   models.ToPointer(false),
        LastUpdate:           models.ToPointer(float64(151.08)),
        RequestTime:          models.ToPointer(float64(71.6)),
        RequestType:          models.ToPointer("vc_create"),
    }

}
```

