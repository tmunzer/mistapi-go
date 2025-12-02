
# Snmpv 3 Config Notify Items

## Structure

`Snmpv3ConfigNotifyItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Tag` | `*string` | Optional | - |
| `Type` | [`*models.Snmpv3ConfigNotifyTypeEnum`](../../doc/models/snmpv-3-config-notify-type-enum.md) | Optional | enum: `inform`, `trap` |

## Example (as JSON)

```json
{
  "name": "name6",
  "tag": "tag0",
  "type": "inform"
}
```

