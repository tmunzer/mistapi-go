
# Port Config 1

eth0 is not allowed here.
Property key is the interface(s) name (e.g. "eth1" or"eth1,eth2")

## Structure

`PortConfig1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AeDisableLacp` | `*bool` | Optional | To disable LACP support for the AE interface |
| `AeIdx` | `*int` | Optional | Users could force to use the designated AE name |
| `AeLacpSlow` | `*bool` | Optional | to use fast timeout<br>**Default**: `true` |
| `Aggregated` | `*bool` | Optional | **Default**: `false` |
| `Critical` | `*bool` | Optional | if want to generate port up/down alarm |
| `Description` | `*string` | Optional | - |
| `DisableAutoneg` | `*bool` | Optional | if `speed` and `duplex` are specified, whether to disable autonegotiation<br>**Default**: `false` |
| `Duplex` | [`*models.JunosPortConfigDuplexEnum`](../../doc/models/junos-port-config-duplex-enum.md) | Optional | **Default**: `"auto"` |
| `DynamicUsage` | `models.Optional[string]` | Optional | Enable dynamic usage for this port. Set to `dynamic` to enable. |
| `Esilag` | `*bool` | Optional | - |
| `Mtu` | `*int` | Optional | media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation<br>**Default**: `1514` |
| `NoLocalOverwrite` | `*bool` | Optional | prevent helpdesk to override the port config |
| `PoeDisabled` | `*bool` | Optional | **Default**: `false` |
| `Speed` | [`*models.JunosPortConfigSpeedEnum`](../../doc/models/junos-port-config-speed-enum.md) | Optional | **Default**: `"auto"` |
| `Usage` | [`models.GatewayPortUsage1Enum`](../../doc/models/gateway-port-usage-1-enum.md) | Required | port usage name.<br><br>If EVPN is used, use `evpn_uplink`or `evpn_downlink` |
| `Disabled` | `*bool` | Optional | port admin up (true) / down (false)<br>**Default**: `false` |
| `DslType` | [`*models.GatewayPortDslTypeEnum`](../../doc/models/gateway-port-dsl-type-enum.md) | Optional | if `wan_type`==`lte`<br>**Default**: `"vdsl"` |
| `DslVci` | `*int` | Optional | if `wan_type`==`dsl`<br>16 bit int<br>**Default**: `35` |
| `DslVpi` | `*int` | Optional | if `wan_type`==`dsl`<br>8 bit int<br>**Default**: `0` |
| `IpConfig` | [`*models.GatewayPortConfigIpConfig`](../../doc/models/gateway-port-config-ip-config.md) | Optional | Junos IP Config |
| `LteApn` | `*string` | Optional | if `wan_type`==`lte` |
| `LteAuth` | [`*models.GatewayPortLteAuthEnum`](../../doc/models/gateway-port-lte-auth-enum.md) | Optional | if `wan_type`==`lte`<br>**Default**: `"none"` |
| `LteBackup` | `*bool` | Optional | - |
| `LtePassword` | `*string` | Optional | if `wan_type`==`lte` |
| `LteUsername` | `*string` | Optional | if `wan_type`==`lte` |
| `Name` | `*string` | Optional | name that we'll use to derive config |
| `Networks` | `[]string` | Optional | if `usage`==`lan` |
| `OuterVlanId` | `*int` | Optional | for Q-in-Q |
| `PortNetwork` | `*string` | Optional | if `usage`==`lan` |
| `PreserveDscp` | `*bool` | Optional | whether to preserve dscp when sending traffic over VPN (SSR-only)<br>**Default**: `true` |
| `Redundant` | `*bool` | Optional | if HA mode |
| `RethIdx` | `*int` | Optional | if HA mode |
| `RethNode` | `*string` | Optional | if HA mode |
| `RethNodes` | `[]string` | Optional | SSR only - supporting vlan-based redundancy (matching the size of `networks`) |
| `SsrNoVirtualMac` | `*bool` | Optional | when SSR is running as VM, this is required on certain hosting platforms<br>**Default**: `false` |
| `SvrPortRange` | `*string` | Optional | for SSR only<br>**Default**: `"none"` |
| `TrafficShaping` | [`*models.GatewayTrafficShaping`](../../doc/models/gateway-traffic-shaping.md) | Optional | - |
| `VlanId` | `*int` | Optional | if WAN interface is on a VLAN |
| `VpnPaths` | [`map[string]models.GatewayPortVpnPath`](../../doc/models/gateway-port-vpn-path.md) | Optional | - |
| `WanArpPolicer` | [`*models.GatewayPortWanArpPolicerEnum`](../../doc/models/gateway-port-wan-arp-policer-enum.md) | Optional | when `wan_type`==`broadband`<br>**Default**: `"recommended"` |
| `WanExtIp` | `*string` | Optional | optional, if spoke should reach this port by a different IP |
| `WanSourceNat` | [`*models.GatewayPortWanSourceNat`](../../doc/models/gateway-port-wan-source-nat.md) | Optional | optional, by default, source-NAT is performed on all WAN Ports using the interface-ip |
| `WanType` | [`*models.GatewayPortWanTypeEnum`](../../doc/models/gateway-port-wan-type-enum.md) | Optional | if `usage`==`wan`<br>**Default**: `"broadband"` |

## Example (as JSON)

```json
{
  "ae_lacp_slow": true,
  "aggregated": false,
  "disable_autoneg": false,
  "duplex": "full",
  "mtu": 1514,
  "poe_disabled": false,
  "speed": "1g",
  "usage": "lan",
  "disabled": false,
  "dsl_type": "vdsl",
  "dsl_vci": 35,
  "dsl_vpi": 0,
  "lte_auth": "none",
  "preserve_dscp": true,
  "reth_nodes": [
    "node0",
    "node1"
  ],
  "ssr_no_virtual_mac": false,
  "svr_port_range": "60000-60005",
  "wan_arp_policer": "recommended",
  "wan_type": "broadband",
  "ae_disable_lacp": false,
  "ae_idx": 156,
  "critical": false
}
```

