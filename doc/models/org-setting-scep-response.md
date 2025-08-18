
# Org Setting Scep Response

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingScepResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CertProviders` | [`[]models.OrgSettingScepCertProviderEnum`](../../doc/models/org-setting-scep-cert-provider-enum.md) | Optional | List of SCEP cert providers, e.g. `intune`, `jamf`, `byod` |
| `Enabled` | `*bool` | Optional | - |
| `IntuneScepUrl` | `*string` | Optional | - |
| `JamfAccessToken` | `*string` | Optional | - |
| `JamfScepUrl` | `*string` | Optional | - |
| `JamfWebhookUrl` | `*string` | Optional | - |
| `Suspended` | `*bool` | Optional | Whether SCEP is suspended for this org<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "intune_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_access_token": "1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=",
  "jamf_scep_url": "https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep",
  "jamf_webhook_url": "https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep",
  "suspended": false,
  "cert_providers": [
    "jamf",
    "intune",
    "byod"
  ],
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

