
# Remote Syslog Facility Enum

enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`

## Enumeration

`RemoteSyslogFacilityEnum`

## Fields

| Name |
|  --- |
| `any` |
| `authorization` |
| `change-log` |
| `config` |
| `conflict-log` |
| `daemon` |
| `dfc` |
| `external` |
| `firewall` |
| `ftp` |
| `interactive-commands` |
| `kernel` |
| `ntp` |
| `pfe` |
| `security` |
| `user` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogFacility := models.RemoteSyslogFacilityEnum_CHANGELOG

}
```

