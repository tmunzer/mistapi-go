
# Response Site Search

Paginated response for organization site search results

## Structure

`ResponseSiteSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the site search window |
| `Limit` | `int` | Required | Maximum number of site records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of site search results |
| `Results` | [`[]models.ResponseSiteSearchItem`](../../doc/models/response-site-search-item.md) | Required | Site records returned by organization site search<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the site search window |
| `Total` | `int` | Required | Number of site records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSiteSearch := models.ResponseSiteSearch{
        End:                  110,
        Limit:                60,
        Next:                 models.ToPointer("next2"),
        Results:              []models.ResponseSiteSearchItem{
            models.ResponseSiteSearchItem{
                AutoUpgradeEnabled:   false,
                AutoUpgradeVersion:   "auto_upgrade_version4",
                CountryCode:          models.NewOptional(models.ToPointer("country_code6")),
                HoneypotEnabled:      false,
                Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                Name:                 "name6",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Timestamp:            float64(2.64),
                Timezone:             "timezone4",
                VnaEnabled:           false,
                WifiEnabled:          false,
            },
        },
        Start:                68,
        Total:                154,
    }

}
```

