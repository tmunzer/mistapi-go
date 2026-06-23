
# Snmp Vacm Access Item

VACM access rule for an SNMP group

## Structure

`SnmpVacmAccessItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GroupName` | `*string` | Optional | SNMP VACM group name |
| `PrefixList` | [`[]models.SnmpVacmAccessItemPrefixListItem`](../../doc/models/snmp-vacm-access-item-prefix-list-item.md) | Optional | Context prefix rules for a VACM access entry |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpVacmAccessItem := models.SnmpVacmAccessItem{
        GroupName:            models.ToPointer("group_name4"),
        PrefixList:           []models.SnmpVacmAccessItemPrefixListItem{
            models.SnmpVacmAccessItemPrefixListItem{
                ContextPrefix:        models.ToPointer("context_prefix2"),
                NotifyView:           models.ToPointer("notify_view6"),
                ReadView:             models.ToPointer("read_view0"),
                SecurityLevel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum_NONE),
                SecurityModel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum_V1),
            },
        },
    }

}
```

