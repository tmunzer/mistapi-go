
# Remote Syslog Console

Console log forwarding filters for remote syslog

## Structure

`RemoteSyslogConsole`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | List of syslog content selectors |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogConsole := models.RemoteSyslogConsole{
        Contents:             []models.RemoteSyslogContent{
            models.RemoteSyslogContent{
                Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
            },
        },
    }

}
```

