
# Response Login Oauth Url

OAuth2 authorization URL response for login

## Structure

`ResponseLoginOauthUrl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthorizationUrl` | `string` | Required | OAuth2 provider authorization URL to open for login |
| `ClientId` | `string` | Required | OAuth2 client identifier used for the authorization request |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseLoginOauthUrl := models.ResponseLoginOauthUrl{
        AuthorizationUrl:     "authorization_url2",
        ClientId:             "client_id4",
    }

}
```

