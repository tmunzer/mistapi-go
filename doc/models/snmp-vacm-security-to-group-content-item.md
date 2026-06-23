
# Snmp Vacm Security to Group Content Item

VACM security-name to group mapping entry

## Structure

`SnmpVacmSecurityToGroupContentItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Group` | `*string` | Optional | VACM group name referenced by this mapping |
| `SecurityName` | `*string` | Optional | Name of the SNMP security principal mapped to a VACM group |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpVacmSecurityToGroupContentItem := models.SnmpVacmSecurityToGroupContentItem{
        Group:                models.ToPointer("group0"),
        SecurityName:         models.ToPointer("security_name2"),
    }

}
```

