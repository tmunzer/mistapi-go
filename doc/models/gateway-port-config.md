
# Gateway Port Config

Gateway port config

## Structure

`GatewayPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AeDisableLacp` | `*bool` | Optional | if `aggregated`==`true`. To disable LCP support for the AE interface<br>**Default**: `false` |
| `AeIdx` | `models.Optional[string]` | Optional | if `aggregated`==`true`. Users could force to use the designated AE name (must be an integer between 0 and 127) |
| `AeLacpForceUp` | `*bool` | Optional | For SRX Only, if `aggregated`==`true`.Sets the state of the interface as UP when the peer has limited LACP capability.\n<br>Use case: When a device connected to this AE port is ZTPing for the first time, it will not have LACP configured on the other end\n<br>Note: Turning this on will enable force-up on one of the interfaces in the bundle only<br>**Default**: `false` |
| `Aggregated` | `*bool` | Optional | **Default**: `false` |
| `Critical` | `*bool` | Optional | if want to generate port up/down alarm, set it to true<br>**Default**: `false` |
| `Description` | `*string` | Optional | - |
| `DisableAutoneg` | `*bool` | Optional | **Default**: `false` |
| `Disabled` | `*bool` | Optional | port admin up (true) / down (false)<br>**Default**: `false` |
| `DslType` | [`*models.GatewayPortDslTypeEnum`](../../doc/models/gateway-port-dsl-type-enum.md) | Optional | if `wan_type`==`dsl`. enum: `adsl`, `vdsl`<br>**Default**: `"vdsl"` |
| `DslVci` | `*int` | Optional | if `wan_type`==`dsl`<br>16 bit int<br>**Default**: `35` |
| `DslVpi` | `*int` | Optional | if `wan_type`==`dsl`<br>8 bit int<br>**Default**: `0` |
| `Duplex` | [`*models.GatewayPortDuplexEnum`](../../doc/models/gateway-port-duplex-enum.md) | Optional | enum: `auto`, `full`, `half`<br>**Default**: `"auto"` |
| `IpConfig` | [`*models.GatewayPortConfigIpConfig`](../../doc/models/gateway-port-config-ip-config.md) | Optional | Junos IP Config |
| `LteApn` | `*string` | Optional | if `wan_type`==`lte` |
| `LteAuth` | [`*models.GatewayPortLteAuthEnum`](../../doc/models/gateway-port-lte-auth-enum.md) | Optional | if `wan_type`==`lte`. enum: `chap`, `none`, `pap`<br>**Default**: `"none"` |
| `LteBackup` | `*bool` | Optional | - |
| `LtePassword` | `*string` | Optional | if `wan_type`==`lte` |
| `LteUsername` | `*string` | Optional | if `wan_type`==`lte` |
| `Mtu` | `*int` | Optional | - |
| `Name` | `*string` | Optional | name that we'll use to derive config |
| `Networks` | `[]string` | Optional | if `usage`==`lan` |
| `OuterVlanId` | `*int` | Optional | for Q-in-Q |
| `PoeDisabled` | `*bool` | Optional | **Default**: `false` |
| `PortNetwork` | `*string` | Optional | if `usage`==`lan` |
| `PreserveDscp` | `*bool` | Optional | whether to preserve dscp when sending traffic over VPN (SSR-only)<br>**Default**: `true` |
| `Redundant` | `*bool` | Optional | if HA mode |
| `RethIdx` | `*int` | Optional | if HA mode |
| `RethNode` | `*string` | Optional | if HA mode |
| `RethNodes` | `[]string` | Optional | SSR only - supporting vlan-based redundancy (matching the size of `networks`) |
| `Speed` | `*string` | Optional | **Default**: `"auto"` |
| `SsrNoVirtualMac` | `*bool` | Optional | when SSR is running as VM, this is required on certain hosting platforms<br>**Default**: `false` |
| `SvrPortRange` | `*string` | Optional | for SSR only<br>**Default**: `"none"` |
| `TrafficShaping` | [`*models.GatewayTrafficShaping`](../../doc/models/gateway-traffic-shaping.md) | Optional | - |
| `Usage` | [`models.GatewayPortUsageEnum`](../../doc/models/gateway-port-usage-enum.md) | Required | port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan` |
| `VlanId` | `*int` | Optional | if WAN interface is on a VLAN |
| `VpnPaths` | [`map[string]models.GatewayPortVpnPath`](../../doc/models/gateway-port-vpn-path.md) | Optional | - |
| `WanArpPolicer` | [`*models.GatewayPortWanArpPolicerEnum`](../../doc/models/gateway-port-wan-arp-policer-enum.md) | Optional | when `wan_type`==`broadband`. enum: `default`, `max`, `recommended`<br>**Default**: `"default"` |
| `WanExtIp` | `*string` | Optional | optional, if spoke should reach this port by a different IP |
| `WanSourceNat` | [`*models.GatewayPortWanSourceNat`](../../doc/models/gateway-port-wan-source-nat.md) | Optional | optional, by default, source-NAT is performed on all WAN Ports using the interface-ip |
| `WanType` | [`*models.GatewayPortWanTypeEnum`](../../doc/models/gateway-port-wan-type-enum.md) | Optional | if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`<br>**Default**: `"broadband"` |

## Example (as JSON)

```json
{
  "ae_disable_lacp": false,
  "ae_lacp_force_up": false,
  "aggregated": false,
  "critical": false,
  "disable_autoneg": false,
  "disabled": false,
  "dsl_type": "vdsl",
  "dsl_vci": 35,
  "dsl_vpi": 0,
  "duplex": "full",
  "lte_auth": "none",
  "poe_disabled": false,
  "preserve_dscp": true,
  "reth_nodes": [
    "node0",
    "node1"
  ],
  "speed": "1g",
  "ssr_no_virtual_mac": false,
  "svr_port_range": "60000-60005",
  "usage": "lan",
  "wan_arp_policer": "default",
  "wan_type": "broadband",
  "ae_idx": "ae_idx6"
}
```

