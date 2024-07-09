
# Snmp Vacm Access Item

## Structure

`SnmpVacmAccessItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GroupName` | `*string` | Optional | - |
| `PrefixList` | [`[]models.SnmpVacmAccessItemPrefixListItem`](../../doc/models/snmp-vacm-access-item-prefix-list-item.md) | Optional | - |

## Example (as JSON)

```json
{
  "group_name": "group_name2",
  "prefix_list": [
    {
      "context_prefix": "context_prefix2",
      "notify_view": "notify_view6",
      "read_view": "read_view0",
      "security_level": "authentication",
      "security_model": "v1"
    },
    {
      "context_prefix": "context_prefix2",
      "notify_view": "notify_view6",
      "read_view": "read_view0",
      "security_level": "authentication",
      "security_model": "v1"
    },
    {
      "context_prefix": "context_prefix2",
      "notify_view": "notify_view6",
      "read_view": "read_view0",
      "security_level": "authentication",
      "security_model": "v1"
    }
  ]
}
```

