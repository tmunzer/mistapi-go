
# Remote Syslog File Config

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "file": "file-name",
  "match": "!alarm|ntp|errors.crc_error[chan]",
  "archive": {
    "files": 122,
    "size": "size8",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "contents": [
    {
      "facility": "ntp",
      "severity": "error",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "facility": "ntp",
      "severity": "error",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "explicit_priority": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

