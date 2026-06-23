
# Template Applies

Where this template should be applied to, can be org_id, site_ids, sitegroup_ids

## Structure

`TemplateApplies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | Organization included in the WLAN template application scope |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    templateApplies := models.TemplateApplies{
        OrgId:                models.ToPointer(uuid.MustParse("00001ce6-0000-0000-0000-000000000000")),
        SiteIds:              []uuid.UUID{
            uuid.MustParse("00001114-0000-0000-0000-000000000000"),
            uuid.MustParse("00001113-0000-0000-0000-000000000000"),
            uuid.MustParse("00001112-0000-0000-0000-000000000000"),
        },
        SitegroupIds:         []uuid.UUID{
            uuid.MustParse("0000073e-0000-0000-0000-000000000000"),
        },
    }

}
```

