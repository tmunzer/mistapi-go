
# Gateway Port Config

Gateway port configuration for LAN, WAN, tunnel, and HA interfaces

## Structure

`GatewayPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AeDisableLacp` | `*bool` | Optional | If `aggregated`==`true`. To disable LCP support for the AE interface<br><br>**Default**: `false` |
| `AeIdx` | `models.Optional[string]` | Optional | If `aggregated`==`true`. Users could force to use the designated AE name (must be an integer between 0 and 127) |
| `AeLacpForceUp` | `*bool` | Optional | For SRX only, if `aggregated`==`true`.Sets the state of the interface as UP when the peer has limited LACP capability. Use case: When a device connected to this AE port is ZTPing for the first time, it will not have LACP configured on the other end. **Note:** Turning this on will enable force-up on one of the interfaces in the bundle only<br><br>**Default**: `false` |
| `Aggregated` | `*bool` | Optional | Whether the port participates in an aggregated Ethernet interface<br><br>**Default**: `false` |
| `Critical` | `*bool` | Optional | To generate port up/down alarm, set it to true<br><br>**Default**: `false` |
| `Description` | `*string` | Optional | Interface Description. Can be a variable (i.e. "{{myvar}}") |
| `DisableAutoneg` | `*bool` | Optional | Whether Ethernet autonegotiation is disabled on the port<br><br>**Default**: `false` |
| `Disabled` | `*bool` | Optional | Port admin up (true) / down (false)<br><br>**Default**: `false` |
| `DslType` | [`*models.GatewayPortDslTypeEnum`](../../doc/models/gateway-port-dsl-type-enum.md) | Optional | if `wan_type`==`dsl`. enum: `adsl`, `vdsl`<br><br>**Default**: `"vdsl"` |
| `DslVci` | `*int` | Optional | If `wan_type`==`dsl`, 16 bit int<br><br>**Default**: `35` |
| `DslVpi` | `*int` | Optional | If `wan_type`==`dsl`, 8 bit int<br><br>**Default**: `0` |
| `Duplex` | [`*models.GatewayPortDuplexEnum`](../../doc/models/gateway-port-duplex-enum.md) | Optional | enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `IpConfig` | [`*models.GatewayPortConfigIpConfig`](../../doc/models/gateway-port-config-ip-config.md) | Optional | Junos IP configuration for a gateway port interface |
| `LteApn` | `*string` | Optional | If `wan_type`==`lte`. APN used by the LTE uplink |
| `LteAuth` | [`*models.GatewayPortLteAuthEnum`](../../doc/models/gateway-port-lte-auth-enum.md) | Optional | if `wan_type`==`lte`. enum: `chap`, `none`, `pap`<br><br>**Default**: `"none"` |
| `LteBackup` | `*bool` | Optional | Whether the LTE uplink is used as a backup WAN connection |
| `LtePassword` | `*string` | Optional | If `wan_type`==`lte`. Password used for LTE uplink authentication |
| `LteUsername` | `*string` | Optional | If `wan_type`==`lte`. Username used for LTE uplink authentication |
| `Mtu` | `*int` | Optional | Layer 3 MTU configured on the port |
| `Name` | `*string` | Optional | Interface name used to derive device configuration |
| `Networks` | `[]string` | Optional | If `usage`==`lan`, name of the [networks]($h/Orgs%20Networks/_overview) to attach to the interface |
| `OuterVlanId` | `*int` | Optional | For Q-in-Q. Outer VLAN ID used for QinQ encapsulation |
| `PoeDisabled` | `*bool` | Optional | Whether PoE output is disabled on the port<br><br>**Default**: `false` |
| `PoeKeepStateWhenReboot` | `*bool` | Optional | Whether Perpetual PoE capabilities are enabled for a port<br><br>**Default**: `false` |
| `PortNetwork` | `*string` | Optional | Only for SRX and if `usage`==`lan`, the name of the Network to be used as the Untagged VLAN |
| `PreserveDscp` | `*bool` | Optional | Whether to preserve dscp when sending traffic over VPN (SSR-only)<br><br>**Default**: `true` |
| `Redundant` | `*bool` | Optional | If HA mode. Whether the port participates in the redundant Ethernet configuration |
| `RedundantGroup` | `*int` | Optional | If HA mode, SRX Only - support redundancy-group. 1-128 for physical SRX, 1-64 for virtual SRX<br><br>**Constraints**: `>= 1`, `<= 128` |
| `RethIdx` | [`*models.GatewayPortConfigRethIdx`](../../doc/models/containers/gateway-port-config-reth-idx.md) | Optional | For SRX only and if HA Mode. `-1` means it will be managed by the device. Use `>= 0` values to manage it manually. Ensure no conflicting values are assigned across all ports. |
| `RethNode` | `*string` | Optional | If HA mode. Node associated with the redundant Ethernet interface |
| `RethNodes` | `[]string` | Optional | SSR only - supporting vlan-based redundancy (matching the size of `networks`) |
| `Speed` | `*string` | Optional | Link speed configured on the port<br><br>**Default**: `"auto"` |
| `SsrNoVirtualMac` | `*bool` | Optional | When SSR is running as VM, this is required on certain hosting platforms<br><br>**Default**: `false` |
| `SvrPortRange` | `*string` | Optional | For SSR only. Port range configured on the interface<br><br>**Default**: `"none"` |
| `TrafficShaping` | [`*models.GatewayTrafficShaping`](../../doc/models/gateway-traffic-shaping.md) | Optional | Traffic shaping settings for a gateway interface or VPN path |
| `Usage` | [`models.GatewayPortUsageEnum`](../../doc/models/gateway-port-usage-enum.md) | Required | port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan` |
| `VlanId` | [`*models.GatewayPortVlanIdWithVariable`](../../doc/models/containers/gateway-port-vlan-id-with-variable.md) | Optional | If WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}") |
| `VpnPaths` | [`map[string]models.GatewayPortVpnPath`](../../doc/models/gateway-port-vpn-path.md) | Optional | Property key is the VPN name |
| `WanArpPolicer` | [`*models.GatewayPortWanArpPolicerEnum`](../../doc/models/gateway-port-wan-arp-policer-enum.md) | Optional | Only when `wan_type`==`broadband`. enum: `default`, `max`, `recommended`<br><br>**Default**: `"default"` |
| `WanExtIp` | `*string` | Optional | Only if `usage`==`wan`, optional. If spoke should reach this port by a different IP |
| `WanExtIp6` | `*string` | Optional | Only if `usage`==`wan`, optional. If spoke should reach this port by a different IPv6 |
| `WanExtraRoutes` | [`map[string]models.WanExtraRoutes`](../../doc/models/wan-extra-routes.md) | Optional | Only if `usage`==`wan`. Property Key is the destination CIDR (e.g. "100.100.100.0/24") |
| `WanExtraRoutes6` | [`map[string]models.WanExtraRoutes6`](../../doc/models/wan-extra-routes-6.md) | Optional | Only if `usage`==`wan`. Property Key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `WanNetworks` | `[]string` | Optional | Only if `usage`==`wan`. If some networks are connected to this WAN port, it can be added here so policies can be defined |
| `WanProbeOverride` | [`*models.GatewayWanProbeOverride`](../../doc/models/gateway-wan-probe-override.md) | Optional | Only if `usage`==`wan`. WAN health probe override for this gateway port |
| `WanSourceNat` | [`*models.GatewayPortWanSourceNat`](../../doc/models/gateway-port-wan-source-nat.md) | Optional | Only if `usage`==`wan`, optional. By default, source-NAT is performed on all WAN Ports using the interface-ip |
| `WanSpeedtestMode` | [`*models.GatewayPortConfigWanSpeedtestModeEnum`](../../doc/models/gateway-port-config-wan-speedtest-mode-enum.md) | Optional | Controls whether Marvis/scheduler can run speedtest on this port. enum: `auto`, `enabled`, `disabled`<br><br>**Default**: `"auto"` |
| `WanType` | [`*models.GatewayPortWanTypeEnum`](../../doc/models/gateway-port-wan-type-enum.md) | Optional | Only if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`<br><br>**Default**: `"broadband"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortConfig := models.GatewayPortConfig{
        AeDisableLacp:          models.ToPointer(false),
        AeIdx:                  models.NewOptional(models.ToPointer("ae_idx6")),
        AeLacpForceUp:          models.ToPointer(false),
        Aggregated:             models.ToPointer(false),
        Critical:               models.ToPointer(false),
        DisableAutoneg:         models.ToPointer(false),
        Disabled:               models.ToPointer(false),
        DslType:                models.ToPointer(models.GatewayPortDslTypeEnum_VDSL),
        DslVci:                 models.ToPointer(35),
        DslVpi:                 models.ToPointer(0),
        Duplex:                 models.ToPointer(models.GatewayPortDuplexEnum_FULL),
        LteAuth:                models.ToPointer(models.GatewayPortLteAuthEnum_NONE),
        PoeDisabled:            models.ToPointer(false),
        PoeKeepStateWhenReboot: models.ToPointer(false),
        PreserveDscp:           models.ToPointer(true),
        RethNodes:              []string{
            "node0",
            "node1",
        },
        Speed:                  models.ToPointer("1g"),
        SsrNoVirtualMac:        models.ToPointer(false),
        SvrPortRange:           models.ToPointer("60000-60005"),
        Usage:                  models.GatewayPortUsageEnum_LAN,
        WanArpPolicer:          models.ToPointer(models.GatewayPortWanArpPolicerEnum_ENUMDEFAULT),
        WanExtIp:               models.ToPointer("64.2.4.3"),
        WanExtIp6:              models.ToPointer("2601:1700:43c0:dc0::10"),
        WanSpeedtestMode:       models.ToPointer(models.GatewayPortConfigWanSpeedtestModeEnum_AUTO),
        WanType:                models.ToPointer(models.GatewayPortWanTypeEnum_BROADBAND),
    }

}
```

