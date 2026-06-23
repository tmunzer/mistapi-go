
# Radius Config

Junos RADIUS authentication and accounting configuration

## Structure

`RadiusConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctInterimInterval` | `*int` | Optional | How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the RADIUS server, 600 and up is recommended when enabled<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | List of RADIUS accounting servers<br><br>**Constraints**: *Unique Items Required* |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | List of RADIUS authentication servers<br><br>**Constraints**: *Unique Items Required* |
| `AuthServersRetries` | `*int` | Optional | Number of RADIUS authentication request retries before failover<br><br>**Default**: `3` |
| `AuthServersTimeout` | `*int` | Optional | RADIUS authentication server timeout, in seconds<br><br>**Default**: `5` |
| `CoaEnabled` | `*bool` | Optional | Whether RADIUS Change of Authorization (CoA) is enabled<br><br>**Default**: `false` |
| `CoaPort` | `*int` | Optional | UDP port used for RADIUS Change of Authorization (CoA)<br><br>**Default**: `3799`<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `Network` | `*string` | Optional | Use `network` or `source_ip`. Network where the RADIUS server resides; if the network has a static IP, Mist uses it as the source IP |
| `SourceIp` | `*string` | Optional | Use `network` or `source_ip`. Explicit source IP address for RADIUS traffic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    radiusConfig := models.RadiusConfig{
        AcctInterimInterval:  models.ToPointer(0),
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
        AuthServersRetries:   models.ToPointer(3),
        AuthServersTimeout:   models.ToPointer(5),
        CoaEnabled:           models.ToPointer(false),
        CoaPort:              models.ToPointer(3799),
    }

}
```

