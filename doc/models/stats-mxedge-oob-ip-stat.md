
# Stats Mxedge Oob Ip Stat

Observed out-of-band management IP settings for a Mist Edge

## Structure

`StatsMxedgeOobIpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Gateway` | `*string` | Optional | IPv4 default gateway for out-of-band management |
| `Gateway6` | `*string` | Optional | IPv6 default gateway for out-of-band management |
| `Ip` | `*string` | Optional | Out-of-band management IPv4 address |
| `Ip6` | `*string` | Optional | Out-of-band management IPv6 address |
| `Netmask` | `*string` | Optional | IPv4 netmask for the out-of-band management interface |
| `Netmask6` | `*string` | Optional | IPv6 prefix length for the out-of-band management interface |
| `Type` | [`*models.MxedgeMgmtOobIpTypeEnum`](../../doc/models/mxedge-mgmt-oob-ip-type-enum.md) | Optional | enum: `dhcp`, `disabled`, `static`<br><br>**Default**: `"dhcp"` |
| `Type8` | [`*models.MxedgeMgmtOobIpType6Enum`](../../doc/models/mxedge-mgmt-oob-ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"autoconf"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeOobIpStat := models.StatsMxedgeOobIpStat{
        Dns:                  []string{
            "dns7",
            "dns8",
        },
        Gateway:              models.ToPointer("gateway0"),
        Gateway6:             models.ToPointer("gateway66"),
        Ip:                   models.ToPointer("ip4"),
        Ip6:                  models.ToPointer("ip60"),
        Type:                 models.ToPointer(models.MxedgeMgmtOobIpTypeEnum_DHCP),
        Type8:                models.ToPointer(models.MxedgeMgmtOobIpType6Enum_AUTOCONF),
    }

}
```

