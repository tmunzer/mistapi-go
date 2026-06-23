
# Custom Header Signature



Documentation for accessing and setting credentials for csrfToken.

## Auth Credentials

| Name | Type | Description | Setter | Getter |
|  --- | --- | --- | --- | --- |
| xCSRFToken | `string` | Session-based authentication for browser or login/password flows. After a successful [Login](../../doc/controllers/admins-login.md#login) request, Mist returns a `csrftoken` cookie. Send that value in the `X-CSRFToken` header on later API requests that use the login session.<br><br>**Format**:<br><br>```<br>X-CSRFToken: vwvBuq9qkqaKh7lu8tNc0gkvBfEaLAmx<br>```<br><br>For automation, API Token authentication is preferred. | `WithXCSRFToken` | `XCSRFToken()` |



**Note:** Required auth credentials can be set using `WithCsrfTokenCredentials()` by providing a credentials instance with `NewCsrfTokenCredentials()` in the configuration initialization and accessed using the `CsrfTokenCredentials()` method in the configuration instance.

## Usage Example

### Client Initialization

You must provide credentials in the client as shown in the following code snippet.

```go
package main

import (
    "mistapi"
)

func main() {
    client := mistapi.NewClient(
    mistapi.CreateConfiguration(
            mistapi.WithCsrfTokenCredentials(
                mistapi.NewCsrfTokenCredentials("X-CSRFToken"),
            ),
        ),
    )
}
```


