
# Org Setting Mist Nac Idp

Mist NAC identity provider realm mapping

## Structure

`OrgSettingMistNacIdp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExcludeRealms` | `[]string` | Optional | When the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `UserRealms` | `[]string` | Optional | Which realm should trigger this IDP. User Realm is extracted from:<br><br>* Username-AVP (`mist.com` from john@mist.com)<br>* Cert CN |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSettingMistNacIdp := models.OrgSettingMistNacIdp{
        ExcludeRealms:        []string{
            "exclude_realms3",
        },
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        UserRealms:           []string{
            "user_realms0",
            "user_realms1",
        },
    }

}
```

