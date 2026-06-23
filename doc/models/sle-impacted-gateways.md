
# Sle Impacted Gateways

Paginated list of gateways impacted by an SLE metric

## Structure

`SleImpactedGateways`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | Requested SLE classifier filter applied to the query |
| `End` | `*int` | Optional | Last timestamp in the impacted gateways window |
| `Failure` | `*string` | Optional | Requested SLE failure filter applied to the query |
| `Gateways` | [`[]models.SleImpactedGatewaysGateway`](../../doc/models/sle-impacted-gateways-gateway.md) | Optional | Impacted gateway rows returned for an SLE query |
| `Limit` | `*int` | Optional | Maximum number of impacted gateway rows returned per page |
| `Metric` | `*string` | Optional | SLE metric name used for the impacted gateways query |
| `Page` | `*int` | Optional | Current page number for impacted gateway results |
| `Start` | `*int` | Optional | First timestamp in the impacted gateways window |
| `TotalCount` | `*int` | Optional | Number of impacted gateway rows matching the query |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedGateways := models.SleImpactedGateways{
        Classifier:           models.ToPointer("classifier2"),
        End:                  models.ToPointer(234),
        Failure:              models.ToPointer("failure0"),
        Gateways:             []models.SleImpactedGatewaysGateway{
            models.SleImpactedGatewaysGateway{
                Degraded:             models.ToPointer(float64(114.84)),
                Duration:             models.ToPointer(70),
                GatewayMac:           models.ToPointer("gateway_mac2"),
                GatewayModel:         models.ToPointer("gateway_model6"),
                GatewayVersion:       models.ToPointer("gateway_version8"),
            },
            models.SleImpactedGatewaysGateway{
                Degraded:             models.ToPointer(float64(114.84)),
                Duration:             models.ToPointer(70),
                GatewayMac:           models.ToPointer("gateway_mac2"),
                GatewayModel:         models.ToPointer("gateway_model6"),
                GatewayVersion:       models.ToPointer("gateway_version8"),
            },
            models.SleImpactedGatewaysGateway{
                Degraded:             models.ToPointer(float64(114.84)),
                Duration:             models.ToPointer(70),
                GatewayMac:           models.ToPointer("gateway_mac2"),
                GatewayModel:         models.ToPointer("gateway_model6"),
                GatewayVersion:       models.ToPointer("gateway_version8"),
            },
        },
        Limit:                models.ToPointer(64),
    }

}
```

