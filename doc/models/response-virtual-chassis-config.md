
# Response Virtual Chassis Config

## Structure

`ResponseVirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigType` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Locating` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Members` | [`[]models.StatsSwitchModuleStatItem`](../../doc/models/stats-switch-module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Model` | `*string` | Optional | - |
| `NumRoutingEngines` | `*int` | Optional | routing-engine count |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `VcMac` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "num_routing_engines": 1,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "config_type": "config_type6",
  "locating": false,
  "mac": "mac2",
  "members": [
    {
      "backup_version": "backup_version8",
      "bios_version": "bios_version6",
      "boot_partition": "boot_partition8",
      "cpld_version": "cpld_version4",
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
  ]
}
```

