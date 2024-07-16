
# Switch Stats

Switch statistics

## Structure

`SwitchStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.SwitchStatsApRedundancy`](../../doc/models/switch-stats-ap-redundancy.md) | Optional | - |
| `ArpTableStats` | [`*models.ArpTableStats`](../../doc/models/arp-table-stats.md) | Optional | - |
| `CertExpiry` | `*int64` | Optional | - |
| `Clients` | [`[]models.SwitchStatsClientItem`](../../doc/models/switch-stats-client-item.md) | Optional | - |
| `ClientsStats` | [`*models.SwitchStatsClientsStats`](../../doc/models/switch-stats-clients-stats.md) | Optional | - |
| `ConfigStatus` | `*string` | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `DhcpdStat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional | - |
| `FwVersionsOutofsync` | `*bool` | Optional | - |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `HasPcap` | `*bool` | Optional | whether the switch supports packet capture |
| `Hostname` | `*string` | Optional | hostname reported by the device |
| `HwRev` | `*string` | Optional | device hardware revision number |
| `Id` | `*uuid.UUID` | Optional | - |
| `IfStat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `Ip` | `*string` | Optional | - |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `LastSeen` | `*float64` | Optional | - |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | last trouble code of switch |
| `Mac` | `*string` | Optional | - |
| `MacTableStats` | [`*models.MacTableStats`](../../doc/models/mac-table-stats.md) | Optional | - |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | - |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `Model` | `*string` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `ModuleStat` | [`[]models.ModuleStatItem`](../../doc/models/module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Name` | `*string` | Optional | device name if configured |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RouteSummaryStats` | [`*models.RouteSummaryStats`](../../doc/models/route-summary-stats.md) | Optional | - |
| `Serial` | `*string` | Optional | - |
| `ServiceStat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Uptime` | `*float64` | Optional | - |
| `VcMac` | `models.Optional[string]` | Optional | - |
| `VcSetupInfo` | [`*models.SwitchStatsVcSetupInfo`](../../doc/models/switch-stats-vc-setup-info.md) | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "has_pcap": false,
  "hostname": "sj-sw1",
  "ip": "10.2.11.137",
  "last_seen": 1553203563,
  "model": "EX4600",
  "name": "sj-sw1",
  "serial": "TC3714190003",
  "status": "connected",
  "type": "switch",
  "uptime": 13501,
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
  "cert_expiry": 18,
  "clients": [
    {
      "device_mac": "device_mac2",
      "hostname": "hostname6",
      "mac": "mac2",
      "port_id": "port_id8"
    },
    {
      "device_mac": "device_mac2",
      "hostname": "hostname6",
      "mac": "mac2",
      "port_id": "port_id8"
    },
    {
      "device_mac": "device_mac2",
      "hostname": "hostname6",
      "mac": "mac2",
      "port_id": "port_id8"
    }
  ],
  "clients_stats": {
    "total": {
      "num_aps": [
        23,
        22,
        21
      ],
      "num_wired_clients": 222
    }
  }
}
```

