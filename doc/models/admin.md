
# Admin

Mist administrator account or organization API token details

*This model accepts additional fields of type interface{}.*

## Structure

`Admin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `*uuid.UUID` | Optional, Read-only | ID of the administrator |
| `ComplianceStatus` | [`*models.AdminComplianceStatusEnum`](../../doc/models/admin-compliance-status-enum.md) | Optional | trade compliance status. enum: `blocked`, `restricted` |
| `Email` | `*string` | Optional | If admin account is not an Org API Token |
| `EnableTwoFactor` | `*bool` | Optional, Read-only | If admin account is not an Org API Token |
| `ExpireTime` | `*int` | Optional | Expiration time for the admin invitation, in epoch seconds |
| `FirstName` | `*string` | Optional | If admin account is not an Org API Token. For an invite, this is the original first name used |
| `Hours` | `*int` | Optional | If admin account is not an Org API Token, how long the invite should be valid<br><br>**Default**: `24`<br><br>**Constraints**: `>= 1`, `<= 168` |
| `LastName` | `*string` | Optional | If admin account is not an Org API Token. For an invite, this is the original last name used |
| `Name` | `*string` | Optional | For Org API Token Only |
| `NoTracking` | `models.Optional[bool]` | Optional | Optional, whether to store privacy-consent information. When it doesn’t exist, it’s assumed true on EU (i.e. no tracking, the user has to opt-in); otherwise, the user would have to opt-out |
| `OauthGoogle` | `*bool` | Optional, Read-only | If admin account is not an Org API Token |
| `PasswordModifiedTime` | `*float64` | Optional | Password last modified time, in epoch |
| `Phone` | `*string` | Optional | If admin account is not an Org API Token. Phone number (numbers only, including country code) |
| `Phone2` | `*string` | Optional | If admin account is not an Org API Token. Secondary phone number (numbers only, including country code) |
| `Privileges` | [`[]models.AdminPrivilege`](../../doc/models/admin-privilege.md) | Optional | List of privileges the admin has<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SessionExpiry` | `*int64` | Optional, Read-only | Session lifetime for the admin, in minutes<br><br>**Constraints**: `>= 10`, `<= 20160` |
| `Tags` | `[]string` | Optional, Read-only | Read-only tags associated with an administrator account |
| `TwoFactorVerified` | `*bool` | Optional, Read-only | If admin account is not an Org API Token. Two factor status |
| `ViaSso` | `*bool` | Optional, Read-only | If admin account is not an Org API Token, an admin login via_sso is more restircted. (password and email cannot be changed) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    admin := models.Admin{
        AdminId:              models.ToPointer(uuid.MustParse("456b7016-a916-a4b1-78dd-72b947c152b7")),
        ComplianceStatus:     models.ToPointer(models.AdminComplianceStatusEnum_BLOCKED),
        Email:                models.ToPointer("jsnow@abc.com"),
        EnableTwoFactor:      models.ToPointer(false),
        ExpireTime:           models.ToPointer(6),
        FirstName:            models.ToPointer("John"),
        Hours:                models.ToPointer(24),
        LastName:             models.ToPointer("Sno"),
        SessionExpiry:        models.ToPointer(int64(1440)),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

