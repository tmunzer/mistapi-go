
# Response Psk Portal Logs Search

Paginated response for organization PSK Portal log search results

## Structure

`ResponsePskPortalLogsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp, in seconds, for the end of the PSK Portal log search window |
| `Limit` | `*int` | Optional | Maximum number of PSK Portal log entries returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of PSK Portal log search results |
| `Results` | [`[]models.ResponsePskPortalLogsSearchItem`](../../doc/models/response-psk-portal-logs-search-item.md) | Optional | PSK Portal log entries returned by search |
| `Start` | `*int` | Optional | Epoch timestamp, in seconds, for the start of the PSK Portal log search window |
| `Total` | `*int` | Optional | Number of PSK Portal log entries matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePskPortalLogsSearch := models.ResponsePskPortalLogsSearch{
        End:                  models.ToPointer(1428954000),
        Limit:                models.ToPointer(100),
        Next:                 models.ToPointer("next6"),
        Results:              []models.ResponsePskPortalLogsSearchItem{
            models.ResponsePskPortalLogsSearchItem{
                Id:                   models.ToPointer(uuid.MustParse("000023ba-0000-0000-0000-000000000000")),
                Message:              models.ToPointer("message6"),
                NameId:               models.ToPointer("name_id8"),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                PskId:                models.ToPointer(uuid.MustParse("00000b40-0000-0000-0000-000000000000")),
            },
        },
        Start:                models.ToPointer(1428939600),
        Total:                models.ToPointer(135),
    }

}
```

