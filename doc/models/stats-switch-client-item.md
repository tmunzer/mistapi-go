
# Stats Switch Client Item

Client observed on a switch port in switch statistics

## Structure

`StatsSwitchClientItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | Switch MAC address for the device reporting this client entry |
| `Hostname` | `*string` | Optional | Reported client hostname, when known |
| `Mac` | `*string` | Optional | Client MAC address observed on the switch port |
| `PortId` | `*string` | Optional | Switch port identifier where the client was observed |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSwitchClientItem := models.StatsSwitchClientItem{
        DeviceMac:            models.ToPointer("device_mac8"),
        Hostname:             models.ToPointer("hostname0"),
        Mac:                  models.ToPointer("mac8"),
        PortId:               models.ToPointer("port_id4"),
    }

}
```

