
# Evpn Options

EVPN Options

*This model accepts additional fields of type interface{}.*

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
| `EnableInbandZtp` | `*bool` | Optional | if the mangement traffic goes inbnd, during installation, only the border/core switches are connected to the Internet to allow initial configuration to be pushed down and leave the downstream access switches stay in the Factory Default state enabling inband-ztp allows upstream switches to use LLDP to assign IP and gives Internet to downstream switches in that state<br><br>**Default**: `false` |
| `Overlay` | [`*models.EvpnOptionsOverlay`](../../doc/models/evpn-options-overlay.md) | Optional | - |
| `PerVlanVgaV4Mac` | `*bool` | Optional | Only for by Core-Distribution architecture when `evpn_options.routed_at`==`core`. By default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address's v4_mac. If enabled, 00-00-5e-00-0X-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)<br><br>**Default**: `false` |
| `PerVlanVgaV6Mac` | `*bool` | Optional | Only for by Core-Distribution architecture when `evpn_options.routed_at`==`core`. By default, JUNOS uses 00-00-5e-00-02-01 as the virtual-gateway-address's v6_mac. If enabled, 00-00-5e-00-1X-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)<br><br>**Default**: `false` |
| `RoutedAt` | [`*models.EvpnOptionsRoutedAtEnum`](../../doc/models/evpn-options-routed-at-enum.md) | Optional | optional, where virtual-gateway should reside. enum: `core`, `distribution`, `edge`<br><br>**Default**: `"edge"` |
| `Underlay` | [`*models.EvpnOptionsUnderlay`](../../doc/models/evpn-options-underlay.md) | Optional | - |
| `VsInstances` | [`map[string]models.EvpnOptionsVsInstance`](../../doc/models/evpn-options-vs-instance.md) | Optional | Optional, for EX9200 only to segregate virtual-switches |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auto_loopback_subnet": "172.16.192.0/24",
  "auto_loopback_subnet6": "fd33:ab00:2::/64",
  "auto_router_id_subnet": "172.16.254.0/23",
  "auto_router_id_subnet6": "fd31:5700:1::/64",
  "core_as_border": false,
  "enable_inband_ztp": false,
  "per_vlan_vga_v4_mac": false,
  "per_vlan_vga_v6_mac": false,
  "routed_at": "edge",
  "vs_instances": {
    "guest": {
      "networks": [
        "guest"
      ]
    },
    "iot": {
      "networks": [
        "iot-wifi",
        "iot-lan"
      ]
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

