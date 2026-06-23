
# Site Mxtunnel Radsec

RadSec proxy settings for a site Mist Tunnel

## Structure

`SiteMxtunnelRadsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | RADIUS accounting servers used by the site Mist Tunnel RadSec proxy |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | RADIUS authentication servers used by the site Mist Tunnel RadSec proxy |
| `Enabled` | `*bool` | Optional | Whether RadSec proxying is enabled for this site Mist Tunnel<br><br>**Default**: `false` |
| `UseMxedge` | `*bool` | Optional | Whether RadSec proxying uses Mist Edge |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteMxtunnelRadsec := models.SiteMxtunnelRadsec{
        AcctServers:          []models.RadiusAcctServer{
            models.RadiusAcctServer{
                Host:                 "host4",
                KeywrapEnabled:       models.ToPointer(false),
                KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_ASCII),
                KeywrapKek:           models.ToPointer("keywrap_kek0"),
                KeywrapMack:          models.ToPointer("keywrap_mack2"),
                Port:                 models.ToPointer(models.RadiusAcctPortContainer.FromNumber(176)),
                Secret:               "secret0",
            },
        },
        AuthServers:          []models.RadiusAuthServer{
            models.RadiusAuthServer{
                Host:                        "host0",
                KeywrapEnabled:              models.ToPointer(false),
                KeywrapFormat:               models.ToPointer(models.RadiusKeywrapFormatEnum_ASCII),
                KeywrapKek:                  models.ToPointer("keywrap_kek4"),
                KeywrapMack:                 models.ToPointer("keywrap_mack6"),
                Port:                        models.ToPointer(models.RadiusAuthPortContainer.FromNumber(36)),
                Secret:                      "secret4",
            },
        },
        Enabled:              models.ToPointer(false),
        UseMxedge:            models.ToPointer(false),
    }

}
```

