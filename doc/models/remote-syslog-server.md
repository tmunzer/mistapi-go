
# Remote Syslog Server

*This model accepts additional fields of type interface{}.*

## Structure

`RemoteSyslogServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | - |
| `ExplicitPriority` | `*bool` | Optional | - |
| `Facility` | [`*models.RemoteSyslogFacilityEnum`](../../doc/models/remote-syslog-facility-enum.md) | Optional | enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`<br>**Default**: `"any"` |
| `Host` | `*string` | Optional | - |
| `Match` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Default**: `514` |
| `Protocol` | [`*models.RemoteSyslogServerProtocolEnum`](../../doc/models/remote-syslog-server-protocol-enum.md) | Optional | enum: `tcp`, `udp`<br>**Default**: `"udp"` |
| `RoutingInstance` | `*string` | Optional | - |
| `Severity` | [`*models.RemoteSyslogSeverityEnum`](../../doc/models/remote-syslog-severity-enum.md) | Optional | enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`<br>**Default**: `"any"` |
| `SourceAddress` | `*string` | Optional | if source_address is configured, will use the vlan firstly otherwise use source_ip |
| `StructuredData` | `*bool` | Optional | - |
| `Tag` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "facility": "config",
  "host": "syslogd.internal",
  "match": "!alarm|ntp|errors.crc_error[chan]",
  "port": 514,
  "protocol": "udp",
  "routing_instance": "routing-instance-name",
  "severity": "any",
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
  "explicit_priority": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

