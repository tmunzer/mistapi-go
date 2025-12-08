
# Snmpv 3 Config Notify Filter Item

## Structure

`Snmpv3ConfigNotifyFilterItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.Snmpv3ConfigNotifyFilterItemContent`](../../doc/models/snmpv-3-config-notify-filter-item-content.md) | Optional | - |
| `ProfileName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "contents": [
    {
      "include": false,
      "oid": "oid4"
    },
    {
      "include": false,
      "oid": "oid4"
    },
    {
      "include": false,
      "oid": "oid4"
    }
  ],
  "profile_name": "profile_name4"
}
```

