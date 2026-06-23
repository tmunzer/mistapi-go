
# Response Events Devices

Paginated response for site device event search results

## Structure

`ResponseEventsDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the device event search window |
| `Limit` | `int` | Required | Maximum number of device event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of device event records |
| `Results` | [`[]models.DeviceEvent`](../../doc/models/device-event.md) | Required | List of device event payloads<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the device event search window |
| `Total` | `int` | Required | Number of device event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseEventsDevices := models.ResponseEventsDevices{
        End:                  30,
        Limit:                140,
        Next:                 models.ToPointer("next8"),
        Results:              []models.DeviceEvent{
            models.DeviceEvent{
                Ap:                   models.ToPointer("ap8"),
                ApName:               models.ToPointer("ap_name6"),
                Apfw:                 models.ToPointer("apfw8"),
                AuditId:              models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                Bandwidth:            models.ToPointer(198),
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Timestamp:            float64(2.64),
                Type:                 "type4",
            },
        },
        Start:                244,
        Total:                22,
    }

}
```

