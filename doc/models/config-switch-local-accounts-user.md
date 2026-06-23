
# Config Switch Local Accounts User

Local switch user account credentials and access role

## Structure

`ConfigSwitchLocalAccountsUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `*string` | Optional | Local password for the switch user account |
| `Role` | [`*models.ConfigSwitchLocalAccountsUserRoleEnum`](../../doc/models/config-switch-local-accounts-user-role-enum.md) | Optional | enum: `admin`, `helpdesk`, `none`, `read`<br><br>**Default**: `"none"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    configSwitchLocalAccountsUser := models.ConfigSwitchLocalAccountsUser{
        Password:             models.ToPointer("Juniper123"),
        Role:                 models.ToPointer(models.ConfigSwitchLocalAccountsUserRoleEnum_NONE),
    }

}
```

