
# Module Stat Item

*This model accepts additional fields of type interface{}.*

## Structure

`ModuleStatItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BackupVersion` | `models.Optional[string]` | Optional | - |
| `BiosVersion` | `models.Optional[string]` | Optional | - |
| `CpldVersion` | `models.Optional[string]` | Optional | - |
| `Errors` | [`[]models.ModuleStatItemErrorsItems`](../../doc/models/module-stat-item-errors-items.md) | Optional | used to report all error states the device node is running into. An error should always have `type` and `since` fields, and could have some other fields specific to that type. |
| `Fans` | [`[]models.ModuleStatItemFansItems`](../../doc/models/module-stat-item-fans-items.md) | Optional | **Constraints**: *Unique Items Required* |
| `FpcIdx` | `*int` | Optional | - |
| `FpgaVersion` | `models.Optional[string]` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | - |
| `Model` | `models.Optional[string]` | Optional | - |
| `OpticsCpldVersion` | `models.Optional[string]` | Optional | - |
| `PendingVersion` | `models.Optional[string]` | Optional | - |
| `Pics` | [`[]models.ModuleStatItemPicsItem`](../../doc/models/module-stat-item-pics-item.md) | Optional | - |
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
| `VcRole` | `models.Optional[string]` | Optional | master / backup / linecard |
| `VcState` | `models.Optional[string]` | Optional | - |
| `Version` | `models.Optional[string]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "model": "EX4300-48P",
  "serial": "PX8716230021",
  "vc_role": "master",
  "backup_version": "backup_version8",
  "bios_version": "bios_version6",
  "cpld_version": "cpld_version4",
  "errors": [
    {
      "feature": "feature4",
      "minimum_version": "minimum_version2",
      "reason": "reason4",
      "since": 174,
      "type": "type0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "feature": "feature4",
      "minimum_version": "minimum_version2",
      "reason": "reason4",
      "since": 174,
      "type": "type0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "feature": "feature4",
      "minimum_version": "minimum_version2",
      "reason": "reason4",
      "since": 174,
      "type": "type0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

