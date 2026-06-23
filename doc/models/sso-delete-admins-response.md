
# Sso Delete Admins Response

Result of deleting SSO admin accounts

## Structure

`SsoDeleteAdminsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Deleted` | `[]string` | Optional | List of email addresses that were successfully deleted |
| `Errors` | `[]string` | Optional | List of error messages for emails that could not be deleted |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoDeleteAdminsResponse := models.SsoDeleteAdminsResponse{
        Deleted:              []string{
            "deleted4",
            "deleted5",
        },
        Errors:               []string{
            "errors7",
            "errors8",
            "errors9",
        },
    }

}
```

