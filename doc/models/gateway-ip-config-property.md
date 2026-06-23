
# Gateway Ip Config Property

Gateway network interface IP configuration

## Structure

`GatewayIpConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | Static IPv4 address for the gateway network interface when `type`==`static` |
| `Ip6` | `*string` | Optional | Static IPv6 address for the gateway network interface when `type6`==`static` |
| `Netmask` | `*string` | Optional | IPv4 netmask or prefix length for the gateway network interface when `type`==`static` |
| `Netmask6` | `*string` | Optional | IPv6 netmask or prefix length for the gateway network interface when `type6`==`static` |
| `SecondaryIps` | `[]string` | Optional | Optional list of secondary IPs in CIDR format |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpType6Enum`](../../doc/models/ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"disabled"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayIpConfigProperty := models.GatewayIpConfigProperty{
        Ip:                   models.ToPointer("ip8"),
        Ip6:                  models.ToPointer("ip64"),
        Netmask:              models.ToPointer("/24"),
        Netmask6:             models.ToPointer("2001:db8:abcd:12::1"),
        SecondaryIps:         []string{
            "192.168.50.1/24",
            "192.168.60.1/26",
        },
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        Type6:                models.ToPointer(models.IpType6Enum_STATIC),
    }

}
```

