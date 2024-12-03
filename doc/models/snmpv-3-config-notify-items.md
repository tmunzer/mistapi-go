
# Snmpv 3 Config Notify Items

*This model accepts additional fields of type interface{}.*

## Structure

`Snmpv3ConfigNotifyItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Tag` | `*string` | Optional | - |
| `Type` | [`*models.Snmpv3ConfigNotifyTypeEnum`](../../doc/models/snmpv-3-config-notify-type-enum.md) | Optional | enum: `inform`, `trap` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name6",
  "tag": "tag0",
  "type": "inform",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

