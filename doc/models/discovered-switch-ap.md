
# Discovered Switch Ap

AP attachment details observed on a discovered switch

## Structure

`DiscoveredSwitchAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostname` | `*string` | Optional | AP hostname observed on the discovered switch |
| `InactiveWiredVlans` | `[]int` | Optional | List of integer values |
| `Mac` | `*string` | Optional | AP MAC address observed on the discovered switch |
| `PoeStatus` | `*bool` | Optional | Whether the upstream switch port provides PoE for the AP |
| `Port` | `*string` | Optional | Switch port name connected to the AP |
| `PortId` | `*string` | Optional | LLDP port identifier for the AP uplink |
| `PowerDraw` | `*float64` | Optional | Power draw reported for the AP connection |
| `When` | `*string` | Optional | Timestamp when this AP attachment was last observed |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    discoveredSwitchAp := models.DiscoveredSwitchAp{
        Hostname:             models.ToPointer("hostname0"),
        InactiveWiredVlans:   []int{
            24,
        },
        Mac:                  models.ToPointer("mac8"),
        PoeStatus:            models.ToPointer(false),
        Port:                 models.ToPointer("port4"),
    }

}
```

