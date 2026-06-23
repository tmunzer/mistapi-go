
# Remote Syslog Severity Enum

enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`

## Enumeration

`RemoteSyslogSeverityEnum`

## Fields

| Name |
|  --- |
| `alert` |
| `any` |
| `critical` |
| `emergency` |
| `error` |
| `info` |
| `notice` |
| `warning` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogSeverity := models.RemoteSyslogSeverityEnum_ENUMERROR

}
```

