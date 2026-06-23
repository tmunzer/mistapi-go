
# Avprofile Protocol Enum

Protocol that can be inspected by an antivirus profile. enum: `ftp`, `http`, `imap`, `pop3`, `smtp`

## Enumeration

`AvprofileProtocolEnum`

## Fields

| Name |
|  --- |
| `ftp` |
| `http` |
| `imap` |
| `pop3` |
| `smtp` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    avprofileProtocol := models.AvprofileProtocolEnum_SMTP

}
```

