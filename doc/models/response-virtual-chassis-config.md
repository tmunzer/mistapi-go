
# Response Virtual Chassis Config

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseVirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigType` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Locating` | `*bool` | Optional | - |
| `Members` | [`[]models.StatsSwitchModuleStatItem`](../../doc/models/stats-switch-module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Model` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | `*string` | Optional | - |
| `VcMac` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "config_type": "config_type6",
  "locating": false,
  "members": [
    {
      "backup_version": "backup_version8",
      "bios_version": "bios_version6",
      "cpld_version": "cpld_version4",
      "cpu_stat": {
        "idle": 102.08,
        "interrupt": 215.84,
        "load_avg": [
          105.91
        ],
        "system": 13.6,
        "user": 204.52,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
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
        }
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "model": "model6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

