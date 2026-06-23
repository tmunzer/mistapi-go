
# Response Login Success

Login response body, empty on normal success or populated with two-factor state

## Structure

`ResponseLoginSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `*string` | Optional | Admin email address for a login flow that requires two-factor authentication |
| `TwoFactorPassed` | `*bool` | Optional | Whether the supplied two-factor code has been accepted for this login |
| `TwoFactorRequired` | `*bool` | Optional | Whether this login requires a two-factor code before a session is established |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseLoginSuccess := models.ResponseLoginSuccess{
        Email:                models.ToPointer("email4"),
        TwoFactorPassed:      models.ToPointer(false),
        TwoFactorRequired:    models.ToPointer(false),
    }

}
```

