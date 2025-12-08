
# Radsec

RadSec settings

## Structure

`Radsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CoaEnabled` | `*bool` | Optional | **Default**: `false` |
| `Enabled` | `*bool` | Optional | - |
| `IdleTimeout` | [`*models.RadsecIdleTimeout`](../../doc/models/containers/radsec-idle-timeout.md) | Optional | Radsec Idle Timeout in seconds. Default is 60 |
| `MxclusterIds` | `[]uuid.UUID` | Optional | To use Org mxedges when this WLAN does not use mxtunnel, specify their mxcluster_ids. Org mxedge(s) identified by mxcluster_ids |
| `ProxyHosts` | `[]string` | Optional | Default is site.mxedge.radsec.proxy_hosts which must be a superset of all `wlans[*].radsec.proxy_hosts`. When `radsec.proxy_hosts` are not used, tunnel peers (org or site mxedges) are used irrespective of `use_site_mxedge` |
| `ServerName` | `*string` | Optional | Name of the server to verify (against the cacerts in Org Setting). Only if not Mist Edge. |
| `Servers` | [`[]models.RadsecServer`](../../doc/models/radsec-server.md) | Optional | List of RadSec Servers. Only if not Mist Edge.<br><br>**Constraints**: *Unique Items Required* |
| `UseMxedge` | `*bool` | Optional | use mxedge(s) as RadSec Proxy |
| `UseSiteMxedge` | `*bool` | Optional | To use Site mxedges when this WLAN does not use mxtunnel<br><br>**Default**: `false` |

## Example (as JSON)

```json
{
  "coa_enabled": false,
  "server_name": "radsec.abc.com",
  "use_site_mxedge": false,
  "enabled": false,
  "idle_timeout": 240,
  "mxcluster_ids": [
    "000007d5-0000-0000-0000-000000000000",
    "000007d6-0000-0000-0000-000000000000"
  ],
  "proxy_hosts": [
    "proxy_hosts4"
  ]
}
```

