
# Response Stats Assets

Paginated response for asset statistics search results

## Structure

`ResponseStatsAssets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the asset statistics search window |
| `Limit` | `int` | Required | Maximum number of asset statistics records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of asset statistics results |
| `Results` | [`[]models.StatsAsset`](../../doc/models/stats-asset.md) | Required | Asset statistics returned by the request |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the asset statistics search window |
| `Total` | `int` | Required | Number of asset statistics records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseStatsAssets := models.ResponseStatsAssets{
        End:                  154,
        Limit:                16,
        Next:                 models.ToPointer("next8"),
        Results:              []models.StatsAsset{
            models.StatsAsset{
                Ttl:                   models.ToPointer(162),
                BatteryPercent:        models.ToPointer(50),
                BatteryVoltage:        models.ToPointer(float64(2970)),
                Beam:                  models.ToPointer(6),
                By:                    models.ToPointer("asset"),
                DeviceId:              models.ToPointer(uuid.MustParse("00000000-0000-0000-1000-5c5b35000001")),
                DeviceName:            models.ToPointer("a"),
                Duration:              models.ToPointer(120),
                EddystoneUidInstance:  models.ToPointer("5c5b35000001"),
                EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
                EddystoneUrlUrl:       models.ToPointer("https://www.abc.com"),
                IbeaconMajor:          models.NewOptional(models.ToPointer(1234)),
                IbeaconMinor:          models.NewOptional(models.ToPointer(1234)),
                IbeaconUuid:           models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
                Id:                    models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                LastSeen:              models.NewOptional(models.ToPointer(float64(1470417522))),
                Mac:                   "6fa474be7ae5",
                MapId:                 models.ToPointer(uuid.MustParse("c45be59f-854d-4ef7-b782-dcd6309c84a9")),
                MfgCompanyId:          models.ToPointer(935),
                MfgData:               models.ToPointer("648520a1020000"),
                Name:                  models.ToPointer("6fa474be7ae5"),
                Rssi:                  models.ToPointer(-60),
                Temperature:           models.ToPointer(float64(23)),
                X:                     models.ToPointer(float64(280.19918140310193)),
                Y:                     models.ToPointer(float64(420.2987721046529)),
            },
        },
        Start:                112,
        Total:                146,
    }

}
```

