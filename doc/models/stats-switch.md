
# Stats Switch

Switch statistics reported by Mist for a site or organization stats response

*This model accepts additional fields of type interface{}.*

## Structure

`StatsSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.StatsSwitchApRedundancy`](../../doc/models/stats-switch-ap-redundancy.md) | Optional | AP switch redundancy coverage summary for a switch |
| `ArpTableStats` | [`*models.ArpTableStats`](../../doc/models/arp-table-stats.md) | Optional | ARP table usage and capacity statistics |
| `AutoUpgradeStat` | [`*models.StatsApAutoUpgrade`](../../doc/models/stats-ap-auto-upgrade.md) | Optional | Auto-upgrade status for an AP |
| `CertExpiry` | `*int64` | Optional | Time when the switch certificate expires, in epoch seconds |
| `Clients` | [`[]models.StatsSwitchClientItem`](../../doc/models/stats-switch-client-item.md) | Optional | Clients observed on switch ports in switch statistics |
| `ClientsStats` | [`*models.StatsSwitchClientsStats`](../../doc/models/stats-switch-clients-stats.md) | Optional | Aggregate switch client counts |
| `ConfigStatus` | `*string` | Optional, Read-only | Configuration synchronization status reported for the switch |
| `ConfigTimestamp` | `*int` | Optional | Time when the switch configuration status was last updated, in epoch seconds |
| `ConfigVersion` | `*int` | Optional | Currently applied configuration version for the switch |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | CPU utilization breakdown for a device |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Applied device profile identifier for the switch, when present |
| `DhcpdStat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Associated EVPN topology identifier for the switch, when present |
| `ExtIp` | `*string` | Optional | Public IP address observed for the switch |
| `FwVersionsOutofsync` | `*bool` | Optional, Read-only | Whether firmware versions are out of sync across switch members |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | Firmware update status for a device |
| `HasPcap` | `*bool` | Optional, Read-only | Whether the switch supports packet capture |
| `Hostname` | `*string` | Optional, Read-only | Device-reported hostname for the switch |
| `HwRev` | `*string` | Optional | Device hardware revision number |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IfStat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `Ip` | `*string` | Optional, Read-only | Management IP address reported for the switch |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | Read-only IP addressing status reported by a device interface |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | Last trouble indicator reported by a switch |
| `Mac` | `*string` | Optional, Read-only | Switch MAC address reported by Mist |
| `MacTableStats` | [`*models.MacTableStats`](../../doc/models/mac-table-stats.md) | Optional | MAC table capacity and usage statistics |
| `MapId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Placement map identifier associated with the switch, when present |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory utilization statistics for a device; in a virtual chassis, this reports the master Routing Engine |
| `Model` | `*string` | Optional, Read-only | Switch model name reported by Mist |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `ModuleStat` | [`[]models.StatsSwitchModuleStatItem`](../../doc/models/stats-switch-module-stat-item.md) | Optional | Hardware module statistics reported by a switch<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Name` | `*string` | Optional, Read-only | Device name if configured |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Ports` | [`[]models.StatsSwitchPort`](../../doc/models/stats-switch-port.md) | Optional | Switch port statistics records returned by a stats response |
| `RouteSummaryStats` | [`*models.RouteSummaryStats`](../../doc/models/route-summary-stats.md) | Optional | Route table capacity and usage summary |
| `Serial` | `*string` | Optional, Read-only | Switch serial number reported by Mist |
| `ServiceStat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | Service version information keyed by service name |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Status` | `*string` | Optional, Read-only | Connection status reported for the switch, such as connected |
| `TagId` | `*int` | Optional | Numeric inventory tag identifier associated with the switch |
| `TagUuid` | `*uuid.UUID` | Optional | Inventory tag UUID associated with the switch |
| `Type` | `string` | Required, Constant | Device Type. enum: `switch`<br><br>**Value**: `"switch"` |
| `Uptime` | `models.Optional[float64]` | Optional, Read-only | Elapsed time since the switch last booted, in seconds |
| `VcMac` | `models.Optional[string]` | Optional, Read-only | Virtual Chassis MAC address reported for the switch, when present |
| `VcSetupInfo` | [`*models.StatsSwitchVcSetupInfo`](../../doc/models/stats-switch-vc-setup-info.md) | Optional | Virtual Chassis setup request and status details reported by a switch |
| `Version` | `models.Optional[string]` | Optional, Read-only | Software version running on the switch |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsSwitch := models.StatsSwitch{
        ApRedundancy:         models.ToPointer(models.StatsSwitchApRedundancy{
            Modules:                    map[string]models.StatsSwitchApRedundancyModule{
                "key0": models.StatsSwitchApRedundancyModule{
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
        CertExpiry:           models.ToPointer(int64(0)),
        Clients:              []models.StatsSwitchClientItem{
            models.StatsSwitchClientItem{
                DeviceMac:            models.ToPointer("device_mac2"),
                Hostname:             models.ToPointer("hostname6"),
                Mac:                  models.ToPointer("mac2"),
                PortId:               models.ToPointer("port_id8"),
            },
            models.StatsSwitchClientItem{
                DeviceMac:            models.ToPointer("device_mac2"),
                Hostname:             models.ToPointer("hostname6"),
                Mac:                  models.ToPointer("mac2"),
                PortId:               models.ToPointer("port_id8"),
            },
            models.StatsSwitchClientItem{
                DeviceMac:            models.ToPointer("device_mac2"),
                Hostname:             models.ToPointer("hostname6"),
                Mac:                  models.ToPointer("mac2"),
                PortId:               models.ToPointer("port_id8"),
            },
        },
        HasPcap:              models.ToPointer(false),
        Hostname:             models.ToPointer("sj-sw1"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Ip:                   models.ToPointer("10.2.11.137"),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Model:                models.ToPointer("EX4600"),
        Name:                 models.ToPointer("sj-sw1"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Serial:               models.ToPointer("TC3714190003"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Status:               models.ToPointer("connected"),
        Type:                 "switch",
        Uptime:               models.NewOptional(models.ToPointer(float64(13501))),
        Version:              models.NewOptional(models.ToPointer("18.4R1.8")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

