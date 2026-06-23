
# Recover

Password recovery request submitted for an administrator account

*This model accepts additional fields of type interface{}.*

## Structure

`Recover`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `string` | Required | Admin email address requesting password recovery<br><br>**Constraints**: *Maximum Length*: `64` |
| `Recaptcha` | `*string` | Optional | CAPTCHA verification token submitted with the recovery request |
| `RecaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | CAPTCHA provider flavor. enum: `google`, `hcaptcha`<br><br>**Default**: `"google"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    recover := models.Recover{
        Email:                "test@mistsys.com",
        Recaptcha:            models.ToPointer("recaptcha4"),
        RecaptchaFlavor:      models.ToPointer(models.RecaptchaFlavorEnum_HCAPTCHA),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

