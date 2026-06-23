
# Snmp Vacm

SNMPv3 View-based Access Control Model configuration

## Structure

`SnmpVacm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Access` | [`[]models.SnmpVacmAccessItem`](../../doc/models/snmp-vacm-access-item.md) | Optional | VACM access rules for SNMPv3 |
| `SecurityToGroup` | [`*models.SnmpVacmSecurityToGroup`](../../doc/models/snmp-vacm-security-to-group.md) | Optional | VACM security-name to group mapping configuration |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpVacm := models.SnmpVacm{
        Access:               []models.SnmpVacmAccessItem{
            models.SnmpVacmAccessItem{
                GroupName:            models.ToPointer("group_name4"),
                PrefixList:           []models.SnmpVacmAccessItemPrefixListItem{
                    models.SnmpVacmAccessItemPrefixListItem{
                        ContextPrefix:        models.ToPointer("context_prefix2"),
                        NotifyView:           models.ToPointer("notify_view6"),
                        ReadView:             models.ToPointer("read_view0"),
                        SecurityLevel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum_NONE),
                        SecurityModel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum_V1),
                    },
                    models.SnmpVacmAccessItemPrefixListItem{
                        ContextPrefix:        models.ToPointer("context_prefix2"),
                        NotifyView:           models.ToPointer("notify_view6"),
                        ReadView:             models.ToPointer("read_view0"),
                        SecurityLevel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum_NONE),
                        SecurityModel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum_V1),
                    },
                },
            },
            models.SnmpVacmAccessItem{
                GroupName:            models.ToPointer("group_name4"),
                PrefixList:           []models.SnmpVacmAccessItemPrefixListItem{
                    models.SnmpVacmAccessItemPrefixListItem{
                        ContextPrefix:        models.ToPointer("context_prefix2"),
                        NotifyView:           models.ToPointer("notify_view6"),
                        ReadView:             models.ToPointer("read_view0"),
                        SecurityLevel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum_NONE),
                        SecurityModel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum_V1),
                    },
                    models.SnmpVacmAccessItemPrefixListItem{
                        ContextPrefix:        models.ToPointer("context_prefix2"),
                        NotifyView:           models.ToPointer("notify_view6"),
                        ReadView:             models.ToPointer("read_view0"),
                        SecurityLevel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum_NONE),
                        SecurityModel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum_V1),
                    },
                },
            },
        },
        SecurityToGroup:      models.ToPointer(models.SnmpVacmSecurityToGroup{
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
        }),
    }

}
```

