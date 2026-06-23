
# Site Mxtunnel Additional Mxtunnel

Additional named Mist Tunnel configuration for a site

## Structure

`SiteMxtunnelAdditionalMxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Clusters` | [`[]models.SiteMxtunnelCluster`](../../doc/models/site-mxtunnel-cluster.md) | Optional | Tunnel peer cluster definitions APs use for an additional Mist Tunnel |
| `HelloInterval` | `*int` | Optional | In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries<br><br>**Default**: `60`<br><br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `*int` | Optional | Number of missed hello heartbeats before an AP tries another tunnel peer<br><br>**Default**: `7`<br><br>**Constraints**: `>= 2`, `<= 30` |
| `Protocol` | [`*models.SiteMxtunnelProtocolEnum`](../../doc/models/site-mxtunnel-protocol-enum.md) | Optional | Encapsulation protocol used for this additional Mist Tunnel. enum: `ip`, `udp` |
| `VlanIds` | `[]int` | Optional | VLAN IDs carried by an additional site Mist Tunnel |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteMxtunnelAdditionalMxtunnel := models.SiteMxtunnelAdditionalMxtunnel{
        Clusters:             []models.SiteMxtunnelCluster{
            models.SiteMxtunnelCluster{
                Name:                 models.ToPointer("name6"),
                TuntermHosts:         []string{
                    "tunterm_hosts6",
                },
            },
            models.SiteMxtunnelCluster{
                Name:                 models.ToPointer("name6"),
                TuntermHosts:         []string{
                    "tunterm_hosts6",
                },
            },
        },
        HelloInterval:        models.ToPointer(60),
        HelloRetries:         models.ToPointer(3),
        Protocol:             models.ToPointer(models.SiteMxtunnelProtocolEnum_UDP),
        VlanIds:              []int{
            300,
            310,
            320,
        },
    }

}
```

