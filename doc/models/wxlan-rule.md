
# Wxlan Rule

WxLAN policy rule controlling traffic between WxLAN tags

*This model accepts additional fields of type interface{}.*

## Structure

`WxlanRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.WxlanRuleActionEnum`](../../doc/models/wxlan-rule-action-enum.md) | Optional | type of action, allow / block. enum: `allow`, `block` |
| `ApplyTags` | `[]string` | Optional | WxLAN tag identifiers applied when a rule matches |
| `BlockedApps` | `[]string` | Optional | Blocked apps (always blocking, ignoring action), the key of Get Application List |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DstAllowWxtags` | `[]string` | Optional | List of WxTag UUID to indicate these tags are allowed access |
| `DstDenyWxtags` | `[]string` | Optional | List of WxTag UUID to indicate these tags are blocked access |
| `DstWxtags` | `[]string` | Optional | List of WxTag UUID |
| `Enabled` | `*bool` | Optional | Whether this WxLAN rule is enabled<br><br>**Default**: `true` |
| `ForSite` | `*bool` | Optional, Read-only | Whether this WxLAN rule is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Order` | `int` | Required | Lookup priority for WxLAN rules; larger positive values match first, and -1 means LAST. Uniqueness is not checked<br><br>**Constraints**: `>= -1` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SrcWxtags` | `[]string` | Required | List of WxTag UUID to determine if this rule would match |
| `TemplateId` | `*uuid.UUID` | Optional | Only for Org Level WxRule |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    wxlanRule := models.WxlanRule{
        Action:               models.ToPointer(models.WxlanRuleActionEnum_ALLOW),
        ApplyTags:            []string{
            "apply_tags8",
            "apply_tags9",
            "apply_tags0",
        },
        BlockedApps:          []string{
            "mist",
            "all-videos",
        },
        CreatedTime:          models.ToPointer(float64(24.6)),
        DstAllowWxtags:       []string{
            "fff34466-eec0-3756-6765-381c728a6037",
            "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3",
        },
        DstDenyWxtags:        []string{
            "aaa34466-eec0-3756-6765-381c728a6037",
            "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3",
        },
        DstWxtags:            []string{
            "aaa34466-eec0-3756-6765-381c728a6037",
            "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3",
        },
        Enabled:              models.ToPointer(true),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Order:                1,
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        SrcWxtags:            []string{
            "8bfc2490-d726-3587-038d-cb2e71bd2330",
            "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9",
        },
        TemplateId:           models.ToPointer(uuid.MustParse("6aa54cbd-e039-4878-846a-04f270de8a5c")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

