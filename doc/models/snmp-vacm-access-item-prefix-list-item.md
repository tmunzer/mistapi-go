
# Snmp Vacm Access Item Prefix List Item

## Structure

`SnmpVacmAccessItemPrefixListItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ContextPrefix` | `*string` | Optional | only required if `type`==`context_prefix` |
| `NotifyView` | `*string` | Optional | refer to view name |
| `ReadView` | `*string` | Optional | refer to view name |
| `SecurityLevel` | [`*models.SnmpVacmAccessItemPrefixListItemLevelEnum`](../../doc/models/snmp-vacm-access-item-prefix-list-item-level-enum.md) | Optional | enum: `authentication`, `none`, `privacy` |
| `SecurityModel` | [`*models.SnmpVacmAccessItemPrefixListItemModelEnum`](../../doc/models/snmp-vacm-access-item-prefix-list-item-model-enum.md) | Optional | enum: `any`, `usm`, `v1`, `v2c` |
| `Type` | [`*models.SnmpVacmAccessItemTypeEnum`](../../doc/models/snmp-vacm-access-item-type-enum.md) | Optional | enum: `context_prefix`, `default_context_prefix` |
| `WriteView` | `*string` | Optional | refer to view name |

## Example (as JSON)

```json
{
  "context_prefix": "iil",
  "notify_view": "all",
  "read_view": "all",
  "write_view": "all",
  "security_level": "privacy",
  "security_model": "v1"
}
```

