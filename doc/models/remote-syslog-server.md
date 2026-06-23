
# Remote Syslog Server

Remote syslog server destination settings

## Structure

`RemoteSyslogServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | List of syslog content selectors |
| `ExplicitPriority` | `*bool` | Optional | Whether to include explicit syslog priority values in messages sent to this server |
| `Facility` | [`*models.RemoteSyslogFacilityEnum`](../../doc/models/remote-syslog-facility-enum.md) | Optional | enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`<br><br>**Default**: `"any"` |
| `Host` | `*string` | Optional | Address or hostname of the remote syslog server |
| `Match` | `*string` | Optional | Expression used to filter log messages sent to this server |
| `Port` | [`*models.RemoteSyslogServerPort`](../../doc/models/containers/remote-syslog-server-port.md) | Optional | Syslog Service Port, value from 1 to 65535 |
| `Protocol` | [`*models.RemoteSyslogServerProtocolEnum`](../../doc/models/remote-syslog-server-protocol-enum.md) | Optional | Transport protocol used for this remote syslog server. enum: `tcp`, `udp`<br><br>**Default**: `"udp"` |
| `RoutingInstance` | `*string` | Optional | Routing instance used to reach this remote syslog server |
| `ServerName` | `*string` | Optional | TLS server name used when verifying the remote syslog server certificate |
| `Severity` | [`*models.RemoteSyslogSeverityEnum`](../../doc/models/remote-syslog-severity-enum.md) | Optional | enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`<br><br>**Default**: `"any"` |
| `SourceAddress` | `*string` | Optional | Source address for syslog traffic. If configured, Mist uses the VLAN first; otherwise it uses `source_ip` |
| `StructuredData` | `*bool` | Optional | Whether to include structured syslog data in messages sent to this server |
| `Tag` | `*string` | Optional | Syslog tag value added to messages sent to this server |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogServer := models.RemoteSyslogServer{
        Contents:             []models.RemoteSyslogContent{
            models.RemoteSyslogContent{
                Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
            },
        },
        ExplicitPriority:     models.ToPointer(false),
        Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_CONFIG),
        Host:                 models.ToPointer("syslogd.internal"),
        Match:                models.ToPointer("!alarm|ntp|errors.crc_error[chan]"),
        Protocol:             models.ToPointer(models.RemoteSyslogServerProtocolEnum_UDP),
        RoutingInstance:      models.ToPointer("routing-instance-name"),
        ServerName:           models.ToPointer("syslogd.internal"),
        Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ANY),
    }

}
```

