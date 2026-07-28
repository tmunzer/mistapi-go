
# Gateway Wan Probe Override

Only if `usage`==`wan`. WAN health probe override for this gateway port

## Structure

`GatewayWanProbeOverride`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostnames` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Http` | [`*models.GatewayWanProbeOverrideHttp`](../../doc/models/gateway-wan-probe-override-http.md) | Optional | HTTP probe settings for a WAN probe override |
| `Ip6s` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Ips` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `ProbeProfile` | [`*models.GatewayWanProbeOverrideProbeProfileEnum`](../../doc/models/gateway-wan-probe-override-probe-profile-enum.md) | Optional | WAN probe profile used for health checks on this port. enum: `broadband`, `lte`<br><br>**Default**: `"broadband"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayWanProbeOverride := models.GatewayWanProbeOverride{
        Hostnames:            []string{
            "hostnames6",
        },
        Http:                 models.ToPointer(models.GatewayWanProbeOverrideHttp{
            AcceptedStatusCodes:  []int{
                189,
                190,
            },
            Urls:                 []string{
                "urls9",
                "urls8",
            },
        }),
        Ip6s:                 []string{
            "ip6s0",
            "ip6s1",
            "ip6s2",
        },
        Ips:                  []string{
            "ips2",
            "ips3",
            "ips4",
        },
        ProbeProfile:         models.ToPointer(models.GatewayWanProbeOverrideProbeProfileEnum_BROADBAND),
    }

}
```

