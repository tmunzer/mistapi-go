
# Org Site Sle Wifi

Paginated Wi-Fi SLE results for organization sites

## Structure

`OrgSiteSleWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Query end timestamp for the returned SLE window |
| `Interval` | `int` | Required | Aggregation interval, in seconds, used for the SLE query |
| `Limit` | `int` | Required | Maximum number of site results returned per page |
| `Page` | `int` | Required | Result page number returned by the query |
| `Results` | [`[]models.OrgSiteSleWifiResult`](../../doc/models/org-site-sle-wifi-result.md) | Required | Wi-Fi SLE results returned for organization sites<br><br>**Constraints**: *Unique Items Required* |
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
    orgSiteSleWifi := models.OrgSiteSleWifi{
        End:                  float64(4.52),
        Interval:             106,
        Limit:                230,
        Page:                 88,
        Results:              []models.OrgSiteSleWifiResult{
            models.OrgSiteSleWifiResult{
                ApAvailability:       float64(204.4),
                ApHealth:             models.ToPointer(float64(236.64)),
                Capacity:             models.ToPointer(float64(106.72)),
                Coverage:             models.ToPointer(float64(128.26)),
                NumAps:               models.ToPointer(float64(247.16)),
                NumClients:           models.ToPointer(float64(68.96)),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
            },
        },
        Start:                float64(216.58),
        Total:                68,
    }

}
```

