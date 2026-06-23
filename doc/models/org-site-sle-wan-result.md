
# Org Site Sle Wan Result

WAN SLE scores and counts for one site

## Structure

`OrgSiteSleWanResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApplicationHealth` | `*float64` | Optional | Application health SLE score for this site |
| `GatewayHealth` | `float64` | Required | Gateway health SLE score for this site |
| `NumClients` | `*float64` | Optional | Number of WAN clients included in this site result |
| `NumGateways` | `*float64` | Optional | Number of gateways included in this site result |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `WanLinkHealth` | `*float64` | Optional | WAN link health SLE score for this site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSiteSleWanResult := models.OrgSiteSleWanResult{
        ApplicationHealth:    models.ToPointer(float64(121.14)),
        GatewayHealth:        float64(110.88),
        NumClients:           models.ToPointer(float64(199.82)),
        NumGateways:          models.ToPointer(float64(88.46)),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        WanLinkHealth:        models.ToPointer(float64(9.68)),
    }

}
```

