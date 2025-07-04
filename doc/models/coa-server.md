
# Coa Server

CoA Server

*This model accepts additional fields of type interface{}.*

## Structure

`CoaServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableEventTimestampCheck` | `*bool` | Optional | Whether to disable Event-Timestamp Check<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Ip` | `string` | Required | - |
| `Port` | [`*models.RadiusCoaPort`](../../doc/models/containers/radius-coa-port.md) | Optional | Radius CoA Port, value from 1 to 65535, default is 3799 |
| `Secret` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disable_event_timestamp_check": false,
  "enabled": false,
  "ip": "1.2.3.4",
  "secret": "testing456",
  "port": 216,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

