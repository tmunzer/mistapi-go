
# Mxedge Das Coa Server

## Structure

`MxedgeDasCoaServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableEventTimestampCheck` | `*bool` | Optional | whether to disable Event-Timestamp Check<br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | - |
| `Host` | `*string` | Optional | this server configured to send CoA\|DM to mist edges |
| `Port` | `*int` | Optional | mist edges will allow this host on this port<br>**Default**: `3799` |
| `Secret` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "disable_event_timestamp_check": false,
  "port": 3799,
  "enabled": false,
  "host": "host2",
  "secret": "secret8"
}
```

