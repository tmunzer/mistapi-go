
# Org Setting Password Policy

password policy

## Structure

`OrgSettingPasswordPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether the policy is enabled<br>**Default**: `false` |
| `Freshness` | `*int` | Optional | days, required if password policy is enabled |
| `MinLength` | `*int` | Optional | required password length<br>**Default**: `8` |
| `RequiresSpecialChar` | `*bool` | Optional | whether to require special character<br>**Default**: `false` |
| `RequiresTwoFactorAuth` | `*bool` | Optional | whether to require two-factor auth<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "enabled": false,
  "freshness": 60,
  "min_length": 8,
  "requires_special_char": false,
  "requires_two_factor_auth": false
}
```

