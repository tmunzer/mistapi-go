
# Remote Syslog

## Structure

`RemoteSyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Archive` | [`*models.RemoteSyslogArchive`](../../doc/models/remote-syslog-archive.md) | Optional | - |
| `Console` | [`*models.RemoteSyslogConsole`](../../doc/models/remote-syslog-console.md) | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Files` | [`[]models.RemoteSyslogFileConfig`](../../doc/models/remote-syslog-file-config.md) | Optional | - |
| `Network` | `*string` | Optional | if source_address is configured, will use the vlan firstly otherwise use source_ip |
| `SendToAllServers` | `*bool` | Optional | **Default**: `false` |
| `Servers` | [`[]models.RemoteSyslogServer`](../../doc/models/remote-syslog-server.md) | Optional | - |
| `TimeFormat` | [`*models.RemoteSyslogTimeFormatEnum`](../../doc/models/remote-syslog-time-format-enum.md) | Optional | - |
| `Users` | [`[]models.RemoteSyslogUser`](../../doc/models/remote-syslog-user.md) | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "network": "default",
  "send_to_all_servers": false,
  "servers": [
    {
      "facility": "config",
      "host": "syslogd.internal",
      "port": 514,
      "protocol": "udp",
      "severity": "info",
      "tag": ""
    }
  ],
  "time_format": "millisecond",
  "archive": {
    "files": 122,
    "size": "size8"
  },
  "console": {
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
  },
  "files": [
    {
      "archive": {
        "files": 122,
        "size": "size8"
      },
      "contents": [
        {
          "facility": "pfe",
          "severity": "warning"
        },
        {
          "facility": "pfe",
          "severity": "warning"
        }
      ],
      "explicit_priority": false,
      "file": "file4",
      "match": "match4"
    }
  ]
}
```

