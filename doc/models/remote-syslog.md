
# Remote Syslog

Remote syslog forwarding settings

## Structure

`RemoteSyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Archive` | [`*models.RemoteSyslogArchive`](../../doc/models/remote-syslog-archive.md) | Optional | Syslog file archive retention settings |
| `Cacerts` | `[]string` | Optional | CA certificates used to verify TLS syslog servers |
| `Console` | [`*models.RemoteSyslogConsole`](../../doc/models/remote-syslog-console.md) | Optional | Console log forwarding filters for remote syslog |
| `Enabled` | `*bool` | Optional | Whether remote syslog forwarding is enabled<br><br>**Default**: `false` |
| `Files` | [`[]models.RemoteSyslogFileConfig`](../../doc/models/remote-syslog-file-config.md) | Optional | Generated syslog file output definitions |
| `Network` | `*string` | Optional | Source network used for syslog traffic. If `source_address` is configured, Mist uses the VLAN first; otherwise it uses `source_ip` |
| `SendToAllServers` | `*bool` | Optional | Whether each log entry is sent to all configured remote syslog servers<br><br>**Default**: `false` |
| `Servers` | [`[]models.RemoteSyslogServer`](../../doc/models/remote-syslog-server.md) | Optional | Remote syslog server destination list |
| `TimeFormat` | [`*models.RemoteSyslogTimeFormatEnum`](../../doc/models/remote-syslog-time-format-enum.md) | Optional | enum: `millisecond`, `year`, `year millisecond` |
| `Users` | [`[]models.RemoteSyslogUser`](../../doc/models/remote-syslog-user.md) | Optional | User-specific syslog logging rules |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslog := models.RemoteSyslog{
        Archive:              models.ToPointer(models.RemoteSyslogArchive{
            Files:                models.ToPointer(models.RemoteSyslogArchiveFilesContainer.FromString("String5")),
            Size:                 models.ToPointer("size8"),
        }),
        Cacerts:              []string{
            "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
            "-----BEGIN CERTIFICATE-----\\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----",
        },
        Console:              models.ToPointer(models.RemoteSyslogConsole{
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
        }),
        Enabled:              models.ToPointer(false),
        Files:                []models.RemoteSyslogFileConfig{
            models.RemoteSyslogFileConfig{
                Archive:              models.ToPointer(models.RemoteSyslogArchive{
                    Files:                models.ToPointer(models.RemoteSyslogArchiveFilesContainer.FromString("String5")),
                    Size:                 models.ToPointer("size8"),
                }),
                Contents:             []models.RemoteSyslogContent{
                    models.RemoteSyslogContent{
                        Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                        Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
                    },
                    models.RemoteSyslogContent{
                        Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                        Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
                    },
                },
                EnableTls:            models.ToPointer(false),
                ExplicitPriority:     models.ToPointer(false),
                File:                 models.ToPointer("file4"),
            },
            models.RemoteSyslogFileConfig{
                Archive:              models.ToPointer(models.RemoteSyslogArchive{
                    Files:                models.ToPointer(models.RemoteSyslogArchiveFilesContainer.FromString("String5")),
                    Size:                 models.ToPointer("size8"),
                }),
                Contents:             []models.RemoteSyslogContent{
                    models.RemoteSyslogContent{
                        Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                        Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
                    },
                    models.RemoteSyslogContent{
                        Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_NTP),
                        Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_ENUMERROR),
                    },
                },
                EnableTls:            models.ToPointer(false),
                ExplicitPriority:     models.ToPointer(false),
                File:                 models.ToPointer("file4"),
            },
        },
        Network:              models.ToPointer("default"),
        SendToAllServers:     models.ToPointer(false),
        Servers:              []models.RemoteSyslogServer{
            models.RemoteSyslogServer{
                Facility:             models.ToPointer(models.RemoteSyslogFacilityEnum_CONFIG),
                Host:                 models.ToPointer("syslogd.internal"),
                Port:                 models.ToPointer(),
                Protocol:             models.ToPointer(models.RemoteSyslogServerProtocolEnum_UDP),
                Severity:             models.ToPointer(models.RemoteSyslogSeverityEnum_INFO),
                Tag:                  models.ToPointer(""),
            },
        },
        TimeFormat:           models.ToPointer(models.RemoteSyslogTimeFormatEnum_MILLISECOND),
    }

}
```

