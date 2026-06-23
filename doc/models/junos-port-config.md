
# Junos Port Config

Junos switch port configuration

## Structure

`JunosPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AeDisableLacp` | `*bool` | Optional | To disable LACP support for the AE interface |
| `AeIdx` | `*int` | Optional | Users could force to use the designated AE name |
| `AeLacpForceUp` | `*bool` | Optional | If `aggregated`==`true`, sets the state of the interface as UP when the peer has limited LACP capability. Use case: When a device connected to this AE port is ZTPing for the first time, it will not have LACP configured on the other end. **Note:** Turning this on will enable force-up on one of the interfaces in the bundle only<br><br>**Default**: `false` |
| `AeLacpPassive` | `*bool` | Optional | If `aggregated`==`true`, sets LACP to passive mode on this AE interface; by default, active (fast) mode is used<br><br>**Default**: `false` |
| `AeLacpSlow` | `*bool` | Optional | To use slow timeout |
| `Aggregated` | `*bool` | Optional | Whether this port is configured as an aggregated Ethernet member<br><br>**Default**: `false` |
| `Critical` | `*bool` | Optional | To generate port up/down alarm<br><br>**Default**: `false` |
| `Description` | `*string` | Optional | Human-readable description for this Junos port |
| `DisableAutoneg` | `*bool` | Optional | If `speed` and `duplex` are specified, whether to disable autonegotiation<br><br>**Default**: `false` |
| `Duplex` | [`*models.JunosPortConfigDuplexEnum`](../../doc/models/junos-port-config-duplex-enum.md) | Optional | enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `DynamicUsage` | `models.Optional[string]` | Optional | Enable dynamic usage for this port. Set to `dynamic` to enable. |
| `Esilag` | `*bool` | Optional | Whether this Junos port participates in an ESI-LAG |
| `Mtu` | `*int` | Optional | Media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation<br><br>**Default**: `1514` |
| `Networks` | `[]string` | Optional | List of network names. Required if `usage`==`inet` |
| `NoLocalOverwrite` | `*bool` | Optional | Prevent helpdesk to override the port config<br><br>**Default**: `true` |
| `PoeDisabled` | `*bool` | Optional | Whether PoE capabilities are disabled for this Junos port<br><br>**Default**: `false` |
| `PortNetwork` | `*string` | Optional | Required if `usage`==`vlan_tunnel`. Q-in-Q tunneling using All-in-one bundling. This also enables standard L2PT for interfaces that are not encapsulation tunnel interfaces and uses MAC rewrite operation. [View more information](https://www.juniper.net/documentation/us/en/software/junos/multicast-l2/topics/topic-map/q-in-q.html#id-understanding-qinq-tunneling-and-vlan-translation) |
| `Speed` | [`*models.JunosPortConfigSpeedEnum`](../../doc/models/junos-port-config-speed-enum.md) | Optional | enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br><br>**Default**: `"auto"` |
| `Usage` | `string` | Required | Port usage name. For Q-in-Q, use `vlan_tunnel`. If EVPN is used, use `evpn_uplink`or `evpn_downlink` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    junosPortConfig := models.JunosPortConfig{
        AeDisableLacp:        models.ToPointer(false),
        AeIdx:                models.ToPointer(52),
        AeLacpForceUp:        models.ToPointer(false),
        AeLacpPassive:        models.ToPointer(false),
        AeLacpSlow:           models.ToPointer(false),
        Aggregated:           models.ToPointer(false),
        Critical:             models.ToPointer(false),
        DisableAutoneg:       models.ToPointer(false),
        Duplex:               models.ToPointer(models.JunosPortConfigDuplexEnum_AUTO),
        Mtu:                  models.ToPointer(1514),
        NoLocalOverwrite:     models.ToPointer(true),
        PoeDisabled:          models.ToPointer(false),
        Speed:                models.ToPointer(models.JunosPortConfigSpeedEnum_AUTO),
        Usage:                "usage2",
    }

}
```

