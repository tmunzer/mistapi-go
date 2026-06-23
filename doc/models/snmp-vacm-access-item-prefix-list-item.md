
# Snmp Vacm Access Item Prefix List Item

Context prefix rule for a VACM access entry

## Structure

`SnmpVacmAccessItemPrefixListItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ContextPrefix` | `*string` | Optional | Context prefix for this VACM access rule. Required only if `type`==`context_prefix` |
| `NotifyView` | `*string` | Optional | Notify view name referenced by this VACM access rule |
| `ReadView` | `*string` | Optional | Read view name referenced by this VACM access rule |
| `SecurityLevel` | [`*models.SnmpVacmAccessItemPrefixListItemLevelEnum`](../../doc/models/snmp-vacm-access-item-prefix-list-item-level-enum.md) | Optional | enum: `authentication`, `none`, `privacy` |
| `SecurityModel` | [`*models.SnmpVacmAccessItemPrefixListItemModelEnum`](../../doc/models/snmp-vacm-access-item-prefix-list-item-model-enum.md) | Optional | enum: `any`, `usm`, `v1`, `v2c` |
| `Type` | [`*models.SnmpVacmAccessItemTypeEnum`](../../doc/models/snmp-vacm-access-item-type-enum.md) | Optional | VACM context matching type for this access rule. enum: `context_prefix`, `default_context_prefix` |
| `WriteView` | `*string` | Optional | Write view name referenced by this VACM access rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpVacmAccessItemPrefixListItem := models.SnmpVacmAccessItemPrefixListItem{
        ContextPrefix:        models.ToPointer("iil"),
        NotifyView:           models.ToPointer("all"),
        ReadView:             models.ToPointer("all"),
        SecurityLevel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemLevelEnum_AUTHENTICATION),
        SecurityModel:        models.ToPointer(models.SnmpVacmAccessItemPrefixListItemModelEnum_ANY),
        WriteView:            models.ToPointer("all"),
    }

}
```

