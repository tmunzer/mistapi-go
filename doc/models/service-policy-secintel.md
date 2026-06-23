
# Service Policy Secintel

SRX SecIntel settings for a service policy

## Structure

`ServicePolicySecintel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether SecIntel inspection is enabled for the service policy<br><br>**Default**: `false` |
| `Profile` | [`*models.ServicePolicySecintelProfileEnum`](../../doc/models/service-policy-secintel-profile-enum.md) | Optional | enum: `default`, `standard`, `strict`<br><br>**Default**: `"default"` |
| `SecintelprofileId` | `*string` | Optional | Organization-level SecIntel profile ID; takes precedence over inline `profile` settings |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySecintel := models.ServicePolicySecintel{
        Enabled:              models.ToPointer(false),
        Profile:              models.ToPointer(models.ServicePolicySecintelProfileEnum_ENUMDEFAULT),
        SecintelprofileId:    models.ToPointer("secintelprofile_id6"),
    }

}
```

