
# Org Setting Password Policy

password policy

## Structure

`OrgSettingPasswordPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether the policy is enabled<br>**Default**: `false` |
| `ExpiryInDays` | `*int` | Optional | password expiry in days |
| `MinLength` | `*int` | Optional | required password length<br>**Default**: `8` |
| `RequiresSpecialChar` | `*bool` | Optional | whether to require special character<br>**Default**: `false` |
| `RequiresTwoFactorAuth` | `*bool` | Optional | whether to require two-factor auth<br>**Default**: `false` |

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

