
# Snmp Config Trap Group

## Structure

`SnmpConfigTrapGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Categories` | `[]string` | Optional | - |
| `GroupName` | `*string` | Optional | Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp_trap-groups-configuring-junos-nm.html |
| `Targets` | `[]string` | Optional | - |
| `Version` | [`*models.SnmpConfigTrapVerionEnum`](../../doc/models/snmp-config-trap-verion-enum.md) | Optional | **Default**: `"v2"` |

## Example (as JSON)

```json
{
  "group_name": "profiler",
  "version": "v2",
  "categories": [
    "categories2",
    "categories3"
  ],
  "targets": [
    "targets2",
    "targets3"
  ]
}
```

