
# Const Device Gateway Ports

Object Key is the interface name (e.g. "ge-0/0/1", ...)

## Structure

`ConstDeviceGatewayPorts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `*string` | Optional | User-facing interface name for the gateway port |
| `PciAddress` | `*string` | Optional | PCI address for the gateway port |
| `Speed` | `*int` | Optional | Port speed for the gateway interface, in Mbps |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceGatewayPorts := models.ConstDeviceGatewayPorts{
        Display:              models.ToPointer("display2"),
        PciAddress:           models.ToPointer("pci_address4"),
        Speed:                models.ToPointer(252),
    }

}
```

