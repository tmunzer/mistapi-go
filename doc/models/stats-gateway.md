
# Stats Gateway

Gateway statistics reported by Mist for a site or organization stats response

*This model accepts additional fields of type interface{}.*

## Structure

`StatsGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.ApRedundancy`](../../doc/models/ap-redundancy.md) | Optional | AP switch redundancy coverage summary |
| `ArpTableStats` | [`*models.ArpTableStats`](../../doc/models/arp-table-stats.md) | Optional | ARP table usage and capacity statistics |
| `AutoUpgradeStat` | [`*models.StatsApAutoUpgrade`](../../doc/models/stats-ap-auto-upgrade.md) | Optional | Auto-upgrade status for an AP |
| `BgpPeers` | [`[]models.BgpPeer`](../../doc/models/bgp-peer.md) | Optional | Only present when `bgp_peers` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/bgp_peers/search` result object, except that org_id, site_id, mac, model are removed |
| `CertExpiry` | `*int64` | Optional | Time when the gateway certificate expires, in epoch seconds |
| `ClusterConfig` | [`*models.StatsClusterConfig`](../../doc/models/stats-cluster-config.md) | Optional | High-availability cluster configuration and health reported by a gateway |
| `ClusterStat` | [`*models.StatsGatewayCluster`](../../doc/models/stats-gateway-cluster.md) | Optional | High-availability cluster state reported by a gateway |
| `ConductorName` | `*string` | Optional, Read-only | SSR conductor name associated with the gateway, when applicable |
| `ConfigStatus` | `*string` | Optional, Read-only | Configuration synchronization status reported for the gateway |
| `ConfigTimestamp` | `*int` | Optional | Time when the gateway configuration status was last updated, in epoch seconds |
| `ConfigVersion` | `*int` | Optional | Currently applied configuration version for the gateway |
| `Cpu2Stat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | CPU utilization breakdown for a device |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | CPU utilization breakdown for a device |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Applied device profile identifier for the gateway, when present |
| `DeviceprofileName` | `*string` | Optional | Applied device profile name for the gateway |
| `Dhcpd2Stat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `DhcpdStat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Associated EVPN topology identifier for the gateway, when present |
| `ExtIp` | `models.Optional[string]` | Optional, Read-only | Public IP address observed for the gateway |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | Firmware update status for a device |
| `HasPcap` | `models.Optional[bool]` | Optional, Read-only | Whether packet capture is available for the gateway |
| `Hostname` | `*string` | Optional | Device-reported hostname for the gateway |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `If2Stat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `IfStat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `Ip` | `models.Optional[string]` | Optional, Read-only | Management IP address reported for the gateway |
| `Ip2Stat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | Read-only IP addressing status reported by a device interface |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | Read-only IP addressing status reported by a device interface |
| `IsHa` | `models.Optional[bool]` | Optional, Read-only | Whether the gateway is part of an HA cluster |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `string` | Required | Gateway MAC address reported by Mist |
| `MacTableStats` | [`*models.StatsGatewayMacTableStats`](../../doc/models/stats-gateway-mac-table-stats.md) | Optional | Gateway MAC table utilization counters |
| `MapId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Placement map identifier associated with the gateway, when present |
| `Memory2Stat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory utilization statistics for a device; in a virtual chassis, this reports the master Routing Engine |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory utilization statistics for a device; in a virtual chassis, this reports the master Routing Engine |
| `Model` | `*string` | Optional | Gateway model name reported by Mist |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Module2Stat` | [`[]models.StatsGatewayModuleStatItem`](../../doc/models/stats-gateway-module-stat-item.md) | Optional | Hardware module statistics reported by a gateway<br><br>**Constraints**: *Minimum Items*: `0`, *Maximum Items*: `1` |
| `ModuleStat` | [`[]models.StatsGatewayModuleStatItem`](../../doc/models/stats-gateway-module-stat-item.md) | Optional | Hardware module statistics reported by a gateway<br><br>**Constraints**: *Minimum Items*: `0`, *Maximum Items*: `1` |
| `Name` | `*string` | Optional, Read-only | Device name if configured |
| `NodeName` | `*string` | Optional, Read-only | HA node name for the gateway, such as node0 or node1 |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Ports` | [`[]models.StatsGatewayPort`](../../doc/models/stats-gateway-port.md) | Optional | Only present when `ports` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/ports/search` result object, except that org_id, site_id, mac, model are removed |
| `RouteSummaryStats` | [`*models.RouteSummaryStats`](../../doc/models/route-summary-stats.md) | Optional | Route table capacity and usage summary |
| `RouterName` | `*string` | Optional, Read-only | Device name if configured |
| `Serial` | `*string` | Optional, Read-only | Gateway serial number reported by Mist |
| `Service2Stat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | Service version information keyed by service name |
| `ServiceStat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | Service version information keyed by service name |
| `ServiceStatus` | [`*models.StatsGatewayServiceStatus`](../../doc/models/stats-gateway-service-status.md) | Optional | Gateway security service installation and runtime status |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Spu2Stat` | [`[]models.StatsGatewaySpuItem`](../../doc/models/stats-gateway-spu-item.md) | Optional | Services Processing Unit statistics reported by a gateway |
| `SpuStat` | [`[]models.StatsGatewaySpuItem`](../../doc/models/stats-gateway-spu-item.md) | Optional | Services Processing Unit statistics reported by a gateway |
| `Status` | `*string` | Optional, Read-only | Connection status reported for the gateway, such as connected |
| `TagId` | `*int` | Optional | Numeric inventory tag identifier associated with the gateway |
| `TagUuid` | `*uuid.UUID` | Optional | Inventory tag UUID associated with the gateway |
| `Tunnels` | [`[]models.StatsGatewayWanTunnel`](../../doc/models/stats-gateway-wan-tunnel.md) | Optional | Only present when `tunnels` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/tunnels/search` result object, except that org_id, site_id, mac, model are removed |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `gateway`<br><br>**Value**: `"gateway"` |
| `Uptime` | `models.Optional[float64]` | Optional, Read-only | Elapsed time since the gateway last booted, in seconds |
| `Version` | `models.Optional[string]` | Optional, Read-only | Software version running on the gateway |
| `VpnPeers` | [`[]models.StatsGatewayVpnPeer`](../../doc/models/stats-gateway-vpn-peer.md) | Optional | Only present when `vpn_peers` in `fields` query parameter. Each port object is same as `GET /api/v1/sites/{site_id}/stats/vpn_peers/search` result object, except that org_id, site_id, mac, model are removed |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsGateway := models.StatsGateway{
        ApRedundancy:         models.ToPointer(models.ApRedundancy{
            Modules:                    map[string]models.ApRedundancyModule{
                "key0": models.ApRedundancyModule{
                    NumAps:                     models.ToPointer(2),
                    NumApsWithSwitchRedundancy: models.ToPointer(254),
                },
            },
            NumAps:                     models.ToPointer(246),
            NumApsWithSwitchRedundancy: models.ToPointer(10),
        }),
        ArpTableStats:        models.ToPointer(models.ArpTableStats{
            ArpTableCount:        models.ToPointer(136),
            MaxEntriesSupported:  models.ToPointer(8),
        }),
        AutoUpgradeStat:      models.ToPointer(models.StatsApAutoUpgrade{
            Lastcheck:            models.NewOptional(models.ToPointer(int64(28))),
        }),
        BgpPeers:             []models.BgpPeer{
            models.BgpPeer{
                EvpnOverlay:          models.ToPointer(false),
                ForOverlay:           models.ToPointer(false),
                LocalAs:              models.ToPointer(models.BgpAsContainer.FromString("String3")),
                Neighbor:             models.ToPointer("neighbor6"),
                NeighborAs:           models.ToPointer(models.BgpAsContainer.FromString("String9")),
            },
        },
        CertExpiry:           models.ToPointer(int64(206)),
        ExtIp:                models.NewOptional(models.ToPointer("66.129.234.224")),
        Hostname:             models.ToPointer("sj1"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Ip:                   models.NewOptional(models.ToPointer("10.2.11.137")),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  "dc38e1dbf3cd",
        Model:                models.ToPointer("SRX320"),
        Name:                 models.ToPointer("sj1"),
        NodeName:             models.ToPointer("node0"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RouterName:           models.ToPointer("sj1"),
        Serial:               models.ToPointer("TC3714190003"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Status:               models.ToPointer("connected"),
        Type:                 "gateway",
        Uptime:               models.NewOptional(models.ToPointer(float64(3671219))),
        Version:              models.NewOptional(models.ToPointer("18.4R1.8")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

