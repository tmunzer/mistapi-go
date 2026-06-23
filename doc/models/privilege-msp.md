
# Privilege Msp

MSP privilege scope and role settings

## Structure

`PrivilegeMsp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | If `scope`==`org`, organization ID this MSP privilege applies to |
| `OrgName` | `*string` | Optional, Read-only | Name of the org (for a site belonging to org) |
| `OrggroupId` | `*uuid.UUID` | Optional | If `scope`==`orggroup`, organization group ID this MSP privilege applies to |
| `Role` | [`models.PrivilegeMspRoleEnum`](../../doc/models/privilege-msp-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.PrivilegeMspScopeEnum`](../../doc/models/privilege-msp-scope-enum.md) | Required | enum: `msp`, `org`, `orggroup` |
| `Views` | [`[]models.AdminPrivilegeViewEnum`](../../doc/models/admin-privilege-view-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users. Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  <br>You can define custom roles by adding the `views` attribute along with `role` when assigning privileges.  <br>Below are the list of supported UI views. Note that this is UI only feature.<br><br>\| UI View \| Required Role \| Description \|<br>\| --- \| --- \| --- \|<br>\| `reporting` \| `read` \| full access to all analytics tools \|<br>\| `marketing` \| `read` \| can view analytics and location maps \|<br>\| `super_observer` \| `read` \| can view all the organization except the subscription page \|<br>\| `location` \| `write` \| can view and manage location maps, can view analytics \|<br>\| `security` \| `write` \| can view and manage site labels, policies and security \|<br>\| `switch_admin` \| `helpdesk` \| can view and manage Switch ports, can view wired clients \|<br>\| `mxedge_admin` \| `admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| `admin` \| full access to Org and Site Pre-shared keys \| |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    privilegeMsp := models.PrivilegeMsp{
        OrgId:                models.ToPointer(uuid.MustParse("0000230a-0000-0000-0000-000000000000")),
        OrgName:              models.ToPointer("org_name4"),
        OrggroupId:           models.ToPointer(uuid.MustParse("00000014-0000-0000-0000-000000000000")),
        Role:                 models.PrivilegeMspRoleEnum_READ,
        Scope:                models.PrivilegeMspScopeEnum_ORG,
        Views:                []models.AdminPrivilegeViewEnum{
            models.AdminPrivilegeViewEnum_MXEDGEADMIN,
        },
    }

}
```

