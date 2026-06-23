
# Response Verify Token Success

Successful registration token verification response

## Structure

`ResponseVerifyTokenSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Result message returned with the verification response |
| `InviteNotApplied` | `*bool` | Optional | Whether the invitation was verified but not applied automatically |
| `MinLength` | `*int` | Optional | Required minimum password length from the applicable password policy |
| `ReturnTo` | `*string` | Optional | URL to redirect the user to after successful registration verification |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseVerifyTokenSuccess := models.ResponseVerifyTokenSuccess{
        Detail:               models.ToPointer("detail0"),
        InviteNotApplied:     models.ToPointer(false),
        MinLength:            models.ToPointer(142),
        ReturnTo:             models.ToPointer("return_to0"),
    }

}
```

