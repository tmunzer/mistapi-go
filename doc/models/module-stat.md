
# Module Stat

## Structure

`ModuleStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BackupVersion` | `*string` | Optional | - |
| `BiosVersion` | `*string` | Optional | - |
| `Errors` | [`[]models.ModuleStatErrorsItems`](../../doc/models/module-stat-errors-items.md) | Optional | used to report all error states the device node is running into.<br>An error should always have `type` and `since` fields, and could have some other fields specific to that type. |
| `Fans` | [`[]models.ModuleStatFansItems`](../../doc/models/module-stat-fans-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Model` | `*string` | Optional | - |
| `Poe` | [`*models.ModuleStatPoe`](../../doc/models/module-stat-poe.md) | Optional | - |
| `Psus` | [`[]models.ModuleStatPsusItems`](../../doc/models/module-stat-psus-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Serial` | `*string` | Optional | - |
| `Temperatures` | [`[]models.ModuleStatTemperaturesItems`](../../doc/models/module-stat-temperatures-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `VcLinks` | [`[]models.ModuleStatVcLinksItems`](../../doc/models/module-stat-vc-links-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `VcMode` | `*string` | Optional | - |
| `VcRole` | `*string` | Optional | master / backup / linecard |
| `VcState` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "model": "EX4300-48P",
  "serial": "PX8716230021",
  "vc_role": "master",
  "backup_version": "backup_version8",
  "bios_version": "bios_version6",
  "errors": [
    {
      "feature": "feature4",
      "minimum_version": "minimum_version2",
      "reason": "reason4",
      "since": 174,
      "type": "type0"
    }
  ],
  "fans": [
    {
      "airflow": "airflow8",
      "name": "name4",
      "status": "status4"
    }
  ]
}
```

