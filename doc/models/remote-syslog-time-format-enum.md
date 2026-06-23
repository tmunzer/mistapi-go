
# Remote Syslog Time Format Enum

enum: `millisecond`, `year`, `year millisecond`

## Enumeration

`RemoteSyslogTimeFormatEnum`

## Fields

| Name |
|  --- |
| `millisecond` |
| `year` |
| `year millisecond` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogTimeFormat := models.RemoteSyslogTimeFormatEnum_ENUMYEARMILLISECOND

}
```

