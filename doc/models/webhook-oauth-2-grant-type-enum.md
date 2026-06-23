
# Webhook Oauth 2 Grant Type Enum

required when `type`==`oauth2`. enum: `client_credentials`, `password`

## Enumeration

`WebhookOauth2GrantTypeEnum`

## Fields

| Name |
|  --- |
| `client_credentials` |
| `password` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookOauth2GrantType := models.WebhookOauth2GrantTypeEnum_CLIENTCREDENTIALS

}
```

