
# Stats Gateway

Gateway statistics

## Structure

`StatsGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.ApRedundancy`](../../doc/models/ap-redundancy.md) | Optional | - |
| `ArpTableStats` | [`*models.ArpTableStats`](../../doc/models/arp-table-stats.md) | Optional | - |
| `BgpPeers` | [`[]models.OptionalStatsBgp`](../../doc/models/optional-stats-bgp.md) | Optional | only present when `bgp_peers` in `fields` query parameter<br>Each port object is same as `GET /api/v1/sites/:site_id/stats/bgp_peers/search` result object, except that org_id, site_id, mac, model are removed |
| `CertExpiry` | `*int64` | Optional | - |
| `ClusterConfig` | [`*models.StatsClusterConfig`](../../doc/models/stats-cluster-config.md) | Optional | - |
| `ClusterStat` | [`*models.StatsGatewayCluster`](../../doc/models/stats-gateway-cluster.md) | Optional | - |
| `ConductorName` | `*string` | Optional | - |
| `ConfigStatus` | `*string` | Optional | - |
| `Cpu2Stat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `Dhcpd2Stat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `DhcpdStat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional | - |
| `ExtIp` | `models.Optional[string]` | Optional | IP address |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `HasPcap` | `models.Optional[bool]` | Optional | - |
| `Hostname` | `*string` | Optional | hostname reported by the device |
| `Id` | `*uuid.UUID` | Optional | serial |
| `If2Stat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `IfStat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `Ip` | `models.Optional[string]` | Optional | IP address |
| `Ip2Stat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `IsHa` | `models.Optional[bool]` | Optional | - |
| `LastSeen` | `*float64` | Optional | last seen timestamp |
| `Mac` | `string` | Required | device mac |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | serial |
| `Memory2Stat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `Model` | `*string` | Optional | device model |
| `ModifiedTime` | `*float64` | Optional | - |
| `Module2Stat` | [`[]models.ModuleStatItem`](../../doc/models/module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `ModuleStat` | [`[]models.ModuleStatItem`](../../doc/models/module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Name` | `*string` | Optional | device name if configured |
| `NodeName` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | serial |
| `Ports` | [`[]models.OptionalStatsPort`](../../doc/models/optional-stats-port.md) | Optional | only present when `ports` in `fields` query parameter<br>Each port object is same as `GET /api/v1/sites/:site_id/stats/ports/search` result object, except that org_id, site_id, mac, model are removed |
| `RouteSummaryStats` | [`*models.RouteSummaryStats`](../../doc/models/route-summary-stats.md) | Optional | - |
| `RouterName` | `*string` | Optional | device name if configured |
| `Serial` | `*string` | Optional | serial |
| `Service2Stat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | - |
| `ServiceStat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | - |
| `ServiceStatus` | [`*models.StatsGatewayServiceStatus`](../../doc/models/stats-gateway-service-status.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | serial |
| `Spu2Stat` | [`[]models.StatsGatewaySpuItem`](../../doc/models/stats-gateway-spu-item.md) | Optional | - |
| `SpuStat` | [`[]models.StatsGatewaySpuItem`](../../doc/models/stats-gateway-spu-item.md) | Optional | - |
| `Status` | `*string` | Optional | - |
| `Tunnels` | [`[]models.OptionalStatWanTunnel`](../../doc/models/optional-stat-wan-tunnel.md) | Optional | only present when `tunnels` in `fields` query parameter<br>Each port object is same as `GET /api/v1/sites/:site_id/stats/tunnels/search` result object, except that org_id, site_id, mac, model are removed |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br>**Default**: `"gateway"` |
| `Uptime` | `*float64` | Optional | - |
| `Version` | `*string` | Optional | - |
| `VpnPeers` | [`[]models.OptionalStatVpnPeer`](../../doc/models/optional-stat-vpn-peer.md) | Optional | only present when `vpn_peers` in `fields` query parameter<br>Each port object is same as `GET /api/v1/sites/:site_id/stats/vpn_peers/search` result object, except that org_id, site_id, mac, model are removed |

## Example (as JSON)

```json
{
  "ext_ip": "66.129.234.224",
  "hostname": "sj1",
  "ip": "10.2.11.137",
  "last_seen": 1553203563,
  "mac": "dc38e1dbf3cd",
  "model": "SRX320",
  "name": "sj1",
  "node_name": "node0",
  "router_name": "sj1",
  "serial": "TC3714190003",
  "status": "connected",
  "type": "gateway",
  "uptime": 3671219,
  "version": "18.4R1.8",
  "ap_redundancy": {
    "modules": {
      "key0": {
        "num_aps": 2,
        "num_aps_with_switch_redundancy": 254
      }
    },
    "num_aps": 246,
    "num_aps_with_switch_redundancy": 10
  },
  "arp_table_stats": {
    "arp_table_count": 136,
    "max_entries_supported": 8
  },
  "bgp_peers": [
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": 146,
      "neighbor": "neighbor6",
      "neighbor_as": 172
    },
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": 146,
      "neighbor": "neighbor6",
      "neighbor_as": 172
    },
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": 146,
      "neighbor": "neighbor6",
      "neighbor_as": 172
    }
  ],
  "cert_expiry": 52,
  "cluster_config": {
    "configuration": "configuration0",
    "control_link_info": {
      "name": "name0",
      "status": "status2"
    },
    "ethernet_connection": [
      {
        "name": "name2",
        "status": "status6"
      }
    ],
    "fabric_link_info": {
      "DataPlaneNotifiedStatus": "DataPlaneNotifiedStatus8",
      "Interface": [
        "Interface0",
        "Interface1",
        "Interface2"
      ],
      "InternalStatus": "InternalStatus2",
      "State": "State4",
      "Status": "Status0"
    },
    "last_status_change_reason": "last_status_change_reason8"
  }
}
```

