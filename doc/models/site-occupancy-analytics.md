
# Site Occupancy Analytics

Analytics settings for site occupancy

## Structure

`SiteOccupancyAnalytics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsEnabled` | `*bool` | Optional | Indicate whether named BLE assets should be included in the zone occupancy calculation<br><br>**Default**: `false` |
| `ClientsEnabled` | `*bool` | Optional | Indicate whether connected Wi-Fi clients should be included in the zone occupancy calculation<br><br>**Default**: `true` |
| `MinDuration` | `*int` | Optional | Minimum dwell duration before a client or asset is counted in occupancy analytics<br><br>**Default**: `3000` |
| `SdkclientsEnabled` | `*bool` | Optional | Indicate whether SDK clients should be included in the zone occupancy calculation<br><br>**Default**: `false` |
| `UnconnectedClientsEnabled` | `*bool` | Optional | Indicate whether unconnected Wi-Fi clients should be included in the zone occupancy calculation<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteOccupancyAnalytics := models.SiteOccupancyAnalytics{
        AssetsEnabled:             models.ToPointer(false),
        ClientsEnabled:            models.ToPointer(true),
        MinDuration:               models.ToPointer(3000),
        SdkclientsEnabled:         models.ToPointer(false),
        UnconnectedClientsEnabled: models.ToPointer(false),
    }

}
```

