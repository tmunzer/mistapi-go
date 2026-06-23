
# Response Login Lookup

Login lookup response indicating whether SSO is available for the account

## Structure

`ResponseLoginLookup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SsoUrl` | `*string` | Optional | URL for SSO login when the account must authenticate through an identity provider |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseLoginLookup := models.ResponseLoginLookup{
        SsoUrl:               models.ToPointer("sso_url4"),
    }

}
```

