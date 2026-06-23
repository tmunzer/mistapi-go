
# Avprofile

Antivirus scanning profile with protocols, limits, and whitelist settings

*This model accepts additional fields of type interface{}.*

## Structure

`Avprofile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `FallbackAction` | [`*models.AvprofileFallbackActionEnum`](../../doc/models/avprofile-fallback-action-enum.md) | Optional | Action applied when antivirus scanning cannot complete. enum: `block`, `log-and-permit`, `permit` |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MaxFilesize` | `*int` | Optional | Maximum file size scanned by this antivirus profile, in KB<br><br>**Default**: `10000`<br><br>**Constraints**: `>= 20`, `<= 40000` |
| `MimeWhitelist` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the antivirus profile |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Protocols` | [`[]models.AvprofileProtocolEnum`](../../doc/models/avprofile-protocol-enum.md) | Optional | List of network protocols monitored by the antivirus profile. enum: `ftp`, `http`, `imap`, `pop3`, `smtp`<br><br>**Constraints**: *Minimum Items*: `1` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `UrlWhitelist` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    avprofile := models.Avprofile{
        CreatedTime:          models.ToPointer(float64(49.94)),
        FallbackAction:       models.ToPointer(models.AvprofileFallbackActionEnum_LOGANDPERMIT),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MaxFilesize:          models.ToPointer(10000),
        MimeWhitelist:        []string{
            "mime_whitelist3",
        },
        Name:                 "name4",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

