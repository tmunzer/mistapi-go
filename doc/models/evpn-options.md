
# Evpn Options

EVPN topology generation options for campus fabric configuration

## Structure

`EvpnOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoLoopbackSubnet` | `*string` | Optional | Optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides<br><br>**Default**: `"172.16.192.0/24"` |
| `AutoLoopbackSubnet6` | `*string` | Optional | Optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides<br><br>**Default**: `"fd33:ab00:2::/64"` |
| `AutoRouterIdSubnet` | `*string` | Optional | Optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored<br><br>**Default**: `"172.16.254.0/23"` |
| `AutoRouterIdSubnet6` | `*string` | Optional | Optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored |
| `CoreAsBorder` | `*bool` | Optional | Optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway. When `routed_at` != `core`, whether to do virtual-gateway at core as well<br><br>**Default**: `false` |
| `EnableInbandMgmt` | `*bool` | Optional | Whether to route management traffic inband; routes will be propagated to downstream switches<br><br>**Default**: `false` |
| `EnableInbandZtp` | `*bool` | Optional | if the mangement traffic goes inbnd, during installation, only the border/core switches are connected to the Internet to allow initial configuration to be pushed down and leave the downstream access switches stay in the Factory Default state enabling inband-ztp allows upstream switches to use LLDP to assign IP and gives Internet to downstream switches in that state<br><br>**Default**: `false` |
| `Overlay` | [`*models.EvpnOptionsOverlay`](../../doc/models/evpn-options-overlay.md) | Optional | EVPN overlay BGP settings |
| `PerVlanVgaV4Mac` | `*bool` | Optional | Only for by Core-Distribution architecture when `evpn_options.routed_at`==`core`. By default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address's v4_mac. If enabled, 00-00-5e-00-0X-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)<br><br>**Default**: `false` |
| `PerVlanVgaV6Mac` | `*bool` | Optional | Only for by Core-Distribution architecture when `evpn_options.routed_at`==`core`. By default, JUNOS uses 00-00-5e-00-02-01 as the virtual-gateway-address's v6_mac. If enabled, 00-00-5e-00-1X-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)<br><br>**Default**: `false` |
| `RoutedAt` | [`*models.EvpnOptionsRoutedAtEnum`](../../doc/models/evpn-options-routed-at-enum.md) | Optional | optional, where virtual-gateway should reside. enum: `core`, `distribution`, `edge`<br><br>**Default**: `"edge"` |
| `Underlay` | [`*models.EvpnOptionsUnderlay`](../../doc/models/evpn-options-underlay.md) | Optional | EVPN underlay BGP and subnet settings |
| `VsInstances` | [`map[string]models.EvpnOptionsVsInstance`](../../doc/models/evpn-options-vs-instance.md) | Optional | Optional, for EX9200 only to segregate virtual-switches |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnOptions := models.EvpnOptions{
        AutoLoopbackSubnet:   models.ToPointer("172.16.192.0/24"),
        AutoLoopbackSubnet6:  models.ToPointer("fd33:ab00:2::/64"),
        AutoRouterIdSubnet:   models.ToPointer("172.16.254.0/23"),
        AutoRouterIdSubnet6:  models.ToPointer("fd31:5700:1::/64"),
        CoreAsBorder:         models.ToPointer(false),
        EnableInbandMgmt:     models.ToPointer(false),
        EnableInbandZtp:      models.ToPointer(false),
        PerVlanVgaV4Mac:      models.ToPointer(false),
        PerVlanVgaV6Mac:      models.ToPointer(false),
        RoutedAt:             models.ToPointer(models.EvpnOptionsRoutedAtEnum_EDGE),
        VsInstances:          map[string]models.EvpnOptionsVsInstance{
            "guest": models.EvpnOptionsVsInstance{
                Networks:             []string{
                    "guest",
                },
            },
            "iot": models.EvpnOptionsVsInstance{
                Networks:             []string{
                    "iot-wifi",
                    "iot-lan",
                },
            },
        },
    }

}
```

