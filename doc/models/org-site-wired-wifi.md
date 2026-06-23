
# Org Site Wired Wifi

Paginated wired SLE results for organization sites

## Structure

`OrgSiteWiredWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Query end timestamp for the returned SLE window |
| `Interval` | `int` | Required | Aggregation interval, in seconds, used for the SLE query |
| `Limit` | `int` | Required | Maximum number of site results returned per page |
| `Page` | `int` | Required | Result page number returned by the query |
| `Results` | [`[]models.OrgSiteSleWiredResult`](../../doc/models/org-site-sle-wired-result.md) | Required | Wired SLE results returned for organization sites<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | Query start timestamp for the returned SLE window |
| `Total` | `int` | Required | Number of matching site results available for the query |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSiteWiredWifi := models.OrgSiteWiredWifi{
        End:                  float64(84.3),
        Interval:             108,
        Limit:                188,
        Page:                 46,
        Results:              []models.OrgSiteSleWiredResult{
            models.OrgSiteSleWiredResult{
                NumClients:           models.ToPointer(float64(68.96)),
                NumSwitches:          models.ToPointer(float64(67.86)),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                SwitchBandwidth:      models.ToPointer(float64(71.16)),
                SwitchHealth:         float64(218.74),
                SwitchThroughput:     models.ToPointer(float64(157.64)),
            },
        },
        Start:                float64(40.36),
        Total:                26,
    }

}
```

