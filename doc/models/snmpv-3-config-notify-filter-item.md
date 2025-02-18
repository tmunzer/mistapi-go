
# Snmpv 3 Config Notify Filter Item

*This model accepts additional fields of type interface{}.*

## Structure

`Snmpv3ConfigNotifyFilterItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.Snmpv3ConfigNotifyFilterItemContent`](../../doc/models/snmpv-3-config-notify-filter-item-content.md) | Optional | - |
| `ProfileName` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "contents": [
    {
      "include": false,
      "oid": "oid4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "include": false,
      "oid": "oid4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "include": false,
      "oid": "oid4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "profile_name": "profile_name4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

