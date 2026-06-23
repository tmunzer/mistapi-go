
# Org Setting Scep Response

Read-only Mist SCEP settings returned for the organization

## Structure

`OrgSettingScepResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CertProviders` | [`[]models.OrgSettingScepCertProviderEnum`](../../doc/models/org-setting-scep-cert-provider-enum.md) | Optional | List of SCEP cert providers, e.g. `intune`, `jamf`, `byod` |
| `Enabled` | `*bool` | Optional, Read-only | Whether Mist SCEP is enabled for this organization |
| `IntuneScepUrl` | `*string` | Optional, Read-only | Intune SCEP enrollment URL for this organization |
| `JamfAccessToken` | `*string` | Optional, Read-only | Access token used by Jamf to call Mist SCEP |
| `JamfScepUrl` | `*string` | Optional, Read-only | Jamf SCEP enrollment URL for this organization |
| `JamfWebhookUrl` | `*string` | Optional, Read-only | Jamf webhook URL for SCEP status callbacks |
| `Suspended` | `*bool` | Optional | Whether SCEP is suspended for this org<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingScepResponse := models.OrgSettingScepResponse{
        CertProviders:        []models.OrgSettingScepCertProviderEnum{
            models.OrgSettingScepCertProviderEnum_INTUNE,
        },
        Enabled:              models.ToPointer(false),
        IntuneScepUrl:        models.ToPointer("https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep"),
        JamfAccessToken:      models.ToPointer("1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81="),
        JamfScepUrl:          models.ToPointer("https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep"),
        JamfWebhookUrl:       models.ToPointer("https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"),
        Suspended:            models.ToPointer(false),
    }

}
```

