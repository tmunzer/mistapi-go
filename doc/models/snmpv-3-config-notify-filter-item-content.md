
# Snmpv 3 Config Notify Filter Item Content

*This model accepts additional fields of type interface{}.*

## Structure

`Snmpv3ConfigNotifyFilterItemContent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Include` | `*bool` | Optional | - |
| `Oid` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "oid": "1.3.6.1.4.1",
  "include": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

