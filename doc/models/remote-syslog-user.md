
# Remote Syslog User

## Structure

`RemoteSyslogUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | - |
| `Match` | `*string` | Optional | - |
| `User` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "match": "\"!alarm|ntp|errors.crc_error[chan]\"",
  "user": "*",
  "contents": [
    {
      "facility": "ntp",
      "severity": "error"
    }
  ]
}
```

