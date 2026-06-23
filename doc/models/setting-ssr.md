
# Setting Ssr

SSR management settings for device onboarding and connectivity

## Structure

`SettingSsr`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.SettingSsrAutoUpgrade`](../../doc/models/setting-ssr-auto-upgrade.md) | Optional | Automatic firmware upgrade settings applied when an SSR device is first onboarded |
| `ConductorHosts` | `[]string` | Optional | Conductor IP addresses or hostnames used by SSR devices |
| `ConductorToken` | `*string` | Optional | Registration token used by SSR devices to connect to the conductor |
| `DisableStats` | `*bool` | Optional | Whether stats collection is disabled on SSR devices |
| `Proxy` | [`*models.SsrProxy`](../../doc/models/ssr-proxy.md) | Optional | SSR proxy configuration to talk to Mist |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    settingSsr := models.SettingSsr{
        AutoUpgrade:          models.ToPointer(models.SettingSsrAutoUpgrade{
            Channel:              models.ToPointer(models.SsrUpgradeChannelEnum_BETA),
            CustomVersions:       map[string]string{
                "key0": "custom_versions3",
                "key1": "custom_versions2",
            },
            Enabled:              models.ToPointer(false),
            Version:              models.ToPointer("version2"),
        }),
        ConductorHosts:       []string{
            "conductor_hosts0",
        },
        ConductorToken:       models.ToPointer("conductor_token2"),
        DisableStats:         models.ToPointer(false),
        Proxy:                models.ToPointer(models.SsrProxy{
            Disabled:             models.ToPointer(false),
            Url:                  models.ToPointer("url6"),
        }),
    }

}
```

