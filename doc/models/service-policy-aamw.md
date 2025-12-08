
# Service Policy Aamw

SRX only

## Structure

`ServicePolicyAamw`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AamwprofileId` | `*uuid.UUID` | Optional | org-level Advanced Advance Anti Malware Profile (SkyAtp) Profile can be used, this takes precedence over 'profile' |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Profile` | [`*models.ServicePolicyAamwProfileEnum`](../../doc/models/service-policy-aamw-profile-enum.md) | Optional | enum: `docsonly`, `executables`, `standard`<br><br>**Default**: `"standard"` |

## Example (as JSON)

```json
{
  "enabled": false,
  "profile": "standard",
  "aamwprofile_id": "00001e70-0000-0000-0000-000000000000"
}
```

