
# Org Setting Password Policy

password policy

## Structure

`OrgSettingPasswordPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether the policy is enabled<br><br>**Default**: `false` |
| `ExpiryInDays` | `*int` | Optional | Password expiry in days. Password Expiry Notice banner will display in the UI 14 days before expiration<br><br>**Constraints**: `>= 1`, `<= 365` |
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

