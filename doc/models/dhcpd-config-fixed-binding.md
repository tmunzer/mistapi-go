
# Dhcpd Config Fixed Binding

Static DHCP binding for a client MAC address

## Structure

`DhcpdConfigFixedBinding`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | Reserved IPv4 address for this fixed DHCP binding |
| `Ip6` | `*string` | Optional | Reserved IPv6 address for this fixed DHCP binding |
| `Name` | `*string` | Optional | Friendly name for this fixed DHCP binding |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpdConfigFixedBinding := models.DhcpdConfigFixedBinding{
        Ip:                   models.ToPointer("192.168.70.35"),
        Ip6:                  models.ToPointer("2607:f8b0:4005:808::2"),
        Name:                 models.ToPointer("name6"),
    }

}
```

