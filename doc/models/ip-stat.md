
# Ip Stat

Read-only IP addressing status reported by a device interface

## Structure

`IpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DhcpServer` | `models.Optional[string]` | Optional, Read-only | Reported DHCP server IPv4 address for the interface, when available |
| `Dns` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `DnsSuffix` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Gateway` | `models.Optional[string]` | Optional, Read-only | Current IPv4 default gateway for the interface, when available |
| `Gateway6` | `models.Optional[string]` | Optional, Read-only | Current IPv6 default gateway for the interface, when available |
| `Ip` | `models.Optional[string]` | Optional, Read-only | Current IPv4 address for the interface, when available |
| `Ip6` | `models.Optional[string]` | Optional, Read-only | Current IPv6 address for the interface, when available |
| `Ips` | `map[string]string` | Optional | Per-VLAN IP address summaries keyed by VLAN name |
| `Netmask` | `models.Optional[string]` | Optional, Read-only | Current IPv4 subnet mask for the interface, when available |
| `Netmask6` | `models.Optional[string]` | Optional, Read-only | Current IPv6 prefix length for the interface, when available |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ipStat := models.IpStat{
        DhcpServer:           models.NewOptional(models.ToPointer("192.168.95.1")),
        Dns:                  []string{
            "dns7",
            "dns8",
        },
        DnsSuffix:            []string{
            "dns_suffix9",
            "dns_suffix0",
        },
        Gateway:              models.NewOptional(models.ToPointer("gateway0")),
        Gateway6:             models.NewOptional(models.ToPointer("fdad:b0bc:f29e::1")),
        Ip:                   models.NewOptional(models.ToPointer("10.3.3.1")),
        Ip6:                  models.NewOptional(models.ToPointer("fdad:b0bc:f29e::3d16")),
        Netmask:              models.NewOptional(models.ToPointer("255.255.255.0")),
        Netmask6:             models.NewOptional(models.ToPointer("/64")),
    }

}
```

