
# Service Policy Aamw

SRX advanced anti-malware settings for a service policy

## Structure

`ServicePolicyAamw`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AamwprofileId` | `*uuid.UUID` | Optional | Organization-level advanced anti-malware profile ID; takes precedence over inline `profile` settings |
| `Enabled` | `*bool` | Optional | Whether advanced anti-malware inspection is enabled for the service policy<br><br>**Default**: `false` |
| `Profile` | [`*models.ServicePolicyAamwProfileEnum`](../../doc/models/service-policy-aamw-profile-enum.md) | Optional | enum: `docsonly`, `executables`, `standard`<br><br>**Default**: `"standard"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    servicePolicyAamw := models.ServicePolicyAamw{
        AamwprofileId:        models.ToPointer(uuid.MustParse("0000180a-0000-0000-0000-000000000000")),
        Enabled:              models.ToPointer(false),
        Profile:              models.ToPointer(models.ServicePolicyAamwProfileEnum_STANDARD),
    }

}
```

