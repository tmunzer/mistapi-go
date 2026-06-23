
# Response Mobile Verify Secret

Mobile SDK invite verification response

## Structure

`ResponseMobileVerifySecret`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | Organization display name associated with the verified SDK invite |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `Secret` | `string` | Required | Device-specific secret returned for mobile SDK activation |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseMobileVerifySecret := models.ResponseMobileVerifySecret{
        Name:                 "name2",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Secret:               "secret2",
    }

}
```

