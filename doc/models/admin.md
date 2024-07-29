
# Admin

## Structure

`Admin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `*uuid.UUID` | Optional | - |
| `ComplianceStatus` | [`*models.AdminComplianceStatusEnum`](../../doc/models/admin-compliance-status-enum.md) | Optional | trade compliance status. enum: `blocked`, `restricted` |
| `Email` | `string` | Required | - |
| `EnableTwoFactor` | `*bool` | Optional | - |
| `ExpireTime` | `*int` | Optional | - |
| `FirstName` | `string` | Required | for an invite, this is the original first name used |
| `Hours` | `*int` | Optional | how long the invite should be valid<br>**Default**: `24`<br>**Constraints**: `>= 1`, `<= 168` |
| `LastName` | `string` | Required | for an invite, this is the original last name used |
| `OauthGoogle` | `*bool` | Optional | - |
| `Phone` | `*string` | Optional | phone number (numbers only, including country code) |
| `Phone2` | `*string` | Optional | secondary phone number (numbers only, including country code) |
| `Privileges` | [`[]models.PrivilegeSelf`](../../doc/models/privilege-self.md) | Optional | list of privileges the admin has<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SessionExpiry` | `*int64` | Optional | **Constraints**: `>= 10`, `<= 20160` |
| `Tags` | `[]string` | Optional | - |
| `TwoFactorVerified` | `*bool` | Optional | two factor status |
| `ViaSso` | `*bool` | Optional | an admin alogin via_sso is more restircted. (password and email cannot be changed) |

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
  "expire_time": 242
}
```

