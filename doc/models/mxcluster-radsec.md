
# Mxcluster Radsec

RadSec proxy configuration for a Mist Edge cluster. Used when the Mist Edge Cluster is used as a RADIUS Proxy between the local devices and external RADIUS Server.

## Structure

`MxclusterRadsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.MxclusterRadsecAcctServer`](../../doc/models/mxcluster-radsec-acct-server.md) | Optional | List of RADIUS accounting servers, optional, order matters where the first one is treated as primary<br><br>**Constraints**: *Unique Items Required* |
| `AuthServers` | [`[]models.MxclusterRadsecAuthServer`](../../doc/models/mxcluster-radsec-auth-server.md) | Optional | List of RADIUS authentication servers, order matters where the first one is treated as primary<br><br>**Constraints**: *Unique Items Required* |
| `Enabled` | `*bool` | Optional | Whether to enable service on Mist Edge i.e. RADIUS proxy over TLS |
| `MatchSsid` | `*bool` | Optional | Whether to match ssid in request message to select from a subset of RADIUS servers |
| `NasIpSource` | [`*models.MxclusterRadsecNasIpSourceEnum`](../../doc/models/mxcluster-radsec-nas-ip-source-enum.md) | Optional | SSpecify NAS-IP-ADDRESS, NAS-IPv6-ADDRESS to use with auth_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`<br><br>**Default**: `"any"` |
| `ProxyHosts` | `[]string` | Optional | Hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts` |
| `ServerSelection` | [`*models.MxclusterRadsecServerSelectionEnum`](../../doc/models/mxcluster-radsec-server-selection-enum.md) | Optional | When ordered, Mist Edge will prefer and go back to the first RADIUS server if possible. enum: `ordered`, `unordered`<br><br>**Default**: `"ordered"` |
| `SrcIpSource` | [`*models.MxclusterRadsecSrcIpSourceEnum`](../../doc/models/mxcluster-radsec-src-ip-source-enum.md) | Optional | Specify IP address to connect to auth_servers and acct_servers. enum: `any`, `oob`, `oob6`, `tunnel`, `tunnel6`<br><br>**Default**: `"any"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxclusterRadsec := models.MxclusterRadsec{
        AcctServers:          []models.MxclusterRadsecAcctServer{
            models.MxclusterRadsecAcctServer{
                Host:                 models.ToPointer("host4"),
                Port:                 models.ToPointer(254),
                Secret:               models.ToPointer("secret0"),
                Ssids:                []string{
                    "ssids5",
                    "ssids6",
                },
            },
            models.MxclusterRadsecAcctServer{
                Host:                 models.ToPointer("host4"),
                Port:                 models.ToPointer(254),
                Secret:               models.ToPointer("secret0"),
                Ssids:                []string{
                    "ssids5",
                    "ssids6",
                },
            },
            models.MxclusterRadsecAcctServer{
                Host:                 models.ToPointer("host4"),
                Port:                 models.ToPointer(254),
                Secret:               models.ToPointer("secret0"),
                Ssids:                []string{
                    "ssids5",
                    "ssids6",
                },
            },
        },
        AuthServers:          []models.MxclusterRadsecAuthServer{
            models.MxclusterRadsecAuthServer{
                Host:                 models.ToPointer("host0"),
                InbandStatusCheck:    models.ToPointer(false),
                InbandStatusInterval: models.ToPointer(160),
                KeywrapEnabled:       models.ToPointer(false),
                KeywrapFormat:        models.NewOptional(models.ToPointer(models.MxclusterRadAuthServerKeywrapFormatEnum_ASCII)),
            },
            models.MxclusterRadsecAuthServer{
                Host:                 models.ToPointer("host0"),
                InbandStatusCheck:    models.ToPointer(false),
                InbandStatusInterval: models.ToPointer(160),
                KeywrapEnabled:       models.ToPointer(false),
                KeywrapFormat:        models.NewOptional(models.ToPointer(models.MxclusterRadAuthServerKeywrapFormatEnum_ASCII)),
            },
        },
        Enabled:              models.ToPointer(false),
        MatchSsid:            models.ToPointer(false),
        NasIpSource:          models.ToPointer(models.MxclusterRadsecNasIpSourceEnum_ANY),
        ServerSelection:      models.ToPointer(models.MxclusterRadsecServerSelectionEnum_ORDERED),
        SrcIpSource:          models.ToPointer(models.MxclusterRadsecSrcIpSourceEnum_ANY),
    }

}
```

