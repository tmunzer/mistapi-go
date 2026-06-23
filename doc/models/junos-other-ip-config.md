
# Junos Other Ip Config

Optional switch L3 presence on an additional network or VLAN

## Structure

`JunosOtherIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EvpnAnycast` | `*bool` | Optional | For EVPN, whether anycast is desired<br><br>**Default**: `false` |
| `Ip` | `*string` | Optional | Required if `type`==`static`; IPv4 address for the additional Junos L3 presence |
| `Ip6` | `*string` | Optional | Required if `type6`==`static`; IPv6 address for the additional Junos L3 presence |
| `Netmask` | `*string` | Optional | Optional IPv4 netmask; `subnet` from `network` definition will be used if defined |
| `Netmask6` | `*string` | Optional | Optional IPv6 prefix length; `subnet` from `network` definition will be used if defined |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpType6Enum`](../../doc/models/ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"disabled"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    junosOtherIpConfig := models.JunosOtherIpConfig{
        EvpnAnycast:          models.ToPointer(false),
        Ip:                   models.ToPointer("10.3.3.1"),
        Ip6:                  models.ToPointer("fdad:b0bc:f29e::3d16"),
        Netmask:              models.ToPointer("255.255.255.0"),
        Netmask6:             models.ToPointer("/64"),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        Type6:                models.ToPointer(models.IpType6Enum_STATIC),
    }

}
```

