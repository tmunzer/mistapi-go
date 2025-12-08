
# Stats Switch Module Stat Item

## Structure

`StatsSwitchModuleStatItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BackupVersion` | `models.Optional[string]` | Optional | - |
| `BiosVersion` | `models.Optional[string]` | Optional | - |
| `BootPartition` | `*string` | Optional | - |
| `CpldVersion` | `models.Optional[string]` | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `Errors` | [`[]models.ModuleStatItemErrorsItems`](../../doc/models/module-stat-item-errors-items.md) | Optional | Used to report all error states the device node is running into. An error should always have `type` and `since` fields, and could have some other fields specific to that type. |
| `Fans` | [`[]models.ModuleStatItemFansItems`](../../doc/models/module-stat-item-fans-items.md) | Optional | **Constraints**: *Unique Items Required* |
| `FpcIdx` | `*int` | Optional | - |
| `FpgaVersion` | `models.Optional[string]` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Locating` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | Memory usage stat (for virtual chassis, memory usage of master RE) |
| `Model` | `models.Optional[string]` | Optional | - |
| `OpticsCpldVersion` | `models.Optional[string]` | Optional | - |
| `PendingVersion` | `models.Optional[string]` | Optional | - |
| `Pics` | [`[]models.ModuleStatItemPicsItem`](../../doc/models/module-stat-item-pics-item.md) | Optional | - |
| `Poe` | [`*models.ModuleStatItemPoe`](../../doc/models/module-stat-item-poe.md) | Optional | - |
| `PoeVersion` | `models.Optional[string]` | Optional | - |
| `PowerCpldVersion` | `models.Optional[string]` | Optional | - |
| `Psus` | [`[]models.ModuleStatItemPsusItem`](../../doc/models/module-stat-item-psus-item.md) | Optional | **Constraints**: *Unique Items Required* |
| `ReFpgaVersion` | `models.Optional[string]` | Optional | - |
| `RecoveryVersion` | `models.Optional[string]` | Optional | - |
| `Serial` | `models.Optional[string]` | Optional | - |
| `Status` | `models.Optional[string]` | Optional | - |
| `Temperatures` | [`[]models.ModuleStatItemTemperaturesItem`](../../doc/models/module-stat-item-temperatures-item.md) | Optional | **Constraints**: *Unique Items Required* |
| `TmcFpgaVersion` | `models.Optional[string]` | Optional | - |
| `Type` | `models.Optional[string]` | Optional | - |
| `UbootVersion` | `models.Optional[string]` | Optional | - |
| `Uptime` | `models.Optional[int]` | Optional | - |
| `VcLinks` | [`[]models.ModuleStatItemVcLinksItem`](../../doc/models/module-stat-item-vc-links-item.md) | Optional | **Constraints**: *Unique Items Required* |
| `VcMode` | `models.Optional[string]` | Optional | - |
| `VcRole` | `models.Optional[string]` | Optional | enum: `master`, `backup`, `linecard` |
| `VcState` | `models.Optional[string]` | Optional | - |
| `Version` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "last_seen": 1470417522,
  "mac": "fc3342123456",
  "model": "EX4300-48P",
  "serial": "PX8716230021",
  "vc_role": "master",
  "backup_version": "backup_version0",
  "bios_version": "bios_version4",
  "boot_partition": "boot_partition0",
  "cpld_version": "cpld_version6",
  "cpu_stat": {
    "idle": 102.08,
    "interrupt": 215.84,
    "load_avg": [
      105.91
    ],
    "system": 13.6,
    "usage": 125.9
  }
}
```

