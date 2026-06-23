
# Sso Delete Admins

Request body listing SSO admin email addresses to delete

## Structure

`SsoDeleteAdmins`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Emails` | `[]string` | Required | List of admin email addresses to delete |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoDeleteAdmins := models.SsoDeleteAdmins{
        Emails:               []string{
            "emails1",
        },
    }

}
```

