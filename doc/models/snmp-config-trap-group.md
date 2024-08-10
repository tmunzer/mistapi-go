
# Snmp Config Trap Group

## Structure

`SnmpConfigTrapGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Categories` | `[]string` | Optional | - |
| `GroupName` | `*string` | Optional | Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp_trap-groups-configuring-junos-nm.html |
| `Targets` | `[]string` | Optional | - |
| `Version` | [`*models.SnmpConfigTrapVerionEnum`](../../doc/models/snmp-config-trap-verion-enum.md) | Optional | enum: `all`, `v1`, `v2`<br>**Default**: `"v2"` |

## Example (as JSON)

```json
{
  "group_name": "profiler",
  "version": "v2",
  "categories": [
    "categories8",
    "categories7"
  ],
  "targets": [
    "targets0",
    "targets1",
    "targets2"
  ]
}
```

