
# Remote Syslog File Config

## Structure

`RemoteSyslogFileConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Archive` | [`*models.RemoteSyslogArchive`](../../doc/models/remote-syslog-archive.md) | Optional | - |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | - |
| `ExplicitPriority` | `*bool` | Optional | - |
| `File` | `*string` | Optional | - |
| `Match` | `*string` | Optional | - |
| `StructuredData` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "file": "file-name",
  "match": "!alarm|ntp|errors.crc_error[chan]",
  "archive": {
    "files": 122,
    "size": "size8"
  },
  "contents": [
    {
      "facility": "ntp",
      "severity": "error"
    },
    {
      "facility": "ntp",
      "severity": "error"
    }
  ],
  "explicit_priority": false
}
```

