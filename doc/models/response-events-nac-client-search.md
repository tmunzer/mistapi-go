
# Response Events Nac Client Search

Paginated response for NAC client event search results

## Structure

`ResponseEventsNacClientSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp for the end of the NAC client event search window |
| `Limit` | `*int` | Optional | Maximum number of NAC client event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of NAC client event records |
| `Results` | [`[]models.NacClientEvent`](../../doc/models/nac-client-event.md) | Optional | List of NAC authentication events |
| `Start` | `*int` | Optional | Epoch timestamp for the start of the NAC client event search window |
| `Total` | `*int` | Optional | Number of NAC client event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseEventsNacClientSearch := models.ResponseEventsNacClientSearch{
        End:                  models.ToPointer(1513176951),
        Limit:                models.ToPointer(10),
        Next:                 models.ToPointer("next2"),
        Results:              []models.NacClientEvent{
            models.NacClientEvent{
                Ap:                   models.ToPointer("ap8"),
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTEAP),
                Bssid:                models.ToPointer("bssid0"),
                ClientType:           models.ToPointer(models.NacAccessTypeEnum_WIRED),
                DeviceMac:            models.ToPointer("device_mac0"),
            },
            models.NacClientEvent{
                Ap:                   models.ToPointer("ap8"),
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTEAP),
                Bssid:                models.ToPointer("bssid0"),
                ClientType:           models.ToPointer(models.NacAccessTypeEnum_WIRED),
                DeviceMac:            models.ToPointer("device_mac0"),
            },
        },
        Start:                models.ToPointer(1512572151),
        Total:                models.ToPointer(1),
    }

}
```

