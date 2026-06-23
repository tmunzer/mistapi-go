
# Account Zscaler Info

OAuth linked Zscaler apps account details

## Structure

`AccountZscalerInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `*string` | Optional | Zscaler Internet Access cloud name configured for the integration |
| `PartnerKey` | `*string` | Optional | Zscaler partner key configured for the Mist integration |
| `Username` | `*string` | Optional | Zscaler partner administrator username configured for Mist |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountZscalerInfo := models.AccountZscalerInfo{
        CloudName:            models.ToPointer("zscalerbeta.net"),
        PartnerKey:           models.ToPointer("K35vrZcK3JvrZc"),
        Username:             models.ToPointer("john@nmo.com"),
    }

}
```

