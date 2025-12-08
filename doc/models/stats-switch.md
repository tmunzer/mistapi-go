
# Stats Switch

Switch statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.StatsSwitchApRedundancy`](../../doc/models/stats-switch-ap-redundancy.md) | Optional | - |
| `ArpTableStats` | [`*models.ArpTableStats`](../../doc/models/arp-table-stats.md) | Optional | - |
| `AutoUpgradeStat` | [`*models.StatsApAutoUpgrade`](../../doc/models/stats-ap-auto-upgrade.md) | Optional | - |
| `CertExpiry` | `*int64` | Optional | - |
| `Clients` | [`[]models.StatsSwitchClientItem`](../../doc/models/stats-switch-client-item.md) | Optional | - |
| `ClientsStats` | [`*models.StatsSwitchClientsStats`](../../doc/models/stats-switch-clients-stats.md) | Optional | - |
| `ConfigStatus` | `*string` | Optional | - |
| `ConfigTimestamp` | `*int` | Optional | - |
| `ConfigVersion` | `*int` | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `DhcpdStat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional | - |
| `ExtIp` | `*string` | Optional | - |
| `FwVersionsOutofsync` | `*bool` | Optional | - |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `HasPcap` | `*bool` | Optional | Whether the switch supports packet capture |
| `Hostname` | `*string` | Optional | Hostname reported by the device |
| `HwRev` | `*string` | Optional | Device hardware revision number |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IfStat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `Ip` | `*string` | Optional | - |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | Last trouble code of switch |
| `Mac` | `*string` | Optional | - |
| `MacTableStats` | [`*models.MacTableStats`](../../doc/models/mac-table-stats.md) | Optional | - |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | - |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory usage stat (for virtual chassis, memory usage of master RE) |
| `Model` | `*string` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `ModuleStat` | [`[]models.StatsSwitchModuleStatItem`](../../doc/models/stats-switch-module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Name` | `*string` | Optional | Device name if configured |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Ports` | [`[]models.StatsSwitchPort`](../../doc/models/stats-switch-port.md) | Optional | - |
| `RouteSummaryStats` | [`*models.RouteSummaryStats`](../../doc/models/route-summary-stats.md) | Optional | - |
| `Serial` | `*string` | Optional | - |
| `ServiceStat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | `*string` | Optional | - |
| `TagId` | `*int` | Optional | - |
| `TagUuid` | `*uuid.UUID` | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `switch`<br><br>**Value**: `"switch"` |
| `Uptime` | `models.Optional[float64]` | Optional | - |
| `VcMac` | `models.Optional[string]` | Optional | - |
| `VcSetupInfo` | [`*models.StatsSwitchVcSetupInfo`](../../doc/models/stats-switch-vc-setup-info.md) | Optional | - |
| `Version` | `models.Optional[string]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "has_pcap": false,
  "hostname": "sj-sw1",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "ip": "10.2.11.137",
  "last_seen": 1470417522,
  "model": "EX4600",
  "name": "sj-sw1",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "serial": "TC3714190003",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
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
  "auto_upgrade_stat": {
    "lastcheck": 28
  },
  "cert_expiry": 122,
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

