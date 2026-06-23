
# Network Vpn Access Destination Nat Property

VPN access destination NAT rule target settings

## Structure

`NetworkVpnAccessDestinationNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | The Destination NAT destination IP address. Must be an IP (i.e. "192.168.70.30") or a Variable (i.e. "{{myvar}}") |
| `Name` | `*string` | Optional | Label for this VPN destination NAT rule |
| `Port` | `*string` | Optional | Destination port or variable for this VPN destination NAT rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkVpnAccessDestinationNatProperty := models.NetworkVpnAccessDestinationNatProperty{
        InternalIp:           models.ToPointer("192.168.70.30"),
        Name:                 models.ToPointer("web server"),
        Port:                 models.ToPointer("443"),
    }

}
```

