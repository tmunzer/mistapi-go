
# Search Wan Client

Paginated response for WAN client searches

## Structure

`SearchWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Upper bound timestamp of the WAN client search window, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of WAN client records returned by this page |
| `Next` | `*string` | Optional | URL for the next page of WAN client records, when more results are available |
| `Results` | [`[]models.StatsWanClient`](../../doc/models/stats-wan-client.md) | Optional | WAN client records returned by the search response |
| `Start` | `*int` | Optional | Lower bound timestamp of the WAN client search window, in epoch seconds |
| `Total` | `*int` | Optional | Count of WAN client records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    searchWanClient := models.SearchWanClient{
        End:                  models.ToPointer(92),
        Limit:                models.ToPointer(78),
        Next:                 models.ToPointer("next8"),
        Results:              []models.StatsWanClient{
            models.StatsWanClient{
                DhcpExpireTime:       models.ToPointer(float64(124.26)),
                DhcpStartTime:        models.ToPointer(float64(9.94)),
                Hostname:             []string{
                    "hostname6",
                    "hostname7",
                    "hostname8",
                },
                Ip:                   []string{
                    "ip7",
                    "ip8",
                },
                IpSrc:                models.ToPointer("ip_src6"),
            },
            models.StatsWanClient{
                DhcpExpireTime:       models.ToPointer(float64(124.26)),
                DhcpStartTime:        models.ToPointer(float64(9.94)),
                Hostname:             []string{
                    "hostname6",
                    "hostname7",
                    "hostname8",
                },
                Ip:                   []string{
                    "ip7",
                    "ip8",
                },
                IpSrc:                models.ToPointer("ip_src6"),
            },
            models.StatsWanClient{
                DhcpExpireTime:       models.ToPointer(float64(124.26)),
                DhcpStartTime:        models.ToPointer(float64(9.94)),
                Hostname:             []string{
                    "hostname6",
                    "hostname7",
                    "hostname8",
                },
                Ip:                   []string{
                    "ip7",
                    "ip8",
                },
                IpSrc:                models.ToPointer("ip_src6"),
            },
        },
        Start:                models.ToPointer(50),
    }

}
```

