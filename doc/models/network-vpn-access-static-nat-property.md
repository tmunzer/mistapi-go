
# Network Vpn Access Static Nat Property

VPN access static NAT rule target settings

## Structure

`NetworkVpnAccessStaticNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | The Static NAT destination IP address. Must be an IP address (i.e. "192.168.70.3") or a Variable (i.e. "{{myvar}}") |
| `Name` | `*string` | Optional | Label for this VPN static NAT rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkVpnAccessStaticNatProperty := models.NetworkVpnAccessStaticNatProperty{
        InternalIp:           models.ToPointer("192.168.70.3"),
        Name:                 models.ToPointer("pos_station-1"),
    }

}
```

