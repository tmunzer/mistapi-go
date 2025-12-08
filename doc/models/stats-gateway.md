
# Stats Gateway

Gateway statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.ApRedundancy`](../../doc/models/ap-redundancy.md) | Optional | - |
| `ArpTableStats` | [`*models.ArpTableStats`](../../doc/models/arp-table-stats.md) | Optional | - |
| `AutoUpgradeStat` | [`*models.StatsApAutoUpgrade`](../../doc/models/stats-ap-auto-upgrade.md) | Optional | - |
| `BgpPeers` | [`[]models.BgpPeer`](../../doc/models/bgp-peer.md) | Optional | Only present when `bgp_peers` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/bgp_peers/search` result object, except that org_id, site_id, mac, model are removed |
| `CertExpiry` | `*int64` | Optional | - |
| `ClusterConfig` | [`*models.StatsClusterConfig`](../../doc/models/stats-cluster-config.md) | Optional | - |
| `ClusterStat` | [`*models.StatsGatewayCluster`](../../doc/models/stats-gateway-cluster.md) | Optional | - |
| `ConductorName` | `*string` | Optional | - |
| `ConfigStatus` | `*string` | Optional | - |
| `ConfigTimestamp` | `*int` | Optional | - |
| `ConfigVersion` | `*int` | Optional | - |
| `Cpu2Stat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `DeviceprofileName` | `*string` | Optional | - |
| `Dhcpd2Stat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `DhcpdStat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional | - |
| `ExtIp` | `models.Optional[string]` | Optional | IP address |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `HasPcap` | `models.Optional[bool]` | Optional | - |
| `Hostname` | `*string` | Optional | Hostname reported by the device |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `If2Stat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `IfStat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `Ip` | `models.Optional[string]` | Optional | IP address |
| `Ip2Stat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `IsHa` | `models.Optional[bool]` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `string` | Required | Device mac |
| `MacTableStats` | [`*models.StatsGatewayMacTableStats`](../../doc/models/stats-gateway-mac-table-stats.md) | Optional | - |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | Serial Number |
| `Memory2Stat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory usage stat (for virtual chassis, memory usage of master RE) |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory usage stat (for virtual chassis, memory usage of master RE) |
| `Model` | `*string` | Optional | Device model |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Module2Stat` | [`[]models.StatsGatewayModuleStatItem`](../../doc/models/stats-gateway-module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `0`, *Maximum Items*: `1` |
| `ModuleStat` | [`[]models.StatsGatewayModuleStatItem`](../../doc/models/stats-gateway-module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `0`, *Maximum Items*: `1` |
| `Name` | `*string` | Optional | Device name if configured |
| `NodeName` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Ports` | [`[]models.StatsGatewayPort`](../../doc/models/stats-gateway-port.md) | Optional | Only present when `ports` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/ports/search` result object, except that org_id, site_id, mac, model are removed |
| `RouteSummaryStats` | [`*models.RouteSummaryStats`](../../doc/models/route-summary-stats.md) | Optional | - |
| `RouterName` | `*string` | Optional | Device name if configured |
| `Serial` | `*string` | Optional | Serial Number |
| `Service2Stat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | - |
| `ServiceStat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | - |
| `ServiceStatus` | [`*models.StatsGatewayServiceStatus`](../../doc/models/stats-gateway-service-status.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Spu2Stat` | [`[]models.StatsGatewaySpuItem`](../../doc/models/stats-gateway-spu-item.md) | Optional | - |
| `SpuStat` | [`[]models.StatsGatewaySpuItem`](../../doc/models/stats-gateway-spu-item.md) | Optional | - |
| `Status` | `*string` | Optional | - |
| `TagId` | `*int` | Optional | - |
| `TagUuid` | `*uuid.UUID` | Optional | - |
| `Tunnels` | [`[]models.StatsGatewayWanTunnel`](../../doc/models/stats-gateway-wan-tunnel.md) | Optional | Only present when `tunnels` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/tunnels/search` result object, except that org_id, site_id, mac, model are removed |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br><br>**Value**: `"gateway"` |
| `Uptime` | `models.Optional[float64]` | Optional | - |
| `Version` | `models.Optional[string]` | Optional | - |
| `VpnPeers` | [`[]models.StatsGatewayVpnPeer`](../../doc/models/stats-gateway-vpn-peer.md) | Optional | Only present when `vpn_peers` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/vpn_peers/search` result object, except that org_id, site_id, mac, model are removed |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ext_ip": "66.129.234.224",
  "hostname": "sj1",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "ip": "10.2.11.137",
  "last_seen": 1470417522,
  "mac": "dc38e1dbf3cd",
  "model": "SRX320",
  "name": "sj1",
  "node_name": "node0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "router_name": "sj1",
  "serial": "TC3714190003",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
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
  "auto_upgrade_stat": {
    "lastcheck": 28
  },
  "bgp_peers": [
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": "String3",
      "neighbor": "neighbor6",
      "neighbor_as": "String9"
    },
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": "String3",
      "neighbor": "neighbor6",
      "neighbor_as": "String9"
    },
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": "String3",
      "neighbor": "neighbor6",
      "neighbor_as": "String9"
    }
  ],
  "cert_expiry": 52,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

