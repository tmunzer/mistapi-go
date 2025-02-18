
# Snmp Vacm Access Item

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpVacmAccessItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GroupName` | `*string` | Optional | - |
| `PrefixList` | [`[]models.SnmpVacmAccessItemPrefixListItem`](../../doc/models/snmp-vacm-access-item-prefix-list-item.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "group_name": "group_name6",
  "prefix_list": [
    {
      "context_prefix": "context_prefix2",
      "notify_view": "notify_view6",
      "read_view": "read_view0",
      "security_level": "none",
      "security_model": "v1",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "context_prefix": "context_prefix2",
      "notify_view": "notify_view6",
      "read_view": "read_view0",
      "security_level": "none",
      "security_model": "v1",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

