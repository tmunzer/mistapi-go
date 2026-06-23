
# Org Site Sle Wired Result

Wired SLE scores and counts for one site

## Structure

`OrgSiteSleWiredResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumClients` | `*float64` | Optional | Number of wired clients included in this site result |
| `NumSwitches` | `*float64` | Optional | Number of switches included in this site result |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `SwitchBandwidth` | `*float64` | Optional | Switch bandwidth SLE score for this site |
| `SwitchHealth` | `float64` | Required | Switch health SLE score for this site |
| `SwitchThroughput` | `*float64` | Optional | Switch throughput SLE score for this site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSiteSleWiredResult := models.OrgSiteSleWiredResult{
        NumClients:           models.ToPointer(float64(185.7)),
        NumSwitches:          models.ToPointer(float64(184.6)),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        SwitchBandwidth:      models.ToPointer(float64(187.9)),
        SwitchHealth:         float64(79.48),
        SwitchThroughput:     models.ToPointer(float64(18.38)),
    }

}
```

