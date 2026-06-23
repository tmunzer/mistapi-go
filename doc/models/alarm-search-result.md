
# Alarm Search Result

Paginated response returned by an alarm search

## Structure

`AlarmSearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Component` | `*string` | Optional | Alarm component matched by the search result |
| `End` | `int` | Required | Upper bound of the alarm search window, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of alarm results returned |
| `Next` | `*string` | Optional | URL for the next page of alarm search results |
| `Page` | `*int` | Optional | Current page number for paginated alarm search results |
| `Results` | [`[]models.Alarm`](../../doc/models/alarm.md) | Required | Alarm records returned by an alarm search response |
| `Start` | `int` | Required | Lower bound of the alarm search window, in epoch seconds |
| `Total` | `int` | Required | Number of alarm records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    alarmSearchResult := models.AlarmSearchResult{
        Component:            models.ToPointer("component6"),
        End:                  1711035686,
        Limit:                10,
        Next:                 models.ToPointer("/api/v1/orgs/b3b9f5e6-67b1-4112-9b4c-6824c565eaeb/alarms/search?end=1711035686&limit=10&search_after=%5B1711031354000%2C+%2256bfa7af-b2db-43ee-a4c8-9b820bbba0e1%22%5D&start=1710949286"),
        Page:                 models.ToPointer(1),
        Results:              []models.Alarm{
            models.Alarm{
                AckAdminId:           models.ToPointer(uuid.MustParse("456b7016-a916-a4b1-78dd-72b947c152b7")),
                AckAdminName:         models.ToPointer("Joe"),
                Acked:                models.ToPointer(true),
                AckedTime:            models.ToPointer(1711031352),
                Aps:                  []string{
                    "ffeeddccbbaa",
                    "ffeeddccbbab",
                },
                Count:                2,
                Gateways:             []string{
                    "ffeeddccbbaa",
                    "ffeeddccbbab",
                },
                Group:                "security",
                Hostnames:            []string{
                    "MC_DavidL",
                    "MCM_AP_33_Nishant",
                },
                Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                LastSeen:             float64(1711031774),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                Severity:             "critical",
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Switches:             []string{
                    "ffeeddccbbaa",
                    "ffeeddccbbab",
                },
                Timestamp:            float64(2.64),
                Type:                 "rogue_client",
            },
        },
        Start:                1710949286,
        Total:                232,
    }

}
```

