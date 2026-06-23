
# Org Setting Password Policy

Admin credential policy settings for the organization

## Structure

`OrgSettingPasswordPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether the policy is enabled<br><br>**Default**: `false` |
| `ExpiryInDays` | `*int` | Optional | Password expiry in days. Password Expiry Notice banner will display in the UI 14 days before expiration<br><br>**Constraints**: `>= 1`, `<= 365` |
| `MinLength` | `*int` | Optional | Minimum number of characters required for passwords<br><br>**Default**: `8` |
| `RequiresSpecialChar` | `*bool` | Optional | Whether to require special character<br><br>**Default**: `false` |
| `RequiresTwoFactorAuth` | `*bool` | Optional | Whether to require two-factor auth<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingPasswordPolicy := models.OrgSettingPasswordPolicy{
        Enabled:               models.ToPointer(false),
        ExpiryInDays:          models.ToPointer(60),
        MinLength:             models.ToPointer(8),
        RequiresSpecialChar:   models.ToPointer(false),
        RequiresTwoFactorAuth: models.ToPointer(false),
    }

}
```

