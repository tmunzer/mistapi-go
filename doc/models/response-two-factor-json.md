
# Response Two Factor Json

Response body containing a generated two-factor authentication secret

## Structure

`ResponseTwoFactorJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TwoFactorSecret` | `*string` | Optional | Generated secret key used to enroll a two-factor authentication app |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseTwoFactorJson := models.ResponseTwoFactorJson{
        TwoFactorSecret:      models.ToPointer("NRMTSTRWNBVECY3GJVYEY3DDJFRGSNCZGJUDO4RVN5FDM3DUMJSA"),
    }

}
```

