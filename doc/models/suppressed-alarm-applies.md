
# Suppressed Alarm Applies

If `scope`==`site`. Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration`

## Structure

`SuppressedAlarmApplies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | UUID of the current org (if provided, the alarms will be suppressed at org level) |
| `SiteIds` | `[]uuid.UUID` | Optional | List of UUID of the sites within the org (if provided, the alarms will be suppressed for all the mentioned sites under the org) |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of UUID of the site groups within the org (if provided, the alarms will be suppressed for all the sites under those site groups) |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    suppressedAlarmApplies := models.SuppressedAlarmApplies{
        OrgId:                models.ToPointer(uuid.MustParse("0000048c-0000-0000-0000-000000000000")),
        SiteIds:              []uuid.UUID{
            uuid.MustParse("0000156a-0000-0000-0000-000000000000"),
            uuid.MustParse("00001569-0000-0000-0000-000000000000"),
            uuid.MustParse("00001568-0000-0000-0000-000000000000"),
        },
        SitegroupIds:         []uuid.UUID{
            uuid.MustParse("000015f4-0000-0000-0000-000000000000"),
            uuid.MustParse("000015f5-0000-0000-0000-000000000000"),
        },
    }

}
```

