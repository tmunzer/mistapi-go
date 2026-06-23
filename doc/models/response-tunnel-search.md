
# Response Tunnel Search

Paginated response for organization tunnel statistics search results

## Structure

`ResponseTunnelSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the tunnel statistics search window |
| `Limit` | `int` | Required | Maximum number of tunnel statistics records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of tunnel statistics results |
| `Results` | [`[]models.ResponseTunnelSearchItem`](../../doc/models/containers/response-tunnel-search-item.md) | Required | Tunnel statistics record; shape depends on the requested tunnel type |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the tunnel statistics search window |
| `Total` | `int` | Required | Number of tunnel statistics records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseTunnelSearch := models.ResponseTunnelSearch{
        End:                  48,
        Limit:                122,
        Next:                 models.ToPointer("next8"),
        Results:              []models.ResponseTunnelSearchItem{
            models.ResponseTunnelSearchItemContainer.FromStatsMxtunnel(models.StatsMxtunnel{
                Ap:                   models.ToPointer("ap4"),
                ForSite:              models.ToPointer(false),
                Fwupdate:             models.ToPointer(models.FwupdateStat{
                    Progress:             models.NewOptional(models.ToPointer(100)),
                    Status:               models.NewOptional(models.ToPointer(models.FwupdateStatStatusEnum_INPROGRESS)),
                    StatusId:             models.NewOptional(models.ToPointer(70)),
                    Timestamp:            models.ToPointer(float64(147.68)),
                    WillRetry:            models.NewOptional(models.ToPointer(false)),
                }),
                LastSeen:             models.NewOptional(models.ToPointer(float64(249.6))),
                Mtu:                  models.ToPointer(34),
                RemoteIp:             "remote_ip4",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            }),
            models.ResponseTunnelSearchItemContainer.FromStatsMxtunnel(models.StatsMxtunnel{
                Ap:                   models.ToPointer("ap4"),
                ForSite:              models.ToPointer(false),
                Fwupdate:             models.ToPointer(models.FwupdateStat{
                    Progress:             models.NewOptional(models.ToPointer(100)),
                    Status:               models.NewOptional(models.ToPointer(models.FwupdateStatStatusEnum_INPROGRESS)),
                    StatusId:             models.NewOptional(models.ToPointer(70)),
                    Timestamp:            models.ToPointer(float64(147.68)),
                    WillRetry:            models.NewOptional(models.ToPointer(false)),
                }),
                LastSeen:             models.NewOptional(models.ToPointer(float64(249.6))),
                Mtu:                  models.ToPointer(34),
                RemoteIp:             "remote_ip4",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            }),
            models.ResponseTunnelSearchItemContainer.FromStatsMxtunnel(models.StatsMxtunnel{
                Ap:                   models.ToPointer("ap4"),
                ForSite:              models.ToPointer(false),
                Fwupdate:             models.ToPointer(models.FwupdateStat{
                    Progress:             models.NewOptional(models.ToPointer(100)),
                    Status:               models.NewOptional(models.ToPointer(models.FwupdateStatStatusEnum_INPROGRESS)),
                    StatusId:             models.NewOptional(models.ToPointer(70)),
                    Timestamp:            models.ToPointer(float64(147.68)),
                    WillRetry:            models.NewOptional(models.ToPointer(false)),
                }),
                LastSeen:             models.NewOptional(models.ToPointer(float64(249.6))),
                Mtu:                  models.ToPointer(34),
                RemoteIp:             "remote_ip4",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            }),
        },
        Start:                6,
        Total:                216,
    }

}
```

