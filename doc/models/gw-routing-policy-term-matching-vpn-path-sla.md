
# Gw Routing Policy Term Matching Vpn Path Sla

SLA thresholds used when matching VPN paths

## Structure

`GwRoutingPolicyTermMatchingVpnPathSla`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxJitter` | `models.Optional[int]` | Optional | Maximum jitter threshold allowed for the VPN path |
| `MaxLatency` | `models.Optional[int]` | Optional | Maximum latency threshold allowed for the VPN path |
| `MaxLoss` | `models.Optional[int]` | Optional | Maximum packet-loss threshold allowed for the VPN path |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gwRoutingPolicyTermMatchingVpnPathSla := models.GwRoutingPolicyTermMatchingVpnPathSla{
        MaxJitter:            models.NewOptional(models.ToPointer(186)),
        MaxLatency:           models.NewOptional(models.ToPointer(1500)),
        MaxLoss:              models.NewOptional(models.ToPointer(30)),
    }

}
```

