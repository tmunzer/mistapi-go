
# Admin

## Structure

`Admin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `*uuid.UUID` | Optional | - |
| `ComplianceStatus` | [`*models.AdminComplianceStatusEnum`](../../doc/models/admin-compliance-status-enum.md) | Optional | trade compliance status. enum: `blocked`, `restricted` |
| `Email` | `*string` | Optional | if admin account is not an Org API Token |
| `EnableTwoFactor` | `*bool` | Optional | if admin account is not an Org API Token |
| `ExpireTime` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | if admin account is not an Org API Token<br>for an invite, this is the original first name used |
| `Hours` | `*int` | Optional | if admin account is not an Org API Token, how long the invite should be valid<br>**Default**: `24`<br>**Constraints**: `>= 1`, `<= 168` |
| `LastName` | `*string` | Optional | if admin account is not an Org API Token<br>for an invite, this is the original last name used |
| `Name` | `*string` | Optional | for Org API Token Only |
| `NoTracking` | `*bool` | Optional | when it doesn’t exist, it’s assumed true on EU (i.e. no tracking, the user has to opt-in); otherwise, the user would have to opt-out |
| `OauthGoogle` | `*bool` | Optional | if admin account is not an Org API Token |
| `PasswordModifiedTime` | `*float64` | Optional | password last modified time, in epoch |
| `Phone` | `*string` | Optional | if admin account is not an Org API Token<br>phone number (numbers only, including country code) |
| `Phone2` | `*string` | Optional | if admin account is not an Org API Token<br>secondary phone number (numbers only, including country code) |
| `Privileges` | [`[]models.PrivilegeSelf`](../../doc/models/privilege-self.md) | Optional | list of privileges the admin has<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SessionExpiry` | `*int64` | Optional | **Constraints**: `>= 10`, `<= 20160` |
| `Tags` | `[]string` | Optional | - |
| `TwoFactorVerified` | `*bool` | Optional | if admin account is not an Org API Token<br>two factor status |
| `ViaSso` | `*bool` | Optional | if admin account is not an Org API Token<br>an admin login via_sso is more restircted. (password and email<br>cannot be changed) |

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
  "expire_time": 6
}
```

