
# Switch Radius Config

Switch RADIUS authentication and accounting configuration

## Structure

`SwitchRadiusConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctImmediateUpdate` | `*bool` | Optional | Whether immediate RADIUS accounting updates are sent |
| `AcctInterimInterval` | `*int` | Optional | How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the RADIUS server, 600 and up is recommended when enabled<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | List of RADIUS accounting servers<br><br>**Constraints**: *Unique Items Required* |
| `AuthServerSelection` | [`*models.SwitchRadiusConfigAuthServerSelectionEnum`](../../doc/models/switch-radius-config-auth-server-selection-enum.md) | Optional | Selection strategy for RADIUS authentication servers. enum: `ordered`, `unordered`<br><br>**Default**: `"ordered"` |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | List of RADIUS authentication servers<br><br>**Constraints**: *Unique Items Required* |
| `AuthServersRetries` | `*int` | Optional | RADIUS auth session retries<br><br>**Default**: `3` |
| `AuthServersTimeout` | `*int` | Optional | RADIUS auth session timeout<br><br>**Default**: `5` |
| `CoaEnabled` | `*bool` | Optional | Whether RADIUS Change of Authorization (CoA) is enabled<br><br>**Default**: `false` |
| `CoaPort` | [`*models.RadiusCoaPort`](../../doc/models/containers/radius-coa-port.md) | Optional | RADIUS CoA Port, value from 1 to 65535, default is 3799 |
| `FastDot1xTimers` | `*bool` | Optional | Whether fast 802.1X timers are enabled for RADIUS authentication<br><br>**Default**: `false` |
| `Network` | `*string` | Optional | Use `network`or `source_ip`. Which network the RADIUS server resides, if there's static IP for this network, we'd use it as source-ip |
| `SourceIp` | `*string` | Optional | Use `network` or `source_ip`. Explicit source IP address for RADIUS traffic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchRadiusConfig := models.SwitchRadiusConfig{
        AcctImmediateUpdate:  models.ToPointer(false),
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
        AuthServerSelection:  models.ToPointer(models.SwitchRadiusConfigAuthServerSelectionEnum_ORDERED),
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
        FastDot1xTimers:      models.ToPointer(false),
    }

}
```

