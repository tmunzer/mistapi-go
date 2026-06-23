
# Site Mxtunnel Cluster

Mist Tunnel peer cluster definition for a site

## Structure

`SiteMxtunnelCluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Peer cluster name used in the site Mist Tunnel configuration |
| `TuntermHosts` | `[]string` | Optional | Tunnel termination hostnames or IP addresses in a site Mist Tunnel cluster |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteMxtunnelCluster := models.SiteMxtunnelCluster{
        Name:                 models.ToPointer("primary"),
        TuntermHosts:         []string{
            "mxedge1",
            "mxedge2.local",
        },
    }

}
```

