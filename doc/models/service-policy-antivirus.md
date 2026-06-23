
# Service Policy Antivirus

SRX antivirus inspection settings for a service policy

## Structure

`ServicePolicyAntivirus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvprofileId` | `*uuid.UUID` | Optional | Organization-level antivirus profile ID; takes precedence over inline `profile` settings |
| `Enabled` | `*bool` | Optional | Whether antivirus inspection is enabled for the service policy<br><br>**Default**: `false` |
| `Profile` | `*string` | Optional | Antivirus profile name to apply, such as `default`, `noftp`, `httponly`, or an AV profile key |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    servicePolicyAntivirus := models.ServicePolicyAntivirus{
        AvprofileId:          models.ToPointer(uuid.MustParse("000024e0-0000-0000-0000-000000000000")),
        Enabled:              models.ToPointer(false),
        Profile:              models.ToPointer("profile4"),
    }

}
```

