
# Tunnel Config Node

Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`

## Structure

`TunnelConfigNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hosts` | `[]string` | Required | Remote gateway host addresses for a tunnel node |
| `InternalIps` | `[]string` | Optional | Only if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`, `provider`==`custom-ipsec` or `provider`==`custom-gre` |
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
        InternalIps:          []string{
            "internal_ips6",
            "internal_ips7",
        },
        ProbeIps:             []string{
            "probe_ips9",
        },
        RemoteIds:            []string{
            "remote_ids7",
        },
        WanNames:             []string{
            "wan_names0",
            "wan_names1",
        },
    }

}
```

