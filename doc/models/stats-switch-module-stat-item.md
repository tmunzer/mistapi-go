
# Stats Switch Module Stat Item

Hardware module status and firmware inventory for a switch

## Structure

`StatsSwitchModuleStatItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BackupVersion` | `models.Optional[string]` | Optional, Read-only | Backup software version stored on the switch module |
| `BiosVersion` | `models.Optional[string]` | Optional, Read-only | BIOS firmware version reported for the switch module |
| `BootPartition` | `*string` | Optional | Active boot partition used by the switch module |
| `CpldVersion` | `models.Optional[string]` | Optional, Read-only | CPLD firmware version reported for the switch module |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | CPU utilization breakdown for a device |
| `Errors` | [`[]models.ModuleStatItemErrorsItems`](../../doc/models/module-stat-item-errors-items.md) | Optional | Used to report all error states the device node is running into. An error should always have `type` and `since` fields, and could have some other fields specific to that type. |
| `Fans` | [`[]models.ModuleStatItemFansItems`](../../doc/models/module-stat-item-fans-items.md) | Optional | Cooling fan status records for a device module<br><br>**Constraints**: *Unique Items Required* |
| `FpcIdx` | `*int` | Optional, Read-only | FPC index identifying this switch module |
| `FpgaVersion` | `models.Optional[string]` | Optional, Read-only | FPGA firmware version reported for the switch module |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Locating` | `*bool` | Optional | Whether the switch module locator indicator is active |
| `Mac` | `*string` | Optional | Switch module MAC address reported by Mist |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory utilization statistics for a device; in a virtual chassis, this reports the master Routing Engine |
| `Model` | `models.Optional[string]` | Optional, Read-only | Switch module model name reported by Mist |
| `OpticsCpldVersion` | `models.Optional[string]` | Optional, Read-only | Optics CPLD firmware version reported for the switch module |
| `PendingVersion` | `models.Optional[string]` | Optional, Read-only | Pending software version staged for the switch module |
| `Pics` | [`[]models.ModuleStatItemPicsItem`](../../doc/models/module-stat-item-pics-item.md) | Optional | Physical Interface Card summaries for a device module |
| `Poe` | [`*models.ModuleStatItemPoe`](../../doc/models/module-stat-item-poe.md) | Optional | Power over Ethernet telemetry for a device module |
| `PoeVersion` | `models.Optional[string]` | Optional, Read-only | PoE controller firmware version reported for the switch module |
| `PowerCpldVersion` | `models.Optional[string]` | Optional, Read-only | Power CPLD firmware version reported for the switch module |
| `Psus` | [`[]models.ModuleStatItemPsusItem`](../../doc/models/module-stat-item-psus-item.md) | Optional | Power supply status records for a device module<br><br>**Constraints**: *Unique Items Required* |
| `ReFpgaVersion` | `models.Optional[string]` | Optional, Read-only | Routing Engine FPGA firmware version reported for the switch module |
| `RecoveryVersion` | `models.Optional[string]` | Optional, Read-only | Recovery software version stored on the switch module |
| `Serial` | `models.Optional[string]` | Optional, Read-only | Switch module serial number reported by Mist |
| `Status` | `models.Optional[string]` | Optional, Read-only | Connection status reported for the switch module |
| `Temperatures` | [`[]models.ModuleStatItemTemperaturesItem`](../../doc/models/module-stat-item-temperatures-item.md) | Optional | Temperature sensor readings for a device module<br><br>**Constraints**: *Unique Items Required* |
| `TmcFpgaVersion` | `models.Optional[string]` | Optional, Read-only | TMC FPGA firmware version reported for the switch module |
| `Type` | `models.Optional[string]` | Optional, Read-only | Module type reported for the switch module |
| `UbootVersion` | `models.Optional[string]` | Optional, Read-only | U-Boot firmware version reported for the switch module |
| `Uptime` | `models.Optional[int]` | Optional, Read-only | Elapsed time since the switch module last booted, in seconds |
| `VcLinks` | [`[]models.ModuleStatItemVcLinksItem`](../../doc/models/module-stat-item-vc-links-item.md) | Optional | Virtual chassis link records for a device module<br><br>**Constraints**: *Unique Items Required* |
| `VcMode` | `models.Optional[string]` | Optional, Read-only | Virtual chassis mode reported for the switch module |
| `VcRole` | `models.Optional[string]` | Optional, Read-only | enum: `master`, `backup`, `linecard` |
| `VcState` | `models.Optional[string]` | Optional, Read-only | Virtual chassis state reported for the switch module |
| `Version` | `models.Optional[string]` | Optional, Read-only | Software version running on the switch module |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSwitchModuleStatItem := models.StatsSwitchModuleStatItem{
        BackupVersion:        models.NewOptional(models.ToPointer("backup_version8")),
        BiosVersion:          models.NewOptional(models.ToPointer("bios_version4")),
        BootPartition:        models.ToPointer("boot_partition8"),
        CpldVersion:          models.NewOptional(models.ToPointer("cpld_version4")),
        CpuStat:              models.ToPointer(models.CpuStat{
            Idle:                 models.NewOptional(models.ToPointer(float64(102.08))),
            Interrupt:            models.NewOptional(models.ToPointer(float64(215.84))),
            LoadAvg:              []float64{
                float64(105.91),
            },
            System:               models.NewOptional(models.ToPointer(float64(13.6))),
            Usage:                models.NewOptional(models.ToPointer(float64(125.9))),
        }),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  models.ToPointer("fc3342123456"),
        Model:                models.NewOptional(models.ToPointer("EX4300-48P")),
        Serial:               models.NewOptional(models.ToPointer("PX8716230021")),
        VcRole:               models.NewOptional(models.ToPointer("master")),
    }

}
```

