
# Recaptcha

CAPTCHA settings returned for admin registration

## Structure

`Recaptcha`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Flavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | CAPTCHA provider flavor. enum: `google`, `hcaptcha`<br><br>**Default**: `"google"` |
| `Required` | `*bool` | Optional | Whether CAPTCHA verification is required for registration |
| `Sitekey` | `*string` | Optional | Public site key used to render the selected CAPTCHA provider |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    recaptcha := models.Recaptcha{
        Flavor:               models.ToPointer(models.RecaptchaFlavorEnum_HCAPTCHA),
        Required:             models.ToPointer(false),
        Sitekey:              models.ToPointer("sitekey0"),
    }

}
```

