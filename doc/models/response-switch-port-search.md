
# Response Switch Port Search

Paginated response for switch and gateway port statistics search results

## Structure

`ResponseSwitchPortSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp, in seconds, for the end of the port statistics reporting window |
| `Limit` | `int` | Required | Maximum number of port statistics records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of switch and gateway port statistics results |
| `Results` | [`[]models.StatsSwitchPort`](../../doc/models/stats-switch-port.md) | Required | Switch and gateway port statistics records returned by search |
| `Start` | `*int` | Optional | Epoch timestamp, in seconds, for the start of the port statistics reporting window |
| `Total` | `int` | Required | Number of port statistics records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSwitchPortSearch := models.ResponseSwitchPortSearch{
        End:                  models.ToPointer(1513177200),
        Limit:                10,
        Next:                 models.ToPointer("next4"),
        Results:              []models.StatsSwitchPort{
            models.StatsSwitchPort{
                Active:               models.ToPointer(false),
                AuthState:            models.ToPointer(models.PortAuthStateEnum_AUTHENTICATED),
                Disabled:             models.ToPointer(false),
                ForSite:              models.ToPointer(false),
                FullDuplex:           models.ToPointer(true),
                Mac:                  "5c4527a96580",
                NeighborMac:          models.ToPointer("64d814353400"),
                NeighborPortDesc:     models.ToPointer("GigabitEthernet1/0/21"),
                NeighborSystemName:   models.ToPointer("CORP-D-SW-2"),
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                PortId:               "ge-0/0/0",
                PortMac:              models.ToPointer("5c4527a96580"),
                PortUsage:            models.ToPointer("lan"),
                RxBps:                models.NewOptional(models.ToPointer(int64(60003))),
                RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
                RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Speed:                models.ToPointer(1000),
                TxBps:                models.NewOptional(models.ToPointer(int64(634301))),
                TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
                TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
                Type:                 models.ToPointer(models.StatsSwitchPortTypeEnum_GATEWAY),
                XcvrModel:            models.ToPointer("SFP+-10G-SR"),
                XcvrPartNumber:       models.ToPointer("740-021487"),
                XcvrSerial:           models.ToPointer("N6AA9HT"),
            },
        },
        Start:                models.ToPointer(1511967600),
        Total:                100,
    }

}
```

