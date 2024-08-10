
# Org Service Policy Antivirus

for SRX-only

## Structure

`OrgServicePolicyAntivirus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvprofileId` | `*uuid.UUID` | Optional | org-level AV Profile can be used, this takes precendence over 'profile' |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Profile` | `*string` | Optional | default / noftp / httponly / or keys from av_profiles |

## Example (as JSON)

```json
{
  "enabled": false,
  "avprofile_id": "00000594-0000-0000-0000-000000000000",
  "profile": "profile8"
}
```

