
# Switch Port Usage

Junos port usages

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchPortUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllNetworks` | `*bool` | Optional | Only if `mode`==`trunk`. Whether to trunk all network/vlans<br><br>**Default**: `false` |
| `AllowDhcpd` | `*bool` | Optional | Only applies when `mode`!=`dynamic`. Controls whether DHCP server traffic is allowed on ports using this configuration if DHCP snooping is enabled. This is a tri-state setting; true: ports become trusted ports allowing DHCP server traffic, false: ports become untrusted blocking DHCP server traffic, undefined: use system defaults (access ports default to untrusted, trunk ports default to trusted). |
| `AllowMultipleSupplicants` | `*bool` | Optional | Only if `mode`!=`dynamic`<br><br>**Default**: `false` |
| `BypassAuthWhenServerDown` | `*bool` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Bypass auth for known clients if set to true when RADIUS server is down<br><br>**Default**: `false` |
| `BypassAuthWhenServerDownForUnknownClient` | `*bool` | Optional | Only if `mode`!=`dynamic` and `port_auth`=`dot1x`. Bypass auth for all (including unknown clients) if set to true when RADIUS server is down<br><br>**Default**: `false` |
| `BypassAuthWhenServerDownForVoip` | `*bool` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Bypass auth for VOIP if set to true when RADIUS server is down<br><br>**Default**: `false` |
| `CommunityVlanId` | `*int` | Optional | Only if `mode`!=`dynamic`. To be used together with `isolation` under networks. Signaling that this port connects to the networks isolated but wired clients belong to the same community can talk to each other |
| `Description` | `*string` | Optional | Only if `mode`!=`dynamic` |
| `DisableAutoneg` | `*bool` | Optional | Only if `mode`!=`dynamic`. If speed and duplex are specified, whether to disable autonegotiation<br><br>**Default**: `false` |
| `Disabled` | `*bool` | Optional | Only if `mode`!=`dynamic`. Whether the port is disabled<br><br>**Default**: `false` |
| `Duplex` | [`*models.SwitchPortUsageDuplexEnum`](../../doc/models/switch-port-usage-duplex-enum.md) | Optional | Only if `mode`!=`dynamic`. Link connection mode. enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `DynamicVlanNetworks` | `[]string` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return |
| `EnableMacAuth` | `*bool` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Whether to enable MAC Auth<br><br>**Default**: `false` |
| `EnableQos` | `*bool` | Optional | Only if `mode`!=`dynamic`<br><br>**Default**: `false` |
| `GuestNetwork` | `models.Optional[string]` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed) |
| `InterIsolationNetworkLink` | `*bool` | Optional | Only if `mode`!=`dynamic`. `inter_isolation_network_link` is used together with `isolation` under networks, signaling that this port connects to isolated networks<br><br>**Default**: `false` |
| `InterSwitchLink` | `*bool` | Optional | Only if `mode`!=`dynamic`. `inter_switch_link` is used together with `isolation` under networks. NOTE: `inter_switch_link` works only between Juniper devices. This has to be applied to both ports connected together<br><br>**Default**: `false` |
| `MacAuthOnly` | `*bool` | Optional | Only if `mode`!=`dynamic` and `enable_mac_auth`==`true` |
| `MacAuthPreferred` | `*bool` | Optional | Only if `mode`!=`dynamic` + `enable_mac_auth`==`true` + `mac_auth_only`==`false`, dot1x will be given priority then mac_auth. Enable this to prefer mac_auth over dot1x. |
| `MacAuthProtocol` | [`*models.SwitchPortUsageMacAuthProtocolEnum`](../../doc/models/switch-port-usage-mac-auth-protocol-enum.md) | Optional | Only if `mode`!=`dynamic` and `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`<br><br>**Default**: `"eap-md5"` |
| `MacLimit` | [`*models.SwitchPortUsageMacLimit`](../../doc/models/containers/switch-port-usage-mac-limit.md) | Optional | Only if `mode`!=`dynamic`, max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform) |
| `Mode` | [`*models.SwitchPortUsageModeEnum`](../../doc/models/switch-port-usage-mode-enum.md) | Optional | `mode`==`dynamic` must only be used if the port usage name is `dynamic`. enum: `access`, `dynamic`, `inet`, `trunk` |
| `Mtu` | [`*models.SwitchPortUsageMtu`](../../doc/models/containers/switch-port-usage-mtu.md) | Optional | Only if `mode`!=`dynamic` media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514. |
| `Networks` | `[]string` | Optional | Only if `mode`==`trunk`, the list of network/vlans |
| `PersistMac` | `*bool` | Optional | Only if `mode`==`access` and `port_auth`!=`dot1x`. Whether the port should retain dynamically learned MAC addresses<br><br>**Default**: `false` |
| `PoeDisabled` | `*bool` | Optional | Only if `mode`!=`dynamic`. Whether PoE capabilities are disabled for a port<br><br>**Default**: `false` |
| `PoePriority` | [`*models.PoePriorityEnum`](../../doc/models/poe-priority-enum.md) | Optional | PoE priority. enum: `low`, `high` |
| `PortAuth` | [`models.Optional[models.SwitchPortUsageDot1xEnum]`](../../doc/models/switch-port-usage-dot-1-x-enum.md) | Optional | Only if `mode`!=`dynamic`. If dot1x is desired, set to dot1x. enum: `dot1x` |
| `PortNetwork` | `*string` | Optional | Only if `mode`!=`dynamic`. Native network/vlan for untagged traffic |
| `ReauthInterval` | [`*models.SwitchPortUsageReauthInterval`](../../doc/models/containers/switch-port-usage-reauth-interval.md) | Optional | Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600) |
| `ResetDefaultWhen` | [`*models.SwitchPortUsageDynamicResetDefaultWhenEnum`](../../doc/models/switch-port-usage-dynamic-reset-default-when-enum.md) | Optional | Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage. enum: `link_down`, `none` (let the DPC port keep at the current port usage)<br><br>**Default**: `"link_down"` |
| `Rules` | [`[]models.SwitchPortUsageDynamicRule`](../../doc/models/switch-port-usage-dynamic-rule.md) | Optional | Only if `mode`==`dynamic` |
| `ServerFailNetwork` | `models.Optional[string]` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Sets server fail fallback vlan |
| `ServerRejectNetwork` | `models.Optional[string]` | Optional | Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. When radius server reject / fails |
| `Speed` | [`*models.SwitchPortUsageSpeedEnum`](../../doc/models/switch-port-usage-speed-enum.md) | Optional | Only if `mode`!=`dynamic`, Port speed, default is auto to automatically negotiate speed enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br><br>**Default**: `"auto"` |
| `StormControl` | [`*models.SwitchPortUsageStormControl`](../../doc/models/switch-port-usage-storm-control.md) | Optional | Switch storm control. Only if `mode`!=`dynamic` |
| `StpDisable` | `*bool` | Optional | Only if `mode`!=`dynamic` and `stp_required`==`false`. Drop bridge protocol data units (BPDUs ) that enter any interface or a specified interface<br><br>**Default**: `false` |
| `StpEdge` | `*bool` | Optional | Only if `mode`!=`dynamic`. When enabled, the port is not expected to receive BPDU frames<br><br>**Default**: `false` |
| `StpNoRootPort` | `*bool` | Optional | Only if `mode`!=`dynamic`<br><br>**Default**: `false` |
| `StpP2p` | `*bool` | Optional | Only if `mode`!=`dynamic`<br><br>**Default**: `false` |
| `StpRequired` | `*bool` | Optional | Only if `mode`!=`dynamic`. Whether to remain in block state if no BPDU is received<br><br>**Default**: `false` |
| `UiEvpntopoId` | `*uuid.UUID` | Optional | Optional for Campus Fabric Core-Distribution ESI-LAG profile. Helper used by the UI to select this port profile as the ESI-Lag between Distribution and Access switches |
| `UseVstp` | `*bool` | Optional | If this is connected to a vstp network<br><br>**Default**: `false` |
| `VoipNetwork` | `models.Optional[string]` | Optional | Only if `mode`!=`dynamic`. Network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "all_networks": false,
  "allow_multiple_supplicants": false,
  "bypass_auth_when_server_down": false,
  "bypass_auth_when_server_down_for_unknown_client": false,
  "bypass_auth_when_server_down_for_voip": false,
  "disable_autoneg": false,
  "disabled": false,
  "duplex": "auto",
  "dynamic_vlan_networks": [
    "corp",
    "user"
  ],
  "enable_mac_auth": false,
  "enable_qos": false,
  "inter_isolation_network_link": false,
  "inter_switch_link": false,
  "mac_auth_protocol": "eap-md5",
  "persist_mac": false,
  "poe_disabled": false,
  "reset_default_when": "link_down",
  "speed": "auto",
  "stp_disable": false,
  "stp_edge": false,
  "stp_no_root_port": false,
  "stp_p2p": false,
  "stp_required": false,
  "use_vstp": false,
  "allow_dhcpd": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

