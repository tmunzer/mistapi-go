
# Search Wan Usage

Paginated response for WAN usage searches

## Structure

`SearchWanUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | Upper bound timestamp of the WAN usage search window, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of WAN usage records returned by this page |
| `Results` | [`[]models.WanUsages`](../../doc/models/wan-usages.md) | Optional | WAN usage records returned by a search response |
| `Start` | `*float64` | Optional | Lower bound timestamp of the WAN usage search window, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    searchWanUsage := models.SearchWanUsage{
        End:                  models.ToPointer(float64(247)),
        Limit:                models.ToPointer(46),
        Results:              []models.WanUsages{
            models.WanUsages{
                Mac:                  models.ToPointer("mac0"),
                PathType:             models.ToPointer("path_type8"),
                PathWeight:           models.ToPointer(242),
                PeerMac:              models.ToPointer("peer_mac6"),
                PeerPortId:           models.ToPointer("peer_port_id4"),
            },
            models.WanUsages{
                Mac:                  models.ToPointer("mac0"),
                PathType:             models.ToPointer("path_type8"),
                PathWeight:           models.ToPointer(242),
                PeerMac:              models.ToPointer("peer_mac6"),
                PeerPortId:           models.ToPointer("peer_port_id4"),
            },
        },
        Start:                models.ToPointer(float64(203.06)),
    }

}
```

