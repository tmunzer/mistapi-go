
# Org Setting Scep

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

## Example (as JSON)

```json
{
  "suspended": false,
  "cert_providers": [
    "intune"
  ],
  "enable": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

