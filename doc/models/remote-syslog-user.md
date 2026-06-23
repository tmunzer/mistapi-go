
# Remote Syslog User

User-specific syslog logging rule

## Structure

`RemoteSyslogUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | List of syslog content selectors |
| `Match` | `*string` | Optional | Expression used to filter user log messages |
| `User` | `*string` | Optional | Account name or wildcard matched by this syslog rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogUser := models.RemoteSyslogUser{
        Contents:             []models.RemoteSyslogContent{
            models.RemoteSyslogContent{
                Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
            },
            models.RemoteSyslogContent{
                Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
            },
            models.RemoteSyslogContent{
                Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
            },
        },
        Match:                models.ToPointer("\"!alarm|ntp|errors.crc_error[chan]\""),
        User:                 models.ToPointer("*"),
    }

}
```

