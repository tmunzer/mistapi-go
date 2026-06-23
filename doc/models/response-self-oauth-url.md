
# Response Self Oauth Url

OAuth2 authorization URL response for linking the current Mist account

## Structure

`ResponseSelfOauthUrl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthorizationUrl` | `string` | Required | OAuth2 provider authorization URL used to link the Mist account |
| `Linked` | `bool` | Required | Whether the Mist account is already linked with this OAuth2 provider |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSelfOauthUrl := models.ResponseSelfOauthUrl{
        AuthorizationUrl:     "authorization_url0",
        Linked:               false,
    }

}
```

