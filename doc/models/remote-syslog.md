
# Remote Syslog

*This model accepts additional fields of type interface{}.*

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
| `TimeFormat` | [`*models.RemoteSyslogTimeFormatEnum`](../../doc/models/remote-syslog-time-format-enum.md) | Optional | enum: `millisecond`, `year`, `year millisecond` |
| `Users` | [`[]models.RemoteSyslogUser`](../../doc/models/remote-syslog-user.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
    "size": "size8",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "console": {
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
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "files": [
    {
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
      "file": "file4",
      "match": "match4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
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
      "file": "file4",
      "match": "match4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
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
      "file": "file4",
      "match": "match4",
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

