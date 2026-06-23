
# Site Setting Juniper Srx Gateway

Juniper SRX gateway API connection settings

## Structure

`SiteSettingJuniperSrxGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiKey` | `*string` | Optional | Authentication key used to access the Juniper SRX gateway API |
| `ApiPassword` | `*string` | Optional | Authentication password used to access the Juniper SRX gateway API |
| `ApiUrl` | `*string` | Optional | Base URL for the Juniper SRX gateway API |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingJuniperSrxGateway := models.SiteSettingJuniperSrxGateway{
        ApiKey:               models.ToPointer("5abf7c8a-1a1c-4398-ba2d-b0c297094d1a"),
        ApiPassword:          models.ToPointer("abc@123"),
        ApiUrl:               models.ToPointer("https://23.43.12.78:8443"),
    }

}
```

