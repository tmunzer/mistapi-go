
# Response Self Oauth Link Success

OAuth2 account-linking success response

## Structure

`ResponseSelfOauthLinkSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | `string` | Required | Completed OAuth2 account-linking action |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSelfOauthLinkSuccess := models.ResponseSelfOauthLinkSuccess{
        Action:               "action6",
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
    }

}
```

