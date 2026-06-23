
# Radsec

RadSec settings for sending RADIUS traffic over TLS

## Structure

`Radsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CoaEnabled` | `*bool` | Optional | Whether RADIUS Change of Authorization (CoA) is enabled for RadSec traffic<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | Whether RadSec is enabled |
| `IdleTimeout` | [`*models.RadsecIdleTimeout`](../../doc/models/containers/radsec-idle-timeout.md) | Optional | RadSec idle timeout in seconds. Default is 60 |
| `MxclusterIds` | `[]uuid.UUID` | Optional | To use Org Mist Edges when this WLAN does not use mxtunnel, specify their mxcluster_ids. Org Mist Edge(s) identified by mxcluster_ids |
| `ProxyHosts` | `[]string` | Optional | Default is site.mxedge.radsec.proxy_hosts which must be a superset of all `wlans[*].radsec.proxy_hosts`. When `radsec.proxy_hosts` are not used, tunnel peers (org or site Mist Edges) are used irrespective of `use_site_mxedge` |
| `ServerName` | `*string` | Optional | TLS server name to verify against the CA certificates in Org Setting. Only if not Mist Edge. |
| `Servers` | [`[]models.RadsecServer`](../../doc/models/radsec-server.md) | Optional | External RadSec servers. Only if not Mist Edge.<br><br>**Constraints**: *Unique Items Required* |
| `UseMxedge` | `*bool` | Optional | Whether to use organization Mist Edge instances as RadSec proxies |
| `UseSiteMxedge` | `*bool` | Optional | Whether to use site Mist Edge instances when this WLAN does not use mxtunnel<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    radsec := models.Radsec{
        CoaEnabled:           models.ToPointer(false),
        Enabled:              models.ToPointer(false),
        IdleTimeout:          models.ToPointer(models.RadsecIdleTimeoutContainer.FromNumber(240)),
        MxclusterIds:         []uuid.UUID{
            uuid.MustParse("000007d5-0000-0000-0000-000000000000"),
            uuid.MustParse("000007d6-0000-0000-0000-000000000000"),
        },
        ProxyHosts:           []string{
            "proxy_hosts4",
        },
        ServerName:           models.ToPointer("radsec.abc.com"),
        UseSiteMxedge:        models.ToPointer(false),
    }

}
```

