
# Account Skyatp Config

Sky ATP account credentials and realm settings used for integration

*This model accepts additional fields of type interface{}.*

## Structure

`AccountSkyatpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | [`*models.AccountSkyatpCloudNameEnum`](../../doc/models/account-skyatp-cloud-name-enum.md) | Optional | Sky ATP cloud name. enum: `www.amerskyatp.com`, `www.apacskyatp.com`, `www.euroskyatp.com`, `www.canadaskyatp.com` |
| `Password` | `string` | Required | Credential password for the Sky ATP realm user |
| `Realm` | `string` | Required | Sky ATP realm to create or link with this Mist organization |
| `Username` | `string` | Required | Sky ATP username used to create or access the realm |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountSkyatpConfig := models.AccountSkyatpConfig{
        CloudName:            models.ToPointer(models.AccountSkyatpCloudNameEnum_ENUMWWWAMERSKYATPCOM),
        Password:             "foryoureyesonly",
        Realm:                "mist-team",
        Username:             "john@abc.com",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

