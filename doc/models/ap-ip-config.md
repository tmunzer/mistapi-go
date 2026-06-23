
# Ap Ip Config

Management IP addressing settings for an access point

## Structure

`ApIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | If `type`==`static`. DNS server IP addresses for AP management traffic |
| `DnsSuffix` | `[]string` | Optional | Required if `type`==`static`; DNS search suffixes applied to AP management lookups |
| `Gateway` | `*string` | Optional | Required if `type`==`static`. IPv4 default gateway for AP management traffic |
| `Gateway6` | `*string` | Optional | Required if `type6`==`static`. IPv6 default gateway for AP management traffic when static IPv6 addressing is used |
| `Ip` | `*string` | Optional | Required if `type`==`static`. Static IPv4 address for the AP management interface |
| `Ip6` | `*string` | Optional | Required if `type6`==`static`. Static IPv6 address for the AP management interface |
| `Mtu` | `*int` | Optional | Maximum transmission unit for AP management traffic |
| `Netmask` | `*string` | Optional | Required if `type`==`static`. IPv4 netmask for the AP management interface |
| `Netmask6` | `*string` | Optional | Required if `type6`==`static`. IPv6 prefix length for the AP management interface |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpType6Enum`](../../doc/models/ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"disabled"` |
| `VlanId` | `*int` | Optional | Management VLAN ID, default is 1 (untagged)<br><br>**Default**: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apIpConfig := models.ApIpConfig{
        Dns:                  []string{
            "8.8.8.8",
            "4.4.4.4",
        },
        DnsSuffix:            []string{
            ".mist.local",
            ".mist.com",
        },
        Gateway:              models.ToPointer("10.2.1.254"),
        Gateway6:             models.ToPointer("2607:f8b0:4005:808::1"),
        Ip:                   models.ToPointer("10.2.1.1"),
        Ip6:                  models.ToPointer("2607:f8b0:4005:808::2004"),
        Mtu:                  models.ToPointer(0),
        Netmask:              models.ToPointer("255.255.255.0"),
        Netmask6:             models.ToPointer("/32"),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        Type6:                models.ToPointer(models.IpType6Enum_STATIC),
        VlanId:               models.ToPointer(1),
    }

}
```

