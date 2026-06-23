
# Org Site Wan Wifi

Paginated WAN SLE results for organization sites

## Structure

`OrgSiteWanWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Query end timestamp for the returned SLE window |
| `Interval` | `int` | Required | Aggregation interval, in seconds, used for the SLE query |
| `Limit` | `int` | Required | Maximum number of site results returned per page |
| `Page` | `int` | Required | Result page number returned by the query |
| `Results` | [`[]models.OrgSiteSleWanResult`](../../doc/models/org-site-sle-wan-result.md) | Required | WAN SLE results returned for organization sites<br><br>**Constraints**: *Unique Items Required* |
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
    orgSiteWanWifi := models.OrgSiteWanWifi{
        End:                  float64(205.68),
        Interval:             254,
        Limit:                82,
        Page:                 196,
        Results:              []models.OrgSiteSleWanResult{
            models.OrgSiteSleWanResult{
                ApplicationHealth:    models.ToPointer(float64(246.28)),
                GatewayHealth:        float64(241.74),
                NumClients:           models.ToPointer(float64(68.96)),
                NumGateways:          models.ToPointer(float64(213.6)),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                WanLinkHealth:        models.ToPointer(float64(134.82)),
            },
        },
        Start:                float64(161.74),
        Total:                80,
    }

}
```

