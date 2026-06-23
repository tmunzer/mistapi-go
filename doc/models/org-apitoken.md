
# Org Apitoken

Organization API token with scoped privileges

**Note:**
`privileges` is required to create the object, but may not be returned in the POST API response. Retrieve the token afterward to inspect it.

*This model accepts additional fields of type interface{}.*

## Structure

`OrgApitoken`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedBy` | `models.Optional[string]` | Optional, Read-only | email of the token creator / null if creator is deleted |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Key` | `*string` | Optional, Read-only | Token secret key. The full API Token is only returned when the API token is created and can only be partially retrieved afterward |
| `LastUsed` | `models.Optional[float64]` | Optional, Read-only | Epoch timestamp when the token was last used, or null if it has not been used |
| `Name` | `string` | Required | Display name of the organization API token |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Privileges` | [`[]models.PrivilegeOrg`](../../doc/models/privilege-org.md) | Optional | List of privileges the token has on the orgs/sites<br><br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `10`, *Unique Items Required* |
| `SrcIps` | `[]string` | Optional | List of allowed IP addresses from where the token can be used from. At most 10 IP addresses can be specified, cannot be changed once the API Token is created. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgApitoken := models.OrgApitoken{
        CreatedBy:            models.NewOptional(models.ToPointer("user@mycorp.com")),
        CreatedTime:          models.ToPointer(float64(88.34)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Key:                  models.ToPointer("1qkb...QQCL"),
        LastUsed:             models.NewOptional(models.ToPointer(float64(1690115110))),
        Name:                 "org_token_xyz",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Privileges:           []models.PrivilegeOrg{
            models.PrivilegeOrg{
                Role:                 models.PrivilegeOrgRoleEnum_ADMIN,
                Scope:                models.PrivilegeOrgScopeEnum_ORG,
            },
        },
        SrcIps:               []string{
            "63.3.56.0/24",
            "63.3.55.4",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

