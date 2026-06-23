
# Sso Role Org

Organization SSO role definition

*This model accepts additional fields of type interface{}.*

## Structure

`SsoRoleOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this organization SSO role is scoped for site-level access |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `Name` | `string` | Required | Display name of the organization SSO role |
| `OrgId` | `*uuid.UUID` | Optional | Owning organization identifier for this SSO role |
| `Privileges` | [`[]models.PrivilegeOrg`](../../doc/models/privilege-org.md) | Required | Access privileges granted by an organization SSO role<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    ssoRoleOrg := models.SsoRoleOrg{
        CreatedTime:          models.ToPointer(float64(121.02)),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(213.94)),
        MspId:                models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        Name:                 "name2",
        OrgId:                models.ToPointer(uuid.MustParse("60f6bfdb-2f45-4022-8e2a-e00d977953fe")),
        Privileges:           []models.PrivilegeOrg{
            models.PrivilegeOrg{
                OrgId:                models.ToPointer(uuid.MustParse("00000cc8-0000-0000-0000-000000000000")),
                Role:                 models.PrivilegeOrgRoleEnum_ADMIN,
                Scope:                models.PrivilegeOrgScopeEnum_SITEGROUP,
                SiteId:               models.ToPointer(uuid.MustParse("00002366-0000-0000-0000-000000000000")),
                SitegroupId:          models.ToPointer(uuid.MustParse("000006ce-0000-0000-0000-000000000000")),
                View:                 models.ToPointer("view4"),
                Views:                []models.AdminPrivilegeViewEnum{
                    models.AdminPrivilegeViewEnum_LOCATION,
                    models.AdminPrivilegeViewEnum_LOBBYADMIN,
                    models.AdminPrivilegeViewEnum_SWITCHADMIN,
                },
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

