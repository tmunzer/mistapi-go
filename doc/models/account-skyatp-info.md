
# Account Skyatp Info

Linked Sky ATP account and realm information

## Structure

`AccountSkyatpInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | [`*models.AccountSkyatpCloudNameEnum`](../../doc/models/account-skyatp-cloud-name-enum.md) | Optional | Sky ATP cloud name. enum: `www.amerskyatp.com`, `www.apacskyatp.com`, `www.euroskyatp.com`, `www.canadaskyatp.com` |
| `Realm` | `*string` | Optional | Sky ATP realm linked with this Mist organization |
| `Username` | `*string` | Optional | Sky ATP username configured for the linked realm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountSkyatpInfo := models.AccountSkyatpInfo{
        CloudName:            models.ToPointer(models.AccountSkyatpCloudNameEnum_ENUMWWWAMERSKYATPCOM),
        Realm:                models.ToPointer("mist-team"),
        Username:             models.ToPointer("john@abc.com"),
    }

}
```

