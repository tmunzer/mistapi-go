
# Snmp Config View

SNMP MIB view definition

## Structure

`SnmpConfigView`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Include` | `*bool` | Optional | Whether the root OID is included in this SNMP view |
| `Oid` | `*string` | Optional | Root OID for this SNMP view |
| `ViewName` | `*string` | Optional | Name of the SNMP MIB view definition |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpConfigView := models.SnmpConfigView{
        Include:              models.ToPointer(false),
        Oid:                  models.ToPointer("1.3.6.1"),
        ViewName:             models.ToPointer("all"),
    }

}
```

