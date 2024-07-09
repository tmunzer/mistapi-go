
# Remote Syslog Console

## Structure

`RemoteSyslogConsole`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | - |

## Example (as JSON)

```json
{
  "contents": [
    {
      "facility": "pfe",
      "severity": "warning"
    },
    {
      "facility": "pfe",
      "severity": "warning"
    },
    {
      "facility": "pfe",
      "severity": "warning"
    }
  ]
}
```

