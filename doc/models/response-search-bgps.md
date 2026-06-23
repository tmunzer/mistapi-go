
# Response Search Bgps

Paginated response for BGP peer statistics search results

## Structure

`ResponseSearchBgps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | Epoch timestamp, in seconds, for the end of the BGP statistics search window |
| `Limit` | `*int` | Optional | Maximum number of BGP peer statistics records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of BGP peer statistics results |
| `Results` | [`[]models.BgpStats`](../../doc/models/bgp-stats.md) | Optional | BGP statistics records returned by a search response |
| `Start` | `*float64` | Optional | Epoch timestamp, in seconds, for the start of the BGP statistics search window |
| `Total` | `*int` | Optional | Number of BGP peer statistics records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSearchBgps := models.ResponseSearchBgps{
        End:                  models.ToPointer(float64(72.86)),
        Limit:                models.ToPointer(52),
        Next:                 models.ToPointer("next6"),
        Results:              []models.BgpStats{
            models.BgpStats{
                EvpnOverlay:          models.ToPointer(false),
                ForOverlay:           models.ToPointer(false),
                LocalAs:              models.ToPointer(models.BgpAsContainer.FromString("String3")),
                Mac:                  models.ToPointer("mac0"),
                Model:                models.ToPointer("model4"),
            },
        },
        Start:                models.ToPointer(float64(28.92)),
    }

}
```

