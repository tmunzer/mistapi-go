
# Org Setting Auto Deviceprofile Assignment

Automatic device profile assignment configuration

## Structure

`OrgSettingAutoDeviceprofileAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | Whether automatic device profile assignment is enabled |
| `Rules` | [`models.Optional[[]models.OrgSettingAutoAssignmentRule]`](../../doc/models/org-setting-auto-assignment-rule.md) | Optional | Automatic device profile assignment rules, or null when automatic assignment is not configured |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingAutoDeviceprofileAssignment := models.OrgSettingAutoDeviceprofileAssignment{
        Enable:               models.ToPointer(false),
        Rules:                models.NewOptional(models.ToPointer([]models.OrgSettingAutoAssignmentRule{
            models.OrgSettingAutoAssignmentRule{
                CreateNewSiteIfNeeded: models.ToPointer(false),
                Expression:            models.NewOptional(models.ToPointer("expression4")),
                GatewaytemplateId:     models.ToPointer("gatewaytemplate_id0"),
                MatchCountry:          models.ToPointer("match_country8"),
                MatchDeviceType:       models.ToPointer(models.DeviceTypeDefaultApEnum_ENUMSWITCH),
                Src:                   models.OrgSettingAutoSiteAssignmentSrcEnum_NAME,
            },
        })),
    }

}
```

