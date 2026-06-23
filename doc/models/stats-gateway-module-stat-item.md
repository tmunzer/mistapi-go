
# Stats Gateway Module Stat Item

Hardware module status and firmware inventory for a gateway

## Structure

`StatsGatewayModuleStatItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BackupVersion` | `models.Optional[string]` | Optional, Read-only | Backup software version stored on the gateway module |
| `BiosVersion` | `models.Optional[string]` | Optional, Read-only | BIOS firmware version reported for the gateway module |
| `BootPartition` | `*string` | Optional | Active boot partition used by the gateway module |
| `CpldVersion` | `models.Optional[string]` | Optional, Read-only | CPLD firmware version reported for the gateway module |
| `Fans` | [`[]models.ModuleStatItemFansItems`](../../doc/models/module-stat-item-fans-items.md) | Optional | Cooling fan status records for a device module<br><br>**Constraints**: *Unique Items Required* |
| `FpgaVersion` | `models.Optional[string]` | Optional, Read-only | FPGA firmware version reported for the gateway module |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Locating` | `*bool` | Optional | Whether the gateway module locator indicator is active |
| `Mac` | `*string` | Optional | Gateway module MAC address reported by Mist |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory utilization statistics for a device; in a virtual chassis, this reports the master Routing Engine |
| `Model` | `models.Optional[string]` | Optional, Read-only | Gateway module model name reported by Mist |
| `NetworkResources` | [`[]models.ModuleStatItemNetworkResource`](../../doc/models/module-stat-item-network-resource.md) | Optional | Network resource usage counters reported by a device module |
| `OpticsCpldVersion` | `models.Optional[string]` | Optional, Read-only | Optics CPLD firmware version reported for the gateway module |
| `PendingVersion` | `models.Optional[string]` | Optional, Read-only | Pending software version staged for the gateway module |
| `Poe` | [`*models.ModuleStatItemPoe`](../../doc/models/module-stat-item-poe.md) | Optional | Power over Ethernet telemetry for a device module |
| `PoeVersion` | `models.Optional[string]` | Optional, Read-only | PoE controller firmware version reported for the gateway module |
| `PowerCpldVersion` | `models.Optional[string]` | Optional, Read-only | Power CPLD firmware version reported for the gateway module |
| `Psus` | [`[]models.ModuleStatItemPsusItem`](../../doc/models/module-stat-item-psus-item.md) | Optional | Power supply status records for a device module<br><br>**Constraints**: *Unique Items Required* |
| `ReFpgaVersion` | `models.Optional[string]` | Optional, Read-only | Routing Engine FPGA firmware version reported for the gateway module |
| `RecoveryVersion` | `models.Optional[string]` | Optional, Read-only | Recovery software version stored on the gateway module |
| `Serial` | `models.Optional[string]` | Optional, Read-only | Gateway module serial number reported by Mist |
| `Status` | `models.Optional[string]` | Optional, Read-only | Connection status reported for the gateway module |
| `Temperatures` | [`[]models.ModuleStatItemTemperaturesItem`](../../doc/models/module-stat-item-temperatures-item.md) | Optional | Temperature sensor readings for a device module<br><br>**Constraints**: *Unique Items Required* |
| `TmcFpgaVersion` | `models.Optional[string]` | Optional, Read-only | TMC FPGA firmware version reported for the gateway module |
| `UbootVersion` | `models.Optional[string]` | Optional, Read-only | U-Boot firmware version reported for the gateway module |
| `Uptime` | `models.Optional[int]` | Optional, Read-only | Elapsed time since the gateway module last booted, in seconds |
| `VcLinks` | [`[]models.ModuleStatItemVcLinksItem`](../../doc/models/module-stat-item-vc-links-item.md) | Optional | Virtual chassis link records for a device module<br><br>**Constraints**: *Unique Items Required* |
| `VcMode` | `models.Optional[string]` | Optional, Read-only | Virtual chassis mode reported for the gateway module |
| `VcRole` | `models.Optional[string]` | Optional, Read-only | Virtual chassis role reported for the gateway module. enum: `master`, `backup`, `linecard` |
| `VcState` | `models.Optional[string]` | Optional, Read-only | Virtual chassis state reported for the gateway module |
| `Version` | `models.Optional[string]` | Optional, Read-only | Software version running on the gateway module |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsGatewayModuleStatItem := models.StatsGatewayModuleStatItem{
        BackupVersion:        models.NewOptional(models.ToPointer("backup_version0")),
        BiosVersion:          models.NewOptional(models.ToPointer("bios_version4")),
        BootPartition:        models.ToPointer("boot_partition0"),
        CpldVersion:          models.NewOptional(models.ToPointer("cpld_version6")),
        Fans:                 []models.ModuleStatItemFansItems{
            models.ModuleStatItemFansItems{
                Airflow:              models.ToPointer("airflow8"),
                Name:                 models.ToPointer("name4"),
                Rpm:                  models.ToPointer(78),
                Status:               models.ToPointer("status4"),
            },
        },
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  models.ToPointer("fc3342123456"),
        Model:                models.NewOptional(models.ToPointer("EX4300-48P")),
        Serial:               models.NewOptional(models.ToPointer("PX8716230021")),
        VcRole:               models.NewOptional(models.ToPointer("master")),
    }

}
```

