
# Org Setting Junos Shell Access

junos_shell_access: Manages role-based web-shell access.  
When junos_shell access is not defined (Default) - No additional users are configured and web-shell uses default `mist` user to login.  
When junos_shell_access is defined - Additional users mist-web-admin (admin permission), mist-web-viewer(viewer permission) are configured on the device and web-shell logs in with the mist-web-admin/mist-web-viewer user depending upon the shell access level. Setting the shell access level to "none", disables web-shell access for that specific role.

## Structure

`OrgSettingJunosShellAccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Admin` | [`*models.OrgSettingJunosShellAccessAdminEnum`](../../doc/models/org-setting-junos-shell-access-admin-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br><br>**Default**: `"admin"` |
| `Helpdesk` | [`*models.OrgSettingJunosShellAccessHelpdeskEnum`](../../doc/models/org-setting-junos-shell-access-helpdesk-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br><br>**Default**: `"none"` |
| `Read` | [`*models.OrgSettingJunosShellAccessReadEnum`](../../doc/models/org-setting-junos-shell-access-read-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br><br>**Default**: `"none"` |
| `Write` | [`*models.OrgSettingJunosShellAccessWriteEnum`](../../doc/models/org-setting-junos-shell-access-write-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br><br>**Default**: `"admin"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingJunosShellAccess := models.OrgSettingJunosShellAccess{
        Admin:                models.ToPointer(models.OrgSettingJunosShellAccessAdminEnum_ADMIN),
        Helpdesk:             models.ToPointer(models.OrgSettingJunosShellAccessHelpdeskEnum_NONE),
        Read:                 models.ToPointer(models.OrgSettingJunosShellAccessReadEnum_NONE),
        Write:                models.ToPointer(models.OrgSettingJunosShellAccessWriteEnum_ADMIN),
    }

}
```

