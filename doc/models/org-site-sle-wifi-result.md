
# Org Site Sle Wifi Result

Wi-Fi SLE scores and counts for one site

## Structure

`OrgSiteSleWifiResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAvailability` | `float64` | Required | AP availability SLE score for this site |
| `ApHealth` | `*float64` | Optional | AP health SLE score for this site |
| `Capacity` | `*float64` | Optional | Wi-Fi capacity SLE score for this site |
| `Coverage` | `*float64` | Optional | Wi-Fi coverage SLE score for this site |
| `NumAps` | `*float64` | Optional | Number of APs included in this site result |
| `NumClients` | `*float64` | Optional | Number of Wi-Fi clients included in this site result |
| `Roaming` | `*float64` | Optional | Wi-Fi roaming SLE score for this site |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `SuccessfulConnect` | `*float64` | Optional | Successful connection SLE score for this site |
| `Throughput` | `*float64` | Optional | Wi-Fi throughput SLE score for this site |
| `TimeToConnect` | `*float64` | Optional | Client connection-time SLE score for this site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSiteSleWifiResult := models.OrgSiteSleWifiResult{
        ApAvailability:       float64(214.74),
        ApHealth:             models.ToPointer(float64(246.98)),
        Capacity:             models.ToPointer(float64(159.62)),
        Coverage:             models.ToPointer(float64(138.6)),
        NumAps:               models.ToPointer(float64(236.82)),
        NumClients:           models.ToPointer(float64(79.3)),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
    }

}
```

