
# Mxedge Das Coa Server

## Structure

`MxedgeDasCoaServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableEventTimestampCheck` | `*bool` | Optional | Whether to disable Event-Timestamp Check<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | - |
| `Host` | `*string` | Optional | This server configured to send CoA\|DM to mist edges |
| `Port` | `*int` | Optional | Mist edges will allow this host on this port<br><br>**Default**: `3799` |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br><br>**Default**: `false` |
| `Secret` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "disable_event_timestamp_check": false,
  "port": 3799,
  "require_message_authenticator": false,
  "enabled": false,
  "host": "host8"
}
```

