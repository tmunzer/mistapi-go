
# Idp Config

Intrusion detection and prevention settings for a service policy

## Structure

`IdpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlertOnly` | `*bool` | Optional | Whether to alert without enforcing IDP prevention actions |
| `Enabled` | `*bool` | Optional | Whether IDP inspection is enabled for the policy<br><br>**Default**: `false` |
| `IdpprofileId` | `*uuid.UUID` | Optional | org_level IDP Profile can be used, this takes precedence over `profile` |
| `Profile` | `*string` | Optional | enum: `Custom`, `strict` (default), `standard` or keys from idp_profiles<br><br>**Default**: `"strict"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    idpConfig := models.IdpConfig{
        AlertOnly:            models.ToPointer(false),
        Enabled:              models.ToPointer(false),
        IdpprofileId:         models.ToPointer(uuid.MustParse("89b9d208-84a4-fa8f-af57-78f92c639cf2")),
        Profile:              models.ToPointer("strict"),
    }

}
```

