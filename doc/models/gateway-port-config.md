
# Gateway Port Config

Gateway port config

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AeDisableLacp` | `*bool` | Optional | If `aggregated`==`true`. To disable LCP support for the AE interface<br>**Default**: `false` |
| `AeIdx` | `models.Optional[string]` | Optional | If `aggregated`==`true`. Users could force to use the designated AE name (must be an integer between 0 and 127) |
| `AeLacpForceUp` | `*bool` | Optional | For SRX Only, if `aggregated`==`true`.Sets the state of the interface as UP when the peer has limited LACP capability. Use case: When a device connected to this AE port is ZTPing for the first time, it will not have LACP configured on the other end. **Note:** Turning this on will enable force-up on one of the interfaces in the bundle only<br>**Default**: `false` |
| `Aggregated` | `*bool` | Optional | **Default**: `false` |
| `Critical` | `*bool` | Optional | To generate port up/down alarm, set it to true<br>**Default**: `false` |
| `Description` | `*string` | Optional | Interface Description. Can be a variable (i.e. "{{myvar}}") |
| `DisableAutoneg` | `*bool` | Optional | **Default**: `false` |
| `Disabled` | `*bool` | Optional | Port admin up (true) / down (false)<br>**Default**: `false` |
| `DslType` | [`*models.GatewayPortDslTypeEnum`](../../doc/models/gateway-port-dsl-type-enum.md) | Optional | if `wan_type`==`dsl`. enum: `adsl`, `vdsl`<br>**Default**: `"vdsl"` |
| `DslVci` | `*int` | Optional | If `wan_type`==`dsl`, 16 bit int<br>**Default**: `35` |
| `DslVpi` | `*int` | Optional | If `wan_type`==`dsl`, 8 bit int<br>**Default**: `0` |
| `Duplex` | [`*models.GatewayPortDuplexEnum`](../../doc/models/gateway-port-duplex-enum.md) | Optional | enum: `auto`, `full`, `half`<br>**Default**: `"auto"` |
| `IpConfig` | [`*models.GatewayPortConfigIpConfig`](../../doc/models/gateway-port-config-ip-config.md) | Optional | Junos IP Config |
| `LteApn` | `*string` | Optional | If `wan_type`==`lte` |
| `LteAuth` | [`*models.GatewayPortLteAuthEnum`](../../doc/models/gateway-port-lte-auth-enum.md) | Optional | if `wan_type`==`lte`. enum: `chap`, `none`, `pap`<br>**Default**: `"none"` |
| `LteBackup` | `*bool` | Optional | - |
| `LtePassword` | `*string` | Optional | If `wan_type`==`lte` |
| `LteUsername` | `*string` | Optional | If `wan_type`==`lte` |
| `Mtu` | `*int` | Optional | - |
| `Name` | `*string` | Optional | Name that we'll use to derive config |
| `Networks` | `[]string` | Optional | If `usage`==`lan`, name of the [networks]($h/Orgs%20Networks/_overview) to attach to the interface |
| `OuterVlanId` | `*int` | Optional | For Q-in-Q |
| `PoeDisabled` | `*bool` | Optional | **Default**: `false` |
| `PortNetwork` | `*string` | Optional | Only for SRX and if `usage`==`lan`, the name of the Network to be used as the Untagged VLAN |
| `PreserveDscp` | `*bool` | Optional | Whether to preserve dscp when sending traffic over VPN (SSR-only)<br>**Default**: `true` |
| `Redundant` | `*bool` | Optional | If HA mode |
| `RedundantGroup` | `*int` | Optional | If HA mode, SRX Only - support redundancy-group. 1-128 for physical SRX, 1-64 for virtual SRX<br>**Constraints**: `>= 1`, `<= 128` |
| `RethIdx` | `*int` | Optional | If HA mode |
| `RethNode` | `*string` | Optional | If HA mode |
| `RethNodes` | `[]string` | Optional | SSR only - supporting vlan-based redundancy (matching the size of `networks`) |
| `Speed` | `*string` | Optional | **Default**: `"auto"` |
| `SsrNoVirtualMac` | `*bool` | Optional | When SSR is running as VM, this is required on certain hosting platforms<br>**Default**: `false` |
| `SvrPortRange` | `*string` | Optional | For SSR only<br>**Default**: `"none"` |
| `TrafficShaping` | [`*models.GatewayTrafficShaping`](../../doc/models/gateway-traffic-shaping.md) | Optional | - |
| `Usage` | [`models.GatewayPortUsageEnum`](../../doc/models/gateway-port-usage-enum.md) | Required | port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan` |
| `VlanId` | [`*models.GatewayPortVlanIdWithVariable`](../../doc/models/containers/gateway-port-vlan-id-with-variable.md) | Optional | If WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}") |
| `VpnPaths` | [`map[string]models.GatewayPortVpnPath`](../../doc/models/gateway-port-vpn-path.md) | Optional | Property key is the VPN name |
| `WanArpPolicer` | [`*models.GatewayPortWanArpPolicerEnum`](../../doc/models/gateway-port-wan-arp-policer-enum.md) | Optional | Only when `wan_type`==`broadband`. enum: `default`, `max`, `recommended`<br>**Default**: `"default"` |
| `WanExtIp` | `*string` | Optional | Only if `usage`==`wan`, optional. If spoke should reach this port by a different IP |
| `WanExtraRoutes` | [`map[string]models.WanExtraRoutes`](../../doc/models/wan-extra-routes.md) | Optional | Only if `usage`==`wan`. Property Key is the destination CIDR (e.g "100.100.100.0/24") |
| `WanNetworks` | `[]string` | Optional | Only if `usage`==`wan`. If some networks are connected to this WAN port, it can be added here so policies can be defined |
| `WanProbeOverride` | [`*models.GatewayWanProbeOverride`](../../doc/models/gateway-wan-probe-override.md) | Optional | Only if `usage`==`wan` |
| `WanSourceNat` | [`*models.GatewayPortWanSourceNat`](../../doc/models/gateway-port-wan-source-nat.md) | Optional | Only if `usage`==`wan`, optional. By default, source-NAT is performed on all WAN Ports using the interface-ip |
| `WanType` | [`*models.GatewayPortWanTypeEnum`](../../doc/models/gateway-port-wan-type-enum.md) | Optional | Only if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`<br>**Default**: `"broadband"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "ae_idx": "ae_idx6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

