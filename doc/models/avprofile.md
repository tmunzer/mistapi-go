
# Avprofile

*This model accepts additional fields of type interface{}.*

## Structure

`Avprofile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `FallbackAction` | [`*models.AvprofileFallbackActionEnum`](../../doc/models/avprofile-fallback-action-enum.md) | Optional | enum: `block`, `permit` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `MaxFilesize` | `*int` | Optional | In KB<br>**Default**: `10000`<br>**Constraints**: `>= 20`, `<= 40000` |
| `MimeWhitelist` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Protocols` | [`[]models.AvprofileProtocolsEnum`](../../doc/models/avprofile-protocols-enum.md) | Optional | List of protocols to monitor. enum: `ftp`, `http`, `imap`, `pop3`, `smtp`<br>**Constraints**: *Minimum Items*: `1` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `UrlWhitelist` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "max_filesize": 10000,
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 49.94,
  "fallback_action": "block",
  "mime_whitelist": [
    "mime_whitelist3"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

