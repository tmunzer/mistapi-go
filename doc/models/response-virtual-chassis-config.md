
# Response Virtual Chassis Config

Combined virtual chassis status, topology, and member statistics

## Structure

`ResponseVirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigType` | `*string` | Optional, Read-only | Provisioning mode of the virtual chassis configuration |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Locating` | `*bool` | Optional, Read-only | Whether locate mode is active for the virtual chassis device |
| `Mac` | `*string` | Optional | Device MAC address for the virtual chassis device record |
| `Members` | [`[]models.StatsSwitchModuleStatItem`](../../doc/models/stats-switch-module-stat-item.md) | Optional | Hardware module statistics reported by a switch<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Model` | `*string` | Optional, Read-only | Switch model for the virtual chassis device record |
| `NumRoutingEngines` | `*int` | Optional | Number of routing engines in the virtual chassis |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Serial` | `*string` | Optional, Read-only | Device serial number reported for the virtual chassis |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Status` | `*string` | Optional, Read-only | Connection status of the virtual chassis device |
| `Type` | `*string` | Optional | Device type for the virtual chassis record |
| `VcMac` | `*string` | Optional, Read-only | Virtual Chassis MAC address used to identify the VC device |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseVirtualChassisConfig := models.ResponseVirtualChassisConfig{
        ConfigType:           models.ToPointer("config_type0"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Locating:             models.ToPointer(false),
        Mac:                  models.ToPointer("mac6"),
        Members:              []models.StatsSwitchModuleStatItem{
            models.StatsSwitchModuleStatItem{
                BackupVersion:        models.NewOptional(models.ToPointer("backup_version8")),
                BiosVersion:          models.NewOptional(models.ToPointer("bios_version6")),
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
            },
            models.StatsSwitchModuleStatItem{
                BackupVersion:        models.NewOptional(models.ToPointer("backup_version8")),
                BiosVersion:          models.NewOptional(models.ToPointer("bios_version6")),
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
            },
        },
        NumRoutingEngines:    models.ToPointer(1),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

