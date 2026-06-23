
# Account Oauth Info Account Service Connection

Prisma Access service connection region for a linked OAuth account

## Structure

`AccountOauthInfoAccountServiceConnection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Region` | `*string` | Optional | Prisma Access region where this service connection is provisioned |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountOauthInfoAccountServiceConnection := models.AccountOauthInfoAccountServiceConnection{
        Region:               models.ToPointer("us-southwest"),
    }

}
```

