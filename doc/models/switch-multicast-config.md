
# Switch Multicast Config

Multicast configuration for a VRF. When set at the network template level it applies to networks in the master VRF (not assigned to any vrf_instances). PIM is automatically enabled when any network in the VRF has `multicast.enabled`==`true`.

## Structure

`SwitchMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AnycastRp` | `*bool` | Optional | When `true`, generates a shared anycast RP on all `is_l3_border` devices in EVPN (ERB/IPClos) topologies. Uses `rp_ip` as the shared RP address, or an internal default when `rp_ip` is omitted. Takes precedence over `rp_mac` and `rp_ip` when multiple RP options are set.<br><br>**Default**: `false` |
| `PegEnabled` | `*bool` | Optional | When `true`, enables the PIM EVPN Gateway on `is_l3_border` devices. Required for external sources or receivers in EVPN topologies.<br><br>**Default**: `false` |
| `RpIp` | `*string` | Optional | RP address used for EVPN anycast RP when `anycast_rp` is true, or for an external RP when it is false. In non-EVPN topologies, a matching device router ID configures a local RP; otherwise a static RP is configured. |
| `RpMac` | `*string` | Optional | Device MAC address of a fabric RP in EVPN topologies. The RP address is the first usable IP of the VRF `evpn_auto_loopback_subnet`, not `rp_ip`; requires `evpn_auto_loopback_subnet`. Takes precedence over `rp_ip` when `anycast_rp` is false. |
| `SbdSubnet` | `*string` | Optional | SBD IRB subnet; Mist auto-assigns per-device IPs from this range (EVPN eOISM only) |
| `SbdVlanId` | `*int` | Optional | Supplemental Bridge Domain VLAN ID (EVPN topology / eOISM only) |
| `SbdWanRpf` | `*bool` | Optional | When `true` on PEG borders, builds an eBGP mesh between PEG borders over SBD IRBs so WAN-learned routes can satisfy the PIM RPF check during a border WAN-uplink failure.<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMulticastConfig := models.SwitchMulticastConfig{
        AnycastRp:            models.ToPointer(false),
        PegEnabled:           models.ToPointer(false),
        RpIp:                 models.ToPointer("10.2.1.10"),
        RpMac:                models.ToPointer("5f263135a66d"),
        SbdSubnet:            models.ToPointer("10.99.0.0/24"),
        SbdVlanId:            models.ToPointer(3900),
        SbdWanRpf:            models.ToPointer(false),
    }

}
```

