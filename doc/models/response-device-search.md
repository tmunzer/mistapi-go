
# Response Device Search

Paginated response for organization or site device search results

## Structure

`ResponseDeviceSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the device search window |
| `Limit` | `int` | Required | Maximum number of device records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of device records |
| `Results` | [`[]models.ResponseDeviceSearchResultsItems`](../../doc/models/containers/response-device-search-results-items.md) | Required | Device search record for an AP, switch, or gateway<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the device search window |
| `Total` | `int` | Required | Number of device records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseDeviceSearch := models.ResponseDeviceSearch{
        End:                  138,
        Limit:                32,
        Next:                 models.ToPointer("next8"),
        Results:              []models.ResponseDeviceSearchResultsItems{
            models.ResponseDeviceSearchResultsItemsContainer.FromApSearch(models.ApSearch{
                Band24Bandwidth:      models.ToPointer("band_24_bandwidth2"),
                Band24Channel:        models.ToPointer(200),
                Band24Power:          models.ToPointer(154),
                Band5Bandwidth:       models.ToPointer("band_5_bandwidth0"),
                Band5Channel:         models.ToPointer(132),
                MxtunnelStatus:       "mxtunnel_status4",
                PowerConstrained:     false,
                PowerOpmode:          "power_opmode0",
                Wlans:                []models.ApSearchWlan{
                    models.ApSearchWlan{
                        Id:                   models.ToPointer(uuid.MustParse("00001c56-0000-0000-0000-000000000000")),
                        Ssid:                 models.ToPointer("ssid8"),
                    },
                    models.ApSearchWlan{
                        Id:                   models.ToPointer(uuid.MustParse("00001c56-0000-0000-0000-000000000000")),
                        Ssid:                 models.ToPointer("ssid8"),
                    },
                    models.ApSearchWlan{
                        Id:                   models.ToPointer(uuid.MustParse("00001c56-0000-0000-0000-000000000000")),
                        Ssid:                 models.ToPointer("ssid8"),
                    },
                },
            }),
        },
        Start:                96,
        Total:                126,
    }

}
```

