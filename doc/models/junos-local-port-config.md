
# Junos Local Port Config

Switch port config

*This model accepts additional fields of type interface{}.*

## Structure

`JunosLocalPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllNetworks` | `*bool` | Optional | Only if `mode`==`trunk` whether to trunk all network/vlans<br><br>**Default**: `false` |
| `AllowDhcpd` | `*bool` | Optional | If DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri_state. When it is not defined, it means using the system's default setting which depends on whether the port is an access or trunk port. |
| `AllowMultipleSupplicants` | `*bool` | Optional | **Default**: `false` |
| `BypassAuthWhenServerDown` | `*bool` | Optional | Only if `port_auth`==`dot1x` bypass auth for known clients if set to true when RADIUS server is down<br><br>**Default**: `false` |
| `BypassAuthWhenServerDownForUnknownClient` | `*bool` | Optional | Only if `port_auth`=`dot1x` bypass auth for all (including unknown clients) if set to true when RADIUS server is down<br><br>**Default**: `false` |
| `Description` | `*string` | Optional | - |
| `DisableAutoneg` | `*bool` | Optional | Only if `mode`!=`dynamic` if speed and duplex are specified, whether to disable autonegotiation<br><br>**Default**: `false` |
| `Disabled` | `*bool` | Optional | Whether the port is disabled<br><br>**Default**: `false` |
| `Duplex` | [`*models.SwitchPortLocalUsageDuplexEnum`](../../doc/models/switch-port-local-usage-duplex-enum.md) | Optional | link connection mode. enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `DynamicVlanNetworks` | `[]string` | Optional | Only if `port_auth`==`dot1x`, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return |
| `EnableMacAuth` | `*bool` | Optional | Only if `port_auth`==`dot1x` whether to enable MAC Auth<br><br>**Default**: `false` |
| `EnableQos` | `*bool` | Optional | **Default**: `false` |
| `GuestNetwork` | `models.Optional[string]` | Optional | Only if `port_auth`==`dot1x` which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) |
| `InterSwitchLink` | `*bool` | Optional | inter_switch_link is used together with "isolation" under networks. NOTE: inter_switch_link works only between Juniper device. This has to be applied to both ports connected together<br><br>**Default**: `false` |
| `MacAuthOnly` | `*bool` | Optional | Only if `enable_mac_auth`==`true` |
| `MacAuthPreferred` | `*bool` | Optional | Only if `enable_mac_auth`==`true` + `mac_auth_only`==`false`, dot1x will be given priority then mac_auth. Enable this to prefer mac_auth over dot1x. |
| `MacAuthProtocol` | [`*models.SwitchPortLocalUsageMacAuthProtocolEnum`](../../doc/models/switch-port-local-usage-mac-auth-protocol-enum.md) | Optional | Only if `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`<br><br>**Default**: `"eap-md5"` |
| `MacLimit` | `*int` | Optional | Max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0` |
| `Mode` | [`*models.SwitchPortLocalUsageModeEnum`](../../doc/models/switch-port-local-usage-mode-enum.md) | Optional | enum: `access`, `inet`, `trunk` |
| `Mtu` | `*int` | Optional | Media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. |
| `Networks` | `[]string` | Optional | Only if `mode`==`trunk`, the list of network/vlans |
| `Note` | `*string` | Optional | Additional note for the port config override |
| `PersistMac` | `*bool` | Optional | Only if `mode`==`access` and `port_auth`!=`dot1x` whether the port should retain dynamically learned MAC addresses<br><br>**Default**: `false` |
| `PoeDisabled` | `*bool` | Optional | Whether PoE capabilities are disabled for a port<br><br>**Default**: `false` |
| `PortAuth` | [`models.Optional[models.SwitchPortLocalUsageDot1xEnum]`](../../doc/models/switch-port-local-usage-dot-1-x-enum.md) | Optional | if dot1x is desired, set to dot1x. enum: `dot1x` |
| `PortNetwork` | `*string` | Optional | Native network/vlan for untagged traffic |
| `ReauthInterval` | [`*models.SwitchPortUsageReauthInterval`](../../doc/models/containers/switch-port-usage-reauth-interval.md) | Optional | Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600) |
| `ServerFailNetwork` | `models.Optional[string]` | Optional | Only if `port_auth`==`dot1x` sets server fail fallback vlan |
| `ServerRejectNetwork` | `models.Optional[string]` | Optional | Only if `port_auth`==`dot1x` when radius server reject / fails |
| `Speed` | [`*models.JunosPortConfigSpeedEnum`](../../doc/models/junos-port-config-speed-enum.md) | Optional | enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br><br>**Default**: `"auto"` |
| `StormControl` | [`*models.SwitchPortLocalUsageStormControl`](../../doc/models/switch-port-local-usage-storm-control.md) | Optional | Switch storm control |
| `StpEdge` | `*bool` | Optional | When enabled, the port is not expected to receive BPDU frames<br><br>**Default**: `false` |
| `StpNoRootPort` | `*bool` | Optional | **Default**: `false` |
| `StpP2p` | `*bool` | Optional | **Default**: `false` |
| `Usage` | `string` | Required | Port usage name. |
| `UseVstp` | `*bool` | Optional | If this is connected to a vstp network<br><br>**Default**: `false` |
| `VoipNetwork` | `*string` | Optional | Network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "all_networks": false,
  "allow_multiple_supplicants": false,
  "bypass_auth_when_server_down": false,
  "bypass_auth_when_server_down_for_unknown_client": false,
  "disable_autoneg": false,
  "disabled": false,
  "duplex": "auto",
  "dynamic_vlan_networks": [
    "corp",
    "user"
  ],
  "enable_mac_auth": false,
  "enable_qos": false,
  "inter_switch_link": false,
  "mac_auth_protocol": "eap-md5",
  "mac_limit": 0,
  "note": "force 100M for camera",
  "persist_mac": false,
  "poe_disabled": false,
  "speed": "auto",
  "stp_edge": false,
  "stp_no_root_port": false,
  "stp_p2p": false,
  "usage": "usage8",
  "use_vstp": false,
  "allow_dhcpd": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

