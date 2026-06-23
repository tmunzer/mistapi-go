
# Stats Marvis Clients Search

Paginated list of Marvis Client stats records

## Structure

`StatsMarvisClientsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `*int` | Optional | Maximum number of results requested |
| `Results` | [`[]models.StatsMarvisClient`](../../doc/models/stats-marvis-client.md) | Optional | List of Marvis Client stats records |
| `Total` | `*int` | Optional | Total number of matching results |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMarvisClientsSearch := models.StatsMarvisClientsSearch{
        Limit:                models.ToPointer(250),
        Results:              []models.StatsMarvisClient{
            models.StatsMarvisClient{
                BatteryCharging:      models.ToPointer(false),
                BatteryLevel:         models.ToPointer(212),
                CpuBackground:        models.ToPointer(float64(143.26)),
                CpuIdle:              models.ToPointer(float64(227.04)),
                CpuSystem:            models.ToPointer(float64(190.88)),
            },
            models.StatsMarvisClient{
                BatteryCharging:      models.ToPointer(false),
                BatteryLevel:         models.ToPointer(212),
                CpuBackground:        models.ToPointer(float64(143.26)),
                CpuIdle:              models.ToPointer(float64(227.04)),
                CpuSystem:            models.ToPointer(float64(190.88)),
            },
            models.StatsMarvisClient{
                BatteryCharging:      models.ToPointer(false),
                BatteryLevel:         models.ToPointer(212),
                CpuBackground:        models.ToPointer(float64(143.26)),
                CpuIdle:              models.ToPointer(float64(227.04)),
                CpuSystem:            models.ToPointer(float64(190.88)),
            },
        },
        Total:                models.ToPointer(88),
    }

}
```

