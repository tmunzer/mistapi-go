
# Org Setting Password Policy

password policy

## Structure

`OrgSettingPasswordPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether the policy is enabled<br><br>**Default**: `false` |
| `ExpiryInDays` | `*int` | Optional | password expiry in days<br><br>**Constraints**: `<= 365` |
| `MinLength` | `*int` | Optional | Required password length<br><br>**Default**: `8` |
| `RequiresSpecialChar` | `*bool` | Optional | Whether to require special character<br><br>**Default**: `false` |
| `RequiresTwoFactorAuth` | `*bool` | Optional | Whether to require two-factor auth<br><br>**Default**: `false` |

## Example (as JSON)

```json
{
  "enabled": false,
  "expiry_in_days": 60,
  "min_length": 8,
  "requires_special_char": false,
  "requires_two_factor_auth": false
}
```

