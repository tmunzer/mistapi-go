
# Nac Rule

NAC authentication policy rule that matches request criteria and applies an allow or block action

*This model accepts additional fields of type interface{}.*

## Structure

`NacRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`models.NacRuleActionEnum`](../../doc/models/nac-rule-action-enum.md) | Required | Allow or block decision applied when the NAC rule matches. enum: `allow`, `block` |
| `ApplyTags` | `[]string` | Optional | All optional, this goes into Access-Accept |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DryRun` | `*bool` | Optional | Whether the NAC rule is in dry-run mode, where matches are logged but the action is not enforced |
| `Enabled` | `*bool` | Optional | Whether the NAC rule is evaluated during policy matching<br><br>**Default**: `true` |
| `GuestAuthState` | [`*models.NacRuleGuestAuthStateEnum`](../../doc/models/nac-rule-guest-auth-state-enum.md) | Optional | Guest portal authorization state. enum: `authorized`, `unknown` |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Matching` | [`*models.NacRuleMatching`](../../doc/models/nac-rule-matching.md) | Optional | Criteria used to include or exclude a NAC authentication request from a rule |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Human-readable name of the NAC rule |
| `NotMatching` | [`*models.NacRuleMatching`](../../doc/models/nac-rule-matching.md) | Optional | Criteria used to include or exclude a NAC authentication request from a rule |
| `Order` | `*int` | Optional | Rule priority; lower values are evaluated with higher priority<br><br>**Constraints**: `>= 0` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    nacRule := models.NacRule{
        Action:               models.NacRuleActionEnum_ALLOW,
        ApplyTags:            []string{
            "c049dfcd-0c73-5014-1c64-062e9903f1e5",
        },
        CreatedTime:          models.ToPointer(float64(2.12)),
        DryRun:               models.ToPointer(false),
        Enabled:              models.ToPointer(true),
        GuestAuthState:       models.ToPointer(models.NacRuleGuestAuthStateEnum_AUTHORIZED),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 "name2",
        Order:                models.ToPointer(1),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

