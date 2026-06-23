
# Junos Local Port Config

Per-port Switch Port Operator (SPO) override configuration used in `local_port_config` to customize settings inherited from `port_config`

## Structure

`JunosLocalPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllNetworks` | `*bool` | Optional | Only if `mode`==`trunk` whether to trunk all network/vlans<br><br>**Default**: `false` |
| `AllowDhcpd` | `*bool` | Optional | Controls whether DHCP server traffic is allowed on ports using this configuration if DHCP snooping is enabled. This is a tri-state setting; `true`: ports become trusted ports allowing DHCP server traffic, `false`: ports become untrusted blocking DHCP server traffic, undefined: use system defaults (access ports default to untrusted, trunk ports default to trusted). |
| `AllowMultipleSupplicants` | `*bool` | Optional | Whether multiple supplicants may authenticate on the port<br><br>**Default**: `false` |
| `BypassAuthWhenServerDown` | `*bool` | Optional | Only if `port_auth`==`dot1x` bypass auth for known clients if set to true when RADIUS server is down<br><br>**Default**: `false` |
| `BypassAuthWhenServerDownForUnknownClient` | `*bool` | Optional | Only if `port_auth`=`dot1x` bypass auth for all (including unknown clients) if set to true when RADIUS server is down<br><br>**Default**: `false` |
| `Description` | `*string` | Optional | Human-readable description for this local port configuration |
| `DisableAutoneg` | `*bool` | Optional | Only if `mode`!=`dynamic` if speed and duplex are specified, whether to disable autonegotiation<br><br>**Default**: `false` |
| `Disabled` | `*bool` | Optional | Whether the port is disabled<br><br>**Default**: `false` |
| `Duplex` | [`*models.SwitchPortLocalUsageDuplexEnum`](../../doc/models/switch-port-local-usage-duplex-enum.md) | Optional | link connection mode. enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `DynamicVlanNetworks` | `[]string` | Optional | Only if `port_auth`==`dot1x`, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return |
| `EnableMacAuth` | `*bool` | Optional | Only if `port_auth`==`dot1x` whether to enable MAC Auth<br><br>**Default**: `false` |
| `EnableQos` | `*bool` | Optional | Whether QoS is enabled on ports using this local configuration<br><br>**Default**: `false` |
| `GuestNetwork` | `models.Optional[string]` | Optional | Only if `port_auth`==`dot1x` which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) |
| `InterSwitchLink` | `*bool` | Optional | Used together with "isolation" under networks for links between Juniper devices; must be applied to both connected ports<br><br>**Default**: `false` |
| `MacAuthOnly` | `*bool` | Optional | Only if `enable_mac_auth`==`true`, whether to use MAC authentication without 802.1X |
| `MacAuthPreferred` | `*bool` | Optional | Only if `enable_mac_auth`==`true` + `mac_auth_only`==`false`, dot1x will be given priority then mac_auth. Enable this to prefer mac_auth over dot1x. |
| `MacAuthProtocol` | [`*models.SwitchPortLocalUsageMacAuthProtocolEnum`](../../doc/models/switch-port-local-usage-mac-auth-protocol-enum.md) | Optional | Only if `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`<br><br>**Default**: `"eap-md5"` |
| `MacLimit` | `*int` | Optional | Max number of MAC addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0` |
| `Mode` | [`*models.SwitchPortLocalUsageModeEnum`](../../doc/models/switch-port-local-usage-mode-enum.md) | Optional | enum: `access`, `inet`, `trunk` |
| `Mtu` | `*int` | Optional | Media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. |
| `Networks` | `[]string` | Optional | Only if `mode`==`trunk`, the list of network/vlans |
| `Note` | `*string` | Optional | Additional note for the port config override |
| `PersistMac` | `*bool` | Optional | Only if `mode`==`access` and `port_auth`!=`dot1x` whether the port should retain dynamically learned MAC addresses<br><br>**Default**: `false` |
| `PoeDisabled` | `*bool` | Optional | Whether PoE capabilities are disabled for a port<br><br>**Default**: `false` |
| `PortAuth` | [`models.Optional[models.SwitchPortLocalUsageDot1xEnum]`](../../doc/models/switch-port-local-usage-dot-1-x-enum.md) | Optional | if dot1x is desired, set to dot1x. enum: `dot1x` |
| `PortNetwork` | `*string` | Optional | Native network/vlan for untagged traffic |
| `ReauthInterval` | [`*models.SwitchPortUsageReauthInterval`](../../doc/models/containers/switch-port-usage-reauth-interval.md) | Optional | Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600). Set to 0 to disable reauthentication (no-reauthentication). |
| `ServerFailNetwork` | `models.Optional[string]` | Optional | Only if `port_auth`==`dot1x` sets server fail fallback vlan |
| `ServerRejectNetwork` | `models.Optional[string]` | Optional | Only if `port_auth`==`dot1x` when RADIUS server reject / fails |
| `Speed` | [`*models.JunosPortConfigSpeedEnum`](../../doc/models/junos-port-config-speed-enum.md) | Optional | enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br><br>**Default**: `"auto"` |
| `StormControl` | [`*models.SwitchPortLocalUsageStormControl`](../../doc/models/switch-port-local-usage-storm-control.md) | Optional | Storm-control settings for this local port configuration |
| `StpEdge` | `*bool` | Optional | When enabled, the port is not expected to receive BPDU frames<br><br>**Default**: `false` |
| `StpNoRootPort` | `*bool` | Optional | Whether STP should prevent this port from becoming a root port<br><br>**Default**: `false` |
| `StpP2p` | `*bool` | Optional | Whether STP treats this port as a point-to-point link<br><br>**Default**: `false` |
| `Usage` | `string` | Required | Port usage profile name for this local port configuration |
| `UseVstp` | `*bool` | Optional | If this is connected to a vstp network<br><br>**Default**: `false` |
| `VoipNetwork` | `*string` | Optional | Network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    junosLocalPortConfig := models.JunosLocalPortConfig{
        AllNetworks:                              models.ToPointer(false),
        AllowDhcpd:                               models.ToPointer(false),
        AllowMultipleSupplicants:                 models.ToPointer(false),
        BypassAuthWhenServerDown:                 models.ToPointer(false),
        BypassAuthWhenServerDownForUnknownClient: models.ToPointer(false),
        DisableAutoneg:                           models.ToPointer(false),
        Disabled:                                 models.ToPointer(false),
        Duplex:                                   models.ToPointer(models.SwitchPortLocalUsageDuplexEnum_AUTO),
        DynamicVlanNetworks:                      []string{
            "corp",
            "user",
        },
        EnableMacAuth:                            models.ToPointer(false),
        EnableQos:                                models.ToPointer(false),
        InterSwitchLink:                          models.ToPointer(false),
        MacAuthProtocol:                          models.ToPointer(models.SwitchPortLocalUsageMacAuthProtocolEnum_EAPMD5),
        MacLimit:                                 models.ToPointer(0),
        Note:                                     models.ToPointer("force 100M for camera"),
        PersistMac:                               models.ToPointer(false),
        PoeDisabled:                              models.ToPointer(false),
        Speed:                                    models.ToPointer(models.JunosPortConfigSpeedEnum_AUTO),
        StpEdge:                                  models.ToPointer(false),
        StpNoRootPort:                            models.ToPointer(false),
        StpP2p:                                   models.ToPointer(false),
        Usage:                                    "usage2",
        UseVstp:                                  models.ToPointer(false),
    }

}
```

