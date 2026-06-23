
# Gateway Oob Ip Config

Out-of-band management IP configuration for gateway interfaces such as vme, em0, or fxp0

## Structure

`GatewayOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | Default gateway for the out-of-band management interface when `type`==`static` |
| `Ip` | `*string` | Optional | Static IPv4 address for the out-of-band management interface when `type`==`static` |
| `Netmask` | `*string` | Optional | IPv4 netmask or prefix length for the out-of-band management interface when `type`==`static` |
| `Node1` | [`*models.GatewayOobIpConfigNode1`](../../doc/models/gateway-oob-ip-config-node-1.md) | Optional | Node1-specific out-of-band management IP configuration for HA clusters |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | If supported on the platform. If enabled, DNS will be using this routing-instance, too<br><br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br><br>**Default**: `false` |
| `VlanId` | [`*models.GatewayPortVlanIdWithVariable`](../../doc/models/containers/gateway-port-vlan-id-with-variable.md) | Optional | If WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayOobIpConfig := models.GatewayOobIpConfig{
        Gateway:              models.ToPointer("gateway2"),
        Ip:                   models.ToPointer("ip6"),
        Netmask:              models.ToPointer("netmask2"),
        Node1:                models.ToPointer(models.GatewayOobIpConfigNode1{
            Gateway:              models.ToPointer("gateway2"),
            Ip:                   models.ToPointer("ip6"),
            Netmask:              models.ToPointer("netmask2"),
            Type:                 models.ToPointer(models.IpTypeEnum_DHCP),
            UseMgmtVrf:           models.ToPointer(false),
        }),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        UseMgmtVrf:           models.ToPointer(false),
        UseMgmtVrfForHostOut: models.ToPointer(false),
    }

}
```

