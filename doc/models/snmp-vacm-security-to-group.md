
# Snmp Vacm Security to Group

VACM security-name to group mapping configuration

## Structure

`SnmpVacmSecurityToGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Content` | [`[]models.SnmpVacmSecurityToGroupContentItem`](../../doc/models/snmp-vacm-security-to-group-content-item.md) | Optional | VACM security-name to group mapping entries |
| `SecurityModel` | [`*models.SnmpVacmSecurityModelEnum`](../../doc/models/snmp-vacm-security-model-enum.md) | Optional | enum: `usm`, `v1`, `v2c` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpVacmSecurityToGroup := models.SnmpVacmSecurityToGroup{
        Content:              []models.SnmpVacmSecurityToGroupContentItem{
            models.SnmpVacmSecurityToGroupContentItem{
                Group:                models.ToPointer("group2"),
                SecurityName:         models.ToPointer("security_name6"),
            },
            models.SnmpVacmSecurityToGroupContentItem{
                Group:                models.ToPointer("group2"),
                SecurityName:         models.ToPointer("security_name6"),
            },
            models.SnmpVacmSecurityToGroupContentItem{
                Group:                models.ToPointer("group2"),
                SecurityName:         models.ToPointer("security_name6"),
            },
        },
        SecurityModel:        models.ToPointer(models.SnmpVacmSecurityModelEnum_V2C),
    }

}
```

