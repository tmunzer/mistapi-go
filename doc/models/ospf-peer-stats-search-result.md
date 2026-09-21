
# Ospf Peer Stats Search Result

Paginated OSPF peer statistics search result

## Structure

`OspfPeerStatsSearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Query end timestamp for the OSPF peer statistics search |
| `Limit` | `*int` | Optional | Maximum number of OSPF peer statistics returned per page |
| `Next` | `*string` | Optional | URL for the next page of OSPF peer statistics, when available |
| `Results` | [`[]models.OspfPeerStatsSearchResultsItems`](../../doc/models/ospf-peer-stats-search-results-items.md) | Optional | List of OSPF peer statistic records |
| `Start` | `*int` | Optional | Query start timestamp for the OSPF peer statistics search |
| `Total` | `*int` | Optional | Number of OSPF peer statistic records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ospfPeerStatsSearchResult := models.OspfPeerStatsSearchResult{
        End:                  models.ToPointer(1711035686),
        Limit:                models.ToPointer(10),
        Next:                 models.ToPointer("next0"),
        Results:              []models.OspfPeerStatsSearchResultsItems{
            models.OspfPeerStatsSearchResultsItems{
                DeadTime:             models.ToPointer(140),
                Mac:                  models.ToPointer("mac0"),
                Neighbor:             models.ToPointer("neighbor6"),
                NeighborId:           models.ToPointer("neighbor_id0"),
            },
            models.OspfPeerStatsSearchResultsItems{
                DeadTime:             models.ToPointer(140),
                Mac:                  models.ToPointer("mac0"),
                Neighbor:             models.ToPointer("neighbor6"),
                NeighborId:           models.ToPointer("neighbor_id0"),
            },
            models.OspfPeerStatsSearchResultsItems{
                DeadTime:             models.ToPointer(140),
                Mac:                  models.ToPointer("mac0"),
                Neighbor:             models.ToPointer("neighbor6"),
                NeighborId:           models.ToPointer("neighbor_id0"),
            },
        },
        Start:                models.ToPointer(1710949286),
        Total:                models.ToPointer(232),
    }

}
```

