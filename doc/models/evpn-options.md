
# Evpn Options

EVPN Options

## Structure

`EvpnOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoLoopbackSubnet` | `*string` | Optional | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides |
| `AutoLoopbackSubnet6` | `*string` | Optional | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides |
| `AutoRouterIdSubnet` | `*string` | Optional | optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored |
| `AutoRouterIdSubnet6` | `*string` | Optional | optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored<br>**Default**: `"fd31:5700:1::/64"` |
| `CoreAsBorder` | `*bool` | Optional | optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway<br>when `routed_at` != `core`, whether to do virtual-gateway at core as well<br>**Default**: `false` |
| `Overlay` | [`*models.EvpnOptionsOverlay`](../../doc/models/evpn-options-overlay.md) | Optional | - |
| `PerVlanVgaV4Mac` | `*bool` | Optional | by default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address's v4_mac<br>if enabled, 00-00-5e-00-XX-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)<br>**Default**: `false` |
| `RoutedAt` | [`*models.EvpnOptionsRoutedAtEnum`](../../doc/models/evpn-options-routed-at-enum.md) | Optional | optional, where virtual-gateway should reside. enum: `core`, `distribution`, `edge`<br>**Default**: `"edge"` |
| `Underlay` | [`*models.EvpnOptionsUnderlay`](../../doc/models/evpn-options-underlay.md) | Optional | - |
| `VsInstances` | [`map[string]models.EvpnOptionsVsInstance`](../../doc/models/evpn-options-vs-instance.md) | Optional | optional, for EX9200 only to seggregate virtual-switches |

## Example (as JSON)

```json
{
  "auto_loopback_subnet": "100.101.0.0/24",
  "auto_loopback_subnet6": "fd33:ab00:2::/64",
  "auto_router_id_subnet": "100.100.0.0/24",
  "auto_router_id_subnet6": "fd31:5700:1::/64",
  "core_as_border": false,
  "per_vlan_vga_v4_mac": false,
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
  }
}
```

