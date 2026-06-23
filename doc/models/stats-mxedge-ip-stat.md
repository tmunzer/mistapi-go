
# Stats Mxedge Ip Stat

IP address statistics reported by a Mist Edge

## Structure

`StatsMxedgeIpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | Primary IPv4 address reported for the Mist Edge |
| `Ip6` | `*string` | Optional | Primary IPv6 address reported for the Mist Edge |
| `Ips` | `map[string]string` | Optional | Interface IP addresses keyed by interface name |
| `Macs` | `map[string]string` | Optional | Interface MAC addresses keyed by interface name |
| `Netmask` | `*string` | Optional | IPv4 netmask reported for the primary Mist Edge address |
| `Netmask6` | `*string` | Optional | IPv6 prefix length reported for the primary Mist Edge address |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeIpStat := models.StatsMxedgeIpStat{
        Ip:                   models.ToPointer("192.168.1.244"),
        Ip6:                  models.ToPointer("fd4e:c615:b27d:5555::45"),
        Ips:                  map[string]string{
            "ens18": "92.168.1.244/24,fd4e:c615:b27d:5555::45/128,fd4e:c615:b27d:5555:20c:29ff:fe44:af25/64,fe80::104c:ffff:fee0:caf8/64",
        },
        Macs:                 map[string]string{
            "ens18": "e4434b217044",
        },
        Netmask:              models.ToPointer("255.255.255.0"),
        Netmask6:             models.ToPointer("/128"),
    }

}
```

