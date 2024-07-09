
# Remote Syslog Server

## Structure

`RemoteSyslogServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | - |
| `ExplicitPriority` | `*bool` | Optional | - |
| `Facility` | [`*models.RemoteSyslogFacilityEnum`](../../doc/models/remote-syslog-facility-enum.md) | Optional | **Default**: `"any"` |
| `Host` | `*string` | Optional | - |
| `Match` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Default**: `514` |
| `Protocol` | [`*models.RemoteSyslogServerProtocolEnum`](../../doc/models/remote-syslog-server-protocol-enum.md) | Optional | **Default**: `"udp"` |
| `RoutingInstance` | `*string` | Optional | - |
| `Severity` | [`*models.RemoteSyslogSeverityEnum`](../../doc/models/remote-syslog-severity-enum.md) | Optional | **Default**: `"any"` |
| `SourceAddress` | `*string` | Optional | if source_address is configured, will use the vlan firstly otherwise use source_ip |
| `StructuredData` | `*bool` | Optional | - |
| `Tag` | `*string` | Optional | - |

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
      "facility": "pfe",
      "severity": "warning"
    },
    {
      "facility": "pfe",
      "severity": "warning"
    }
  ],
  "explicit_priority": false
}
```

