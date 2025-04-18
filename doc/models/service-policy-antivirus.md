
# Service Policy Antivirus

For SRX-only

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePolicyAntivirus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvprofileId` | `*uuid.UUID` | Optional | org-level AV Profile can be used, this takes precedence over 'profile' |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Profile` | `*string` | Optional | Default / noftp / httponly / or keys from av_profiles |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "avprofile_id": "00000696-0000-0000-0000-000000000000",
  "profile": "profile0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

