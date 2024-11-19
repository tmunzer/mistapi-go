
# Evpn Topology Switch Config

## Structure

`EvpnTopologySwitchConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DhcpdConfig` | [`*models.EvpnTopologySwitchConfigDhcpdConfig`](../../doc/models/evpn-topology-switch-config-dhcpd-config.md) | Optional | - |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `OtherIpConfigs` | [`map[string]models.JunosOtherIpConfig`](../../doc/models/junos-other-ip-config.md) | Optional | additional IP Addresses configured on the switch. Property key is the port network name |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `RouterId` | `*string` | Optional | used for OSPF / BGP / EVPN |
| `VrfConfig` | [`*models.EvpnTopologySwitchConfigVrfConfig`](../../doc/models/evpn-topology-switch-config-vrf-config.md) | Optional | - |

## Example (as JSON)

```json
{
  "router_id": "10.2.1.10",
  "dhcpd_config": {
    "enabled": false
  },
  "networks": {
    "key0": {
      "isolation": false,
      "isolation_vlan_id": "isolation_vlan_id8",
      "subnet": "subnet6",
      "vlan_id": "String7"
    },
    "key1": {
      "isolation": false,
      "isolation_vlan_id": "isolation_vlan_id8",
      "subnet": "subnet6",
      "vlan_id": "String7"
    },
    "key2": {
      "isolation": false,
      "isolation_vlan_id": "isolation_vlan_id8",
      "subnet": "subnet6",
      "vlan_id": "String7"
    }
  },
  "other_ip_configs": {
    "key0": {
      "evpn_anycast": false,
      "ip": "ip4",
      "ip6": "ip60",
      "netmask": "netmask0",
      "netmask6": "netmask60"
    }
  },
  "port_config": {
    "key0": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6"
    },
    "key1": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6"
    },
    "key2": {
      "ae_disable_lacp": false,
      "ae_idx": 230,
      "ae_lacp_slow": false,
      "aggregated": false,
      "critical": false,
      "usage": "usage6"
    }
  },
  "port_usages": {
    "key0": {
      "all_networks": false,
      "allow_dhcpd": false,
      "allow_multiple_supplicants": false,
      "bypass_auth_when_server_down": false,
      "bypass_auth_when_server_down_for_unkonwn_client": false
    }
  }
}
```
