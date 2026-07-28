
# Tunnel Config Node

Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`

## Structure

`TunnelConfigNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hosts` | `[]string` | Required | Remote gateway host addresses for a tunnel node |
| `InternalIp6s` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `InternalIps` | `[]string` | Optional | Only if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`, `provider`==`custom-ipsec` or `provider`==`custom-gre` |
| `ProbeHostnames` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `ProbeHttp` | [`*models.TunnelConfigNodeProbeHttp`](../../doc/models/tunnel-config-node-probe-http.md) | Optional | HTTP probe settings for a custom IPsec tunnel node |
| `ProbeIp6s` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `ProbeIps` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `RemoteIds` | `[]string` | Optional | Only if `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `WanNames` | `[]string` | Required | Interface names that source tunnel traffic for a tunnel node |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigNode := models.TunnelConfigNode{
        Hosts:                []string{
            "hosts5",
        },
        InternalIp6s:         []string{
            "internal_ip6s2",
            "internal_ip6s3",
            "internal_ip6s4",
        },
        InternalIps:          []string{
            "internal_ips6",
            "internal_ips7",
        },
        ProbeHostnames:       []string{
            "probe_hostnames9",
            "probe_hostnames0",
        },
        ProbeHttp:            models.ToPointer(models.TunnelConfigNodeProbeHttp{
            AcceptedStatusCodes:  []int{
                247,
                248,
            },
            Urls:                 []string{
                "urls5",
                "urls4",
            },
        }),
        ProbeIp6s:            []string{
            "probe_ip6s0",
        },
        WanNames:             []string{
            "wan_names0",
            "wan_names1",
        },
    }

}
```

