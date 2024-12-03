
# Mxedge Das Coa Server

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeDasCoaServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableEventTimestampCheck` | `*bool` | Optional | whether to disable Event-Timestamp Check<br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | - |
| `Host` | `*string` | Optional | this server configured to send CoA\|DM to mist edges |
| `Port` | `*int` | Optional | mist edges will allow this host on this port<br>**Default**: `3799` |
| `RequireMessageAuthenticator` | `*bool` | Optional | whether to require Message-Authenticator in requests<br>**Default**: `false` |
| `Secret` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disable_event_timestamp_check": false,
  "port": 3799,
  "require_message_authenticator": false,
  "enabled": false,
  "host": "host8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

