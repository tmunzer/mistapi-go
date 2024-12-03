
# Remote Syslog User

*This model accepts additional fields of type interface{}.*

## Structure

`RemoteSyslogUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | - |
| `Match` | `*string` | Optional | - |
| `User` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "match": "\"!alarm|ntp|errors.crc_error[chan]\"",
  "user": "*",
  "contents": [
    {
      "facility": "ntp",
      "severity": "error",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

