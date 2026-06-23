
# Response Logout

Logout response with optional SSO forwarding URL

## Structure

`ResponseLogout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForwardUrl` | `*string` | Optional | If configured in SSO as custom_logout_url |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseLogout := models.ResponseLogout{
        ForwardUrl:           models.ToPointer("forward_url6"),
    }

}
```

