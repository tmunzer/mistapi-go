
# Privilege Org

Organization privilege scope and role settings

## Structure

`PrivilegeOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | If `scope`==`org`, organization ID this privilege applies to |
| `Role` | [`models.PrivilegeOrgRoleEnum`](../../doc/models/privilege-org-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.PrivilegeOrgScopeEnum`](../../doc/models/privilege-org-scope-enum.md) | Required | enum: `org`, `site`, `sitegroup`, `orgsites` |
| `SiteId` | `*uuid.UUID` | Optional | If `scope`==`site`, site ID this privilege applies to |
| `SitegroupId` | `*uuid.UUID` | Optional | If `scope`==`sitegroup`, site group ID this privilege applies to |
| `View` | `*string` | Optional | Used for backward compatibility. Use `views` instead. |
| `Views` | [`[]models.AdminPrivilegeViewEnum`](../../doc/models/admin-privilege-view-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users. Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  <br>You can define custom roles by adding the `views` attribute along with `role` when assigning privileges.  <br>Below are the list of supported UI views. Note that this is UI only feature.<br><br>\| UI View \| Required Role \| Description \|<br>\| --- \| --- \| --- \|<br>\| `reporting` \| `read` \| full access to all analytics tools \|<br>\| `marketing` \| `read` \| can view analytics and location maps \|<br>\| `super_observer` \| `read` \| can view all the organization except the subscription page \|<br>\| `location` \| `write` \| can view and manage location maps, can view analytics \|<br>\| `security` \| `write` \| can view and manage site labels, policies and security \|<br>\| `switch_admin` \| `helpdesk` \| can view and manage Switch ports, can view wired clients \|<br>\| `mxedge_admin` \| `admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| `admin` \| full access to Org and Site Pre-shared keys \| |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    privilegeOrg := models.PrivilegeOrg{
        OrgId:                models.ToPointer(uuid.MustParse("000020f8-0000-0000-0000-000000000000")),
        Role:                 models.PrivilegeOrgRoleEnum_INSTALLER,
        Scope:                models.PrivilegeOrgScopeEnum_SITEGROUP,
        SiteId:               models.ToPointer(uuid.MustParse("00001086-0000-0000-0000-000000000000")),
        SitegroupId:          models.ToPointer(uuid.MustParse("00001afe-0000-0000-0000-000000000000")),
        View:                 models.ToPointer("view2"),
        Views:                []models.AdminPrivilegeViewEnum{
            models.AdminPrivilegeViewEnum_LOCATION,
            models.AdminPrivilegeViewEnum_LOBBYADMIN,
        },
    }

}
```

