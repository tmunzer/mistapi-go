
# Dhcpd Stat Lan

DHCP lease statistics for one network

## Structure

`DhcpdStatLan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumIps` | `*int` | Optional | Total number of IP addresses in the DHCP pool |
| `NumLeased` | `*int` | Optional | Number of DHCP pool addresses currently leased |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpdStatLan := models.DhcpdStatLan{
        NumIps:               models.ToPointer(100),
        NumLeased:            models.ToPointer(20),
    }

}
```

