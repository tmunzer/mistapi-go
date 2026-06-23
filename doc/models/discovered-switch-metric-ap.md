
# Discovered Switch Metric Ap

AP attachment details included in a discovered switch metric

## Structure

`DiscoveredSwitchMetricAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostname` | `*string` | Optional | AP hostname included in the discovered switch metric |
| `Mac` | `*string` | Optional | AP MAC address included in the discovered switch metric |
| `PoeStatus` | `*bool` | Optional | Whether the upstream switch port provides PoE for this AP |
| `Port` | `*string` | Optional | Switch port name connected to this AP |
| `PortId` | `*string` | Optional | LLDP port identifier for this AP uplink |
| `PowerDraw` | `*int` | Optional | Power draw reported for this AP connection |
| `When` | `*string` | Optional | Timestamp when this AP attachment metric was observed |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    discoveredSwitchMetricAp := models.DiscoveredSwitchMetricAp{
        Hostname:             models.ToPointer("hostname2"),
        Mac:                  models.ToPointer("mac6"),
        PoeStatus:            models.ToPointer(false),
        Port:                 models.ToPointer("port2"),
        PortId:               models.ToPointer("port_id2"),
    }

}
```

