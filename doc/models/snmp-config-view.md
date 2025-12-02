
# Snmp Config View

## Structure

`SnmpConfigView`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Include` | `*bool` | Optional | If the root oid configured is included |
| `Oid` | `*string` | Optional | - |
| `ViewName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "oid": "1.3.6.1",
  "view_name": "all",
  "include": false
}
```

