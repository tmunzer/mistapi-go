
# Remote Syslog File Config

Generated syslog file output settings

## Structure

`RemoteSyslogFileConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Archive` | [`*models.RemoteSyslogArchive`](../../doc/models/remote-syslog-archive.md) | Optional | Syslog file archive retention settings |
| `Contents` | [`[]models.RemoteSyslogContent`](../../doc/models/remote-syslog-content.md) | Optional | List of syslog content selectors |
| `EnableTls` | `*bool` | Optional | Only if `protocol`==`tcp`, enable TLS for this syslog file destination |
| `ExplicitPriority` | `*bool` | Optional | Whether to include explicit syslog priority values in file output |
| `File` | `*string` | Optional | Generated syslog file name |
| `Match` | `*string` | Optional | Expression used to filter log messages written to this file |
| `StructuredData` | `*bool` | Optional | Whether to include structured syslog data in file output |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogFileConfig := models.RemoteSyslogFileConfig{
        Archive:              models.ToPointer(models.RemoteSyslogArchive{
            Files:                models.ToPointer(models.RemoteSyslogArchiveFilesContainer.FromString("String5")),
            Size:                 models.ToPointer("size8"),
        }),
        Contents:             []models.RemoteSyslogContent{
            models.RemoteSyslogContent{
                Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
            },
        },
        EnableTls:            models.ToPointer(false),
        ExplicitPriority:     models.ToPointer(false),
        File:                 models.ToPointer("file-name"),
        Match:                models.ToPointer("!alarm|ntp|errors.crc_error[chan]"),
    }

}
```

