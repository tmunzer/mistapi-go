
# If Stat Property Servp Info

Service-provider and geolocation details associated with an interface address

## Structure

`IfStatPropertyServpInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Asn` | `*string` | Optional | Autonomous System Number associated with the service provider |
| `City` | `*string` | Optional | Detected city for the service provider address |
| `CountryCode` | `*string` | Optional | ISO country code for the service provider address |
| `Latitude` | `*float64` | Optional | Geographic latitude for the service provider address |
| `Longitude` | `*float64` | Optional | Geographic longitude for the service provider address |
| `Org` | `*string` | Optional | Service provider organization name |
| `RegionCode` | `*string` | Optional | Administrative region code for the service provider address |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ifStatPropertyServpInfo := models.IfStatPropertyServpInfo{
        Asn:                  models.ToPointer("asn6"),
        City:                 models.ToPointer("city8"),
        CountryCode:          models.ToPointer("country_code8"),
        Latitude:             models.ToPointer(float64(79.28)),
        Longitude:            models.ToPointer(float64(174.52)),
    }

}
```

