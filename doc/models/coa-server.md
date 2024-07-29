
# Coa Server

CoA Server

## Structure

`CoaServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableEventTimestampCheck` | `*bool` | Optional | whether to disable Event-Timestamp Check<br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Ip` | `string` | Required | - |
| `Port` | `*int` | Optional | **Default**: `3799`<br>**Constraints**: `>= 1`, `<= 65535` |
| `Secret` | `string` | Required | - |

## Example (as JSON)

```json
{
  "disable_event_timestamp_check": false,
  "enabled": false,
  "ip": "1.2.3.4",
  "port": 3799,
  "secret": "testing456"
}
```

