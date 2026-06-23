
# Junos Ip Config

Junos management IP configuration

## Structure

`JunosIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `DnsSuffix` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Gateway` | `*string` | Optional | Default gateway IPv4 address for this Junos IP configuration |
| `Ip` | `*string` | Optional | Configured IPv4 address for this Junos IP configuration |
| `Netmask` | `*string` | Optional | Used only if `subnet` is not specified in `networks` |
| `Network` | `*string` | Optional | Management network for this IP configuration; used as the default source network for outbound SSH, DNS, NTP, TACACS+, RADIUS, syslog, and SNMP |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    junosIpConfig := models.JunosIpConfig{
        Dns:                  []string{
            "dns5",
            "dns4",
            "dns3",
        },
        DnsSuffix:            []string{
            "dns_suffix3",
            "dns_suffix2",
            "dns_suffix1",
        },
        Gateway:              models.ToPointer("gateway8"),
        Ip:                   models.ToPointer("ip2"),
        Netmask:              models.ToPointer("netmask8"),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
    }

}
```

