
# Remote Syslog Archive

Syslog file archive retention settings

## Structure

`RemoteSyslogArchive`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Files` | [`*models.RemoteSyslogArchiveFiles`](../../doc/models/containers/remote-syslog-archive-files.md) | Optional | Number of archived syslog files to retain |
| `Size` | `*string` | Optional | Maximum size of each archived syslog file, such as 5m |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogArchive := models.RemoteSyslogArchive{
        Files:                models.ToPointer(models.RemoteSyslogArchiveFilesContainer.FromNumber(20)),
        Size:                 models.ToPointer("5m"),
    }

}
```

