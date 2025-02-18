
# Org Setting Password Policy

password policy

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingPasswordPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether the policy is enabled<br>**Default**: `false` |
| `ExpiryInDays` | `*int` | Optional | password expiry in days |
| `MinLength` | `*int` | Optional | Required password length<br>**Default**: `8` |
| `RequiresSpecialChar` | `*bool` | Optional | Whether to require special character<br>**Default**: `false` |
| `RequiresTwoFactorAuth` | `*bool` | Optional | Whether to require two-factor auth<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "expiry_in_days": 60,
  "min_length": 8,
  "requires_special_char": false,
  "requires_two_factor_auth": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

