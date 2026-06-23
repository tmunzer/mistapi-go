
# Sso Mxedge Proxy Auth Server

RADIUS authentication server for the Mist Edge SSO proxy

## Structure

`SsoMxedgeProxyAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | RADIUS authentication server hostname or IP address |
| `Port` | `*int` | Optional | UDP port for RADIUS authentication requests<br><br>**Default**: `1812` |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br><br>**Default**: `false` |
| `Retry` | `*int` | Optional | Number of retry attempts for RADIUS authentication requests<br><br>**Default**: `2` |
| `Secret` | `*string` | Optional | Shared secret used with the RADIUS authentication server |
| `Timeout` | `*int` | Optional | Authentication request timeout, in seconds<br><br>**Default**: `5` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoMxedgeProxyAuthServer := models.SsoMxedgeProxyAuthServer{
        Host:                        models.ToPointer("1.2.3.4"),
        Port:                        models.ToPointer(1812),
        RequireMessageAuthenticator: models.ToPointer(false),
        Retry:                       models.ToPointer(2),
        Secret:                      models.ToPointer("testing123"),
        Timeout:                     models.ToPointer(5),
    }

}
```

