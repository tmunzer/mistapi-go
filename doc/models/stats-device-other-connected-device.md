
# Stats Device Other Connected Device

LLDP neighbor device discovered from an other-device statistics record

## Structure

`StatsDeviceOtherConnectedDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Connected device MAC address learned from LLDP |
| `Name` | `*string` | Optional | LLDP system name reported for the connected device |
| `PortId` | `*string` | Optional | LLDP port identifier reported for the connected device |
| `Type` | `*string` | Optional | Detected Mist device type for the connected neighbor |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsDeviceOtherConnectedDevice := models.StatsDeviceOtherConnectedDevice{
        Mac:                  models.ToPointer("020001abcdef"),
        Name:                 models.ToPointer("DNT-NTR-GWE"),
        PortId:               models.ToPointer("ge-0/0/1"),
        Type:                 models.ToPointer("gateway"),
    }

}
```

