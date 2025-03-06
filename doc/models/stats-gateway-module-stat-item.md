
# Stats Gateway Module Stat Item

*This model accepts additional fields of type interface{}.*

## Structure

`StatsGatewayModuleStatItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BackupVersion` | `models.Optional[string]` | Optional | - |
| `BiosVersion` | `models.Optional[string]` | Optional | - |
| `CpldVersion` | `models.Optional[string]` | Optional | - |
| `Fans` | [`[]models.ModuleStatItemFansItems`](../../doc/models/module-stat-item-fans-items.md) | Optional | **Constraints**: *Unique Items Required* |
| `FpgaVersion` | `models.Optional[string]` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | - |
| `Locating` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Model` | `models.Optional[string]` | Optional | - |
| `OpticsCpldVersion` | `models.Optional[string]` | Optional | - |
| `PendingVersion` | `models.Optional[string]` | Optional | - |
| `Poe` | [`*models.ModuleStatItemPoe`](../../doc/models/module-stat-item-poe.md) | Optional | - |
| `PoeVersion` | `models.Optional[string]` | Optional | - |
| `PowerCpldVersion` | `models.Optional[string]` | Optional | - |
| `Psus` | [`[]models.ModuleStatItemPsusItem`](../../doc/models/module-stat-item-psus-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `ReFpgaVersion` | `models.Optional[string]` | Optional | - |
| `RecoveryVersion` | `models.Optional[string]` | Optional | - |
| `Serial` | `models.Optional[string]` | Optional | - |
| `Status` | `models.Optional[string]` | Optional | - |
| `Temperatures` | [`[]models.ModuleStatItemTemperaturesItem`](../../doc/models/module-stat-item-temperatures-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `TmcFpgaVersion` | `models.Optional[string]` | Optional | - |
| `UbootVersion` | `models.Optional[string]` | Optional | - |
| `Uptime` | `models.Optional[int]` | Optional | - |
| `VcLinks` | [`[]models.ModuleStatItemVcLinksItem`](../../doc/models/module-stat-item-vc-links-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `VcMode` | `models.Optional[string]` | Optional | - |
| `VcRole` | `models.Optional[string]` | Optional | enum: `master`, `backup`, `linecard` |
| `VcState` | `models.Optional[string]` | Optional | - |
| `Version` | `models.Optional[string]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "fc3342123456",
  "model": "EX4300-48P",
  "serial": "PX8716230021",
  "vc_role": "master",
  "backup_version": "backup_version6",
  "bios_version": "bios_version8",
  "cpld_version": "cpld_version2",
  "fans": [
    {
      "airflow": "airflow8",
      "name": "name4",
      "status": "status4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "airflow": "airflow8",
      "name": "name4",
      "status": "status4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "airflow": "airflow8",
      "name": "name4",
      "status": "status4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "fpga_version": "fpga_version4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

