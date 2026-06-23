
# Snmp Config Trap Group

SNMP trap group definition

## Structure

`SnmpConfigTrapGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Categories` | `[]string` | Optional | Trap categories included in an SNMP trap group |
| `GroupName` | `*string` | Optional | Trap group name for this SNMP trap group |
| `Targets` | `[]string` | Optional | Trap target addresses for an SNMP trap group |
| `Version` | [`*models.SnmpConfigTrapVersionEnum`](../../doc/models/snmp-config-trap-version-enum.md) | Optional | enum: `all`, `v1`, `v2`<br><br>**Default**: `"v2"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpConfigTrapGroup := models.SnmpConfigTrapGroup{
        Categories:           []string{
            "categories6",
            "categories5",
            "categories4",
        },
        GroupName:            models.ToPointer("profiler"),
        Targets:              []string{
            "targets8",
            "targets9",
        },
        Version:              models.ToPointer(models.SnmpConfigTrapVersionEnum_V2),
    }

}
```

