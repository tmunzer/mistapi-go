
# Mxedge Oob Ip Config

IP configuration for the Mist Edge out-of-band management interface

## Structure

`MxedgeOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Autoconf6` | `*bool` | Optional | Whether IPv6 autoconfiguration is enabled on the out-of-band management interface<br><br>**Default**: `true` |
| `Dhcp6` | `*bool` | Optional | Whether DHCPv6 is enabled on the out-of-band management interface<br><br>**Default**: `true` |
| `Dns` | `[]string` | Optional | IPv4 ignored if `type`!=`static`, IPv6 ignored if `type6`!=`static` |
| `Gateway` | `*string` | Optional | If `type`=`static`, IPv4 default gateway for the out-of-band management interface |
| `Gateway6` | `*string` | Optional | If `type6`=`static`, IPv6 default gateway for the out-of-band management interface |
| `Ip` | `*string` | Optional | If `type`=`static`, IPv4 address for the out-of-band management interface |
| `Ip6` | `*string` | Optional | If `type6`=`static`, IPv6 address for the out-of-band management interface |
| `Netmask` | `*string` | Optional | If `type`=`static`, IPv4 netmask for the out-of-band management interface |
| `Netmask6` | `*string` | Optional | If `type6`=`static`, IPv6 prefix length for the out-of-band management interface |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeOobIpConfig := models.MxedgeOobIpConfig{
        Autoconf6:            models.ToPointer(true),
        Dhcp6:                models.ToPointer(true),
        Dns:                  []string{
            "8.8.8.8",
            "4.4.4.4",
            "2001:4860:4860::8888",
            "2001:4860:4860::8844",
        },
        Gateway:              models.ToPointer("10.2.1.254"),
        Gateway6:             models.ToPointer("2601:1700:43c0:dc0::1"),
        Ip:                   models.ToPointer("10.2.1.2"),
        Ip6:                  models.ToPointer("2601:1700:43c0:dc0:20c:29ff:fea7:93bc"),
        Netmask:              models.ToPointer("255.255.255.0"),
        Netmask6:             models.ToPointer("/64"),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        Type6:                models.ToPointer(models.IpTypeEnum_STATIC),
    }

}
```

