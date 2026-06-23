
# Search Wireless Client Session

Paginated response for wireless client session searches

## Structure

`SearchWirelessClientSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Upper bound timestamp of the wireless client session search window, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of wireless client session records returned by this page |
| `Next` | `*string` | Optional | URL for the next page of wireless client session records, when more results are available |
| `Results` | [`[]models.WirelessClientSession`](../../doc/models/wireless-client-session.md) | Required | Wireless client session records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | Lower bound timestamp of the wireless client session search window, in epoch seconds |
| `Total` | `int` | Required | Count of wireless client session records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    searchWirelessClientSession := models.SearchWirelessClientSession{
        End:                  float64(4.8),
        Limit:                202,
        Next:                 models.ToPointer("next0"),
        Results:              []models.WirelessClientSession{
            models.WirelessClientSession{
                Ap:                   "ap8",
                Band:                 "band8",
                ClientManufacture:    models.NewOptional(models.ToPointer("client_manufacture0")),
                Connect:              198,
                Disconnect:           148,
                Duration:             float64(104.42),
                ForSite:              models.ToPointer(false),
                Mac:                  "mac0",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Ssid:                 "ssid6",
                Tags:                 []string{
                    "tags1",
                    "tags2",
                },
                Timestamp:            float64(2.64),
                WlanId:               uuid.MustParse("00000742-0000-0000-0000-000000000000"),
            },
        },
        Start:                float64(216.86),
        Total:                40,
    }

}
```

