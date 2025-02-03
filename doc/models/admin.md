
# Admin

*This model accepts additional fields of type interface{}.*

## Structure

`Admin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `*uuid.UUID` | Optional | ID of the administrator |
| `ComplianceStatus` | [`*models.AdminComplianceStatusEnum`](../../doc/models/admin-compliance-status-enum.md) | Optional | trade compliance status. enum: `blocked`, `restricted` |
| `Email` | `*string` | Optional | If admin account is not an Org API Token |
| `EnableTwoFactor` | `*bool` | Optional | If admin account is not an Org API Token |
| `ExpireTime` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | If admin account is not an Org API Token. For an invite, this is the original first name used |
| `Hours` | `*int` | Optional | If admin account is not an Org API Token, how long the invite should be valid<br>**Default**: `24`<br>**Constraints**: `>= 1`, `<= 168` |
| `LastName` | `*string` | Optional | If admin account is not an Org API Token. For an invite, this is the original last name used |
| `Name` | `*string` | Optional | For Org API Token Only |
| `NoTracking` | `models.Optional[bool]` | Optional | Optional, whether to store privacy-consent information. When it doesn’t exist, it’s assumed true on EU (i.e. no tracking, the user has to opt-in); otherwise, the user would have to opt-out |
| `OauthGoogle` | `*bool` | Optional | If admin account is not an Org API Token |
| `PasswordModifiedTime` | `*float64` | Optional | Password last modified time, in epoch |
| `Phone` | `*string` | Optional | If admin account is not an Org API Token. Phone number (numbers only, including country code) |
| `Phone2` | `*string` | Optional | If admin account is not an Org API Token. Secondary phone number (numbers only, including country code) |
| `Privileges` | [`[]models.AdminPrivilege`](../../doc/models/admin-privilege.md) | Optional | List of privileges the admin has<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SessionExpiry` | `*int64` | Optional | **Constraints**: `>= 10`, `<= 20160` |
| `Tags` | `[]string` | Optional | - |
| `TwoFactorVerified` | `*bool` | Optional | If admin account is not an Org API Token. Two factor status |
| `ViaSso` | `*bool` | Optional | If admin account is not an Org API Token, an admin login via_sso is more restircted. (password and email cannot be changed) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "admin_id": "456b7016-a916-a4b1-78dd-72b947c152b7",
  "email": "jsnow@abc.com",
  "first_name": "John",
  "hours": 24,
  "last_name": "Sno",
  "session_expiry": 1440,
  "compliance_status": "blocked",
  "enable_two_factor": false,
  "expire_time": 6,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

