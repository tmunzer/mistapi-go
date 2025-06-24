
# Idp Config

*This model accepts additional fields of type interface{}.*

## Structure

`IdpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlertOnly` | `*bool` | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `IdpprofileId` | `*uuid.UUID` | Optional | org_level IDP Profile can be used, this takes precedence over `profile` |
| `Profile` | `*string` | Optional | enum: `Custom`, `strict` (default), `standard` or keys from idp_profiles<br><br>**Default**: `"strict"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "idpprofile_id": "89b9d208-84a4-fa8f-af57-78f92c639cf2",
  "profile": "strict",
  "alert_only": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

