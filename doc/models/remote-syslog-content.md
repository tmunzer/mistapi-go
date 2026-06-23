
# Remote Syslog Content

Syslog message content selector for remote logging

## Structure

`RemoteSyslogContent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Facility` | [`*models.RemoteSyslogFacilityEnum`](../../doc/models/remote-syslog-facility-enum.md) | Optional | enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`<br><br>**Default**: `"any"` |
| `Severity` | [`*models.RemoteSyslogSeverityEnum`](../../doc/models/remote-syslog-severity-enum.md) | Optional | enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`<br><br>**Default**: `"any"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogContent := models.RemoteSyslogContent{
        Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_CONFIG),
        Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ANY),
    }

}
```

