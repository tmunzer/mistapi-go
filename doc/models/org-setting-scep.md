
# Org Setting Scep

Mist SCEP settings for certificate enrollment

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingScep`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CertProviders` | [`[]models.OrgSettingScepCertProviderEnum`](../../doc/models/org-setting-scep-cert-provider-enum.md) | Optional | List of SCEP cert providers, e.g. `intune`, `jamf`, `byod` |
| `Enable` | `*bool` | Optional | Whether SCEP is enabled for this org |
| `Suspended` | `*bool` | Optional | Whether SCEP is suspended for this org<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingScep := models.OrgSettingScep{
        CertProviders:        []models.OrgSettingScepCertProviderEnum{
            models.OrgSettingScepCertProviderEnum_JAMF,
            models.OrgSettingScepCertProviderEnum_INTUNE,
            models.OrgSettingScepCertProviderEnum_BYOD,
        },
        Enable:               models.ToPointer(false),
        Suspended:            models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

