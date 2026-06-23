
# Site Setting Juniper Srx

Site-level Juniper SRX integration settings

## Structure

`SiteSettingJuniperSrx`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.JuniperSrxAutoUpgrade`](../../doc/models/juniper-srx-auto-upgrade.md) | Optional | SRX firmware auto-upgrade settings applied when a device is first onboarded |
| `Gateways` | [`[]models.SiteSettingJuniperSrxGateway`](../../doc/models/site-setting-juniper-srx-gateway.md) | Optional | Juniper SRX gateways integrated with a site |
| `SendMistNacUserInfo` | `*bool` | Optional | Whether Mist NAC user information is sent to Juniper SRX gateways |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingJuniperSrx := models.SiteSettingJuniperSrx{
        AutoUpgrade:          models.ToPointer(models.JuniperSrxAutoUpgrade{
            CustomVersions:       map[string]string{
                "key0": "custom_versions3",
                "key1": "custom_versions2",
            },
            Enabled:              models.ToPointer(false),
            Snapshot:             models.ToPointer(false),
            Version:              models.ToPointer("version2"),
        }),
        Gateways:             []models.SiteSettingJuniperSrxGateway{
            models.SiteSettingJuniperSrxGateway{
                ApiKey:               models.ToPointer("api_key8"),
                ApiPassword:          models.ToPointer("api_password8"),
                ApiUrl:               models.ToPointer("api_url0"),
            },
        },
        SendMistNacUserInfo:  models.ToPointer(false),
    }

}
```

