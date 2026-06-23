
# Switch Oob Ip Config

Switch OOB IP Config:

- If HA configuration: key parameter will be nodeX (eg: node1)
- If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1`

## Structure

`SwitchOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | Default gateway for the out-of-band management interface when `type`==`static` |
| `Ip` | `*string` | Optional | Static IPv4 address for the out-of-band management interface when `type`==`static` |
| `Netmask` | `*string` | Optional | Used only if `subnet` is not specified in `networks` |
| `Network` | `*string` | Optional | Optional, the network to be used for mgmt |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | If supported on the platform. If enabled, DNS will be using this routing-instance, too<br><br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchOobIpConfig := models.SwitchOobIpConfig{
        Gateway:              models.ToPointer("gateway8"),
        Ip:                   models.ToPointer("ip2"),
        Netmask:              models.ToPointer("netmask8"),
        Network:              models.ToPointer("network4"),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        UseMgmtVrf:           models.ToPointer(false),
        UseMgmtVrfForHostOut: models.ToPointer(false),
    }

}
```

