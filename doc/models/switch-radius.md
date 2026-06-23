
# Switch Radius

Switch RADIUS override settings. By default, the switch `radius_config` is used; set `use_different_radius` to use alternate RADIUS settings.

## Structure

`SwitchRadius`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether RADIUS is enabled for switch management authentication |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Switch RADIUS authentication and accounting configuration |
| `UseDifferentRadius` | `*string` | Optional | Selector for alternate RADIUS settings instead of the default switch `radius_config` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchRadius := models.SwitchRadius{
        Enabled:              models.ToPointer(false),
        RadiusConfig:         models.ToPointer(models.SwitchRadiusConfig{
            AcctImmediateUpdate:  models.ToPointer(false),
            AcctInterimInterval:  models.ToPointer(118),
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
        }),
        UseDifferentRadius:   models.ToPointer("use_different_radius0"),
    }

}
```

