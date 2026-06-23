
# Site Setting Paloalto Networks

Palo Alto Networks integration settings for the site

## Structure

`SiteSettingPaloaltoNetworks`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateways` | [`[]models.SiteSettingPaloaltoNetworkGateway`](../../doc/models/site-setting-paloalto-network-gateway.md) | Optional | Palo Alto Networks gateways integrated with a site |
| `SendMistNacUserInfo` | `*bool` | Optional | Whether Mist NAC user information is sent to Palo Alto Networks gateways<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingPaloaltoNetworks := models.SiteSettingPaloaltoNetworks{
        Gateways:             []models.SiteSettingPaloaltoNetworkGateway{
            models.SiteSettingPaloaltoNetworkGateway{
                ApiKey:               models.ToPointer("api_key8"),
                ApiUrl:               models.ToPointer("api_url0"),
            },
        },
        SendMistNacUserInfo:  models.ToPointer(false),
    }

}
```

