
# Avprofile

## Structure

`Avprofile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `FallbackAction` | [`*models.AvprofileFallbackActionEnum`](../../doc/models/avprofile-fallback-action-enum.md) | Optional | enum: `block`, `permit` |
| `Id` | `*uuid.UUID` | Optional | - |
| `MexFilesize` | `*int` | Optional | in KB<br>**Constraints**: `>= 20`, `<= 40000` |
| `MimeWhitelist` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Protocols` | [`[]models.AvprofileProtocolsEnum`](../../doc/models/avprofile-protocols-enum.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `UrlWhitelist` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 49.94,
  "fallback_action": "block",
  "id": "00001a6c-0000-0000-0000-000000000000",
  "mex_filesize": 94,
  "mime_whitelist": [
    "mime_whitelist3"
  ]
}
```

