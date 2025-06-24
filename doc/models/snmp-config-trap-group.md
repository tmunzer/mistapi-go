
# Snmp Config Trap Group

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpConfigTrapGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Categories` | `[]string` | Optional | - |
| `GroupName` | `*string` | Optional | Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp_trap-groups-configuring-junos-nm.html |
| `Targets` | `[]string` | Optional | - |
| `Version` | [`*models.SnmpConfigTrapVersionEnum`](../../doc/models/snmp-config-trap-version-enum.md) | Optional | enum: `all`, `v1`, `v2`<br><br>**Default**: `"v2"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

