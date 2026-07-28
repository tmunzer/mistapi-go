
# Switch Multicast Config

Multicast configuration for a VRF. When set at the network template level it applies to networks in the master VRF (not assigned to any vrf_instances). PIM is automatically enabled when any network in the VRF has `multicast.enabled`==`true`.

## Structure

`SwitchMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AnycastRp` | `*bool` | Optional | When `true`, auto-generates a shared RP on `is_l3_border` devices (ERB/IPClos topologies only)<br><br>**Default**: `false` |
| `RpIp` | `*string` | Optional | RP address used when `anycast_rp`==`false`. If the address matches a device SVI, it is configured as a local RP; otherwise a static RP is configured |
| `SbdSubnet` | `*string` | Optional | SBD IRB subnet; Mist auto-assigns per-device IPs from this range (EVPN eOISM only) |
| `SbdVlanId` | `*int` | Optional | Supplemental Bridge Domain VLAN ID (EVPN topology / eOISM only) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMulticastConfig := models.SwitchMulticastConfig{
        AnycastRp:            models.ToPointer(false),
        RpIp:                 models.ToPointer("10.2.1.10"),
        SbdSubnet:            models.ToPointer("10.99.0.0/24"),
        SbdVlanId:            models.ToPointer(3900),
    }

}
```

