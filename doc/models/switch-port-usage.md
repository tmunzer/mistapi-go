
# Switch Port Usage

Junos port usages

## Structure

`SwitchPortUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllNetworks` | `*bool` | Optional | Only if `mode`==`trunk` whether to trunk all network/vlans<br>**Default**: `false` |
| `AllowDhcpd` | `*bool` | Optional | Only if `mode`!=`dynamic` if DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri_state.<br><br>When it is not defined, it means using the system’s default setting which depends on whether the port is a access or trunk port. |
| `AllowMultipleSupplicants` | `*bool` | Optional | Only if `mode`!=`dynamic`<br>**Default**: `false` |
| `BypassAuthWhenServerDown` | `*bool` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x` bypass auth for known clients if set to true when RADIUS server is down<br>**Default**: `false` |
| `BypassAuthWhenServerDownForUnkonwnClient` | `*bool` | Optional | Only if `mode`!=`dynamic` and `port_auth`=`dot1x` bypass auth for all (including unknown clients) if set to true when RADIUS server is down<br>**Default**: `false` |
| `Description` | `*string` | Optional | Only if `mode`!=`dynamic` |
| `DisableAutoneg` | `*bool` | Optional | Only if `mode`!=`dynamic` if speed and duplex are specified, whether to disable autonegotiation<br>**Default**: `false` |
| `Disabled` | `*bool` | Optional | Only if `mode`!=`dynamic` whether the port is disabled<br>**Default**: `false` |
| `Duplex` | [`*models.SwitchPortUsageDuplexEnum`](../../doc/models/switch-port-usage-duplex-enum.md) | Optional | Only if `mode`!=`dynamic` link connection mode. enum: `auto`, `full`, `half`<br>**Default**: `"auto"` |
| `DynamicVlanNetworks` | `[]string` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return |
| `EnableMacAuth` | `*bool` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x` whether to enable MAC Auth<br>**Default**: `false` |
| `EnableQos` | `*bool` | Optional | Only if `mode`!=`dynamic`<br>**Default**: `false` |
| `GuestNetwork` | `models.Optional[string]` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x` which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) |
| `InterSwitchLink` | `*bool` | Optional | Only if `mode`!=`dynamic` inter_switch_link is used together with "isolation" under networks<br>NOTE: inter_switch_link works only between Juniper device. This has to be applied to both ports connected together<br>**Default**: `false` |
| `MacAuthOnly` | `*bool` | Optional | Only if `mode`!=`dynamic` and `enable_mac_auth`==`true` |
| `MacAuthPreferred` | `*bool` | Optional | Only if `mode`!=`dynamic` + `enable_mac_auth`==`true` + `mac_auth_only`==`false`, dot1x will be given priority then mac_auth. Enable this to prefer mac_auth over dot1x. |
| `MacAuthProtocol` | [`*models.SwitchPortUsageMacAuthProtocolEnum`](../../doc/models/switch-port-usage-mac-auth-protocol-enum.md) | Optional | Only if `mode`!=`dynamic` and `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`<br>**Default**: `"eap-md5"` |
| `MacLimit` | `*int` | Optional | Only if `mode`!=`dynamic` max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform<br>**Default**: `0`<br>**Constraints**: `>= 0` |
| `Mode` | [`*models.SwitchPortUsageModeEnum`](../../doc/models/switch-port-usage-mode-enum.md) | Optional | `mode`==`dynamic` must only be used with the port usage with the name `dynamic`. enum: `access`, `dynamic`, `inet`, `trunk` |
| `Mtu` | `*int` | Optional | Only if `mode`!=`dynamic` media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. |
| `Networks` | `[]string` | Optional | Only if `mode`==`trunk`, the list of network/vlans |
| `PersistMac` | `*bool` | Optional | Only if `mode`==`access` and `port_auth`!=`dot1x` whether the port should retain dynamically learned MAC addresses<br>**Default**: `false` |
| `PoeDisabled` | `*bool` | Optional | Only if `mode`!=`dynamic` whether PoE capabilities are disabled for a port<br>**Default**: `false` |
| `PortAuth` | [`models.Optional[models.SwitchPortUsageDot1xEnum]`](../../doc/models/switch-port-usage-dot-1-x-enum.md) | Optional | Only if `mode`!=`dynamic` if dot1x is desired, set to dot1x. enum: `dot1x` |
| `PortNetwork` | `*string` | Optional | Only if `mode`!=`dynamic` native network/vlan for untagged traffic |
| `ReauthInterval` | `*int` | Optional | Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range<br>**Default**: `3600`<br>**Constraints**: `>= 10`, `<= 65535` |
| `ResetDefaultWhen` | [`*models.SwitchPortUsageDynamicResetDefaultWhenEnum`](../../doc/models/switch-port-usage-dynamic-reset-default-when-enum.md) | Optional | Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage. enum: `link_down`, `none` (let the DPC port keep at the current port usage)<br>**Default**: `"link_down"` |
| `Rules` | [`[]models.SwitchPortUsageDynamicRule`](../../doc/models/switch-port-usage-dynamic-rule.md) | Optional | Only if `mode`==`dynamic` |
| `ServerFailNetwork` | `models.Optional[string]` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x` sets server fail fallback vlan |
| `ServerRejectNetwork` | `models.Optional[string]` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x` when radius server reject / fails |
| `Speed` | `*string` | Optional | Only if `mode`!=`dynamic` speed, default is auto to automatically negotiate speed |
| `StormControl` | [`*models.SwitchPortUsageStormControl`](../../doc/models/switch-port-usage-storm-control.md) | Optional | Switch storm control<br>Only if `mode`!=`dynamic` |
| `StpEdge` | `*bool` | Optional | Only if `mode`!=`dynamic` when enabled, the port is not expected to receive BPDU frames<br>**Default**: `false` |
| `StpNoRootPort` | `*bool` | Optional | **Default**: `false` |
| `StpP2p` | `*bool` | Optional | **Default**: `false` |
| `VoipNetwork` | `*string` | Optional | Only if `mode`!=`dynamic` network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth |

## Example (as JSON)

```json
{
  "all_networks": false,
  "allow_multiple_supplicants": false,
  "bypass_auth_when_server_down": false,
  "bypass_auth_when_server_down_for_unkonwn_client": false,
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
  "persist_mac": false,
  "poe_disabled": false,
  "reauth_interval": 3600,
  "reset_default_when": "link_down",
  "stp_edge": false,
  "stp_no_root_port": false,
  "stp_p2p": false,
  "allow_dhcpd": false
}
```

