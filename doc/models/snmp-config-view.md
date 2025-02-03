
# Snmp Config View

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpConfigView`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Include` | `*bool` | Optional | If the root oid configured is included |
| `Oid` | `*string` | Optional | - |
| `ViewName` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "oid": "1.3.6.1",
  "view_name": "all",
  "include": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

