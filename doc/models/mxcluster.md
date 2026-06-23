
# Mxcluster

Mist Edge cluster that groups one or more Mist Edge devices for tunneling, RadSec, and related edge services

*This model accepts additional fields of type interface{}.*

## Structure

`Mxcluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this Mist Edge cluster is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MistDas` | [`*models.MxedgeDas`](../../doc/models/mxedge-das.md) | Optional | Cloud-assisted Dynamic Authorization Service settings for a Mist Edge cluster |
| `MistNac` | [`*models.MxclusterNac`](../../doc/models/mxcluster-nac.md) | Optional | Mist NAC RADIUS settings for a Mist Edge cluster. Used when the Mist Edge Cluster is used as a RADIUS Proxy between the local devices and the Mist NAC |
| `MistNacedge` | [`*models.MxclusterNacedge`](../../doc/models/mxcluster-nacedge.md) | Optional | NAC Edge survivability settings for a Mist Edge cluster. Requires `mist_nac` to be enabled on the cluster |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | Management settings for a Mist Edge appliance |
| `Name` | `*string` | Optional | Display name of the Mist Edge cluster |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Proxy` | [`*models.Proxy`](../../doc/models/proxy.md) | Optional | Proxy Configuration to talk to Mist |
| `Radsec` | [`*models.MxclusterRadsec`](../../doc/models/mxcluster-radsec.md) | Optional | RadSec proxy configuration for a Mist Edge cluster. Used when the Mist Edge Cluster is used as a RADIUS Proxy between the local devices and external RADIUS Server. |
| `RadsecTls` | [`*models.MxclusterRadsecTls`](../../doc/models/mxcluster-radsec-tls.md) | Optional | TLS settings for RadSec on a Mist Edge cluster |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TuntermApSubnets` | `[]string` | Optional | List of subnets where we allow AP to establish Mist Tunnels from |
| `TuntermDhcpdConfig` | [`*models.TuntermDhcpdConfig`](../../doc/models/tunterm-dhcpd-config.md) | Optional | DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID |
| `TuntermExtraRoutes` | [`map[string]models.MxclusterTuntermExtraRoute`](../../doc/models/mxcluster-tunterm-extra-route.md) | Optional | Extra routes for Mist Tunneled VLANs. Property key is a CIDR |
| `TuntermHosts` | `[]string` | Optional | Hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP) |
| `TuntermHostsOrder` | `[]int` | Optional | List of index of tunterm_hosts |
| `TuntermHostsSelection` | [`*models.MxclusterTuntermHostsSelectionEnum`](../../doc/models/mxcluster-tunterm-hosts-selection-enum.md) | Optional | Ordering of tunterm_hosts for Mist Edge within the same mxcluster. enum:<br><br>* `shuffle`: the ordering of tunterm_hosts is randomized by the device''s MAC<br>* `shuffle-by-site`: shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels)<br>* `ordered`: order decided by tunterm_hosts_order<br><br>**Default**: `"shuffle"` |
| `TuntermMonitoring` | [`[][]models.TuntermMonitoringItem`](../../doc/models/tunterm-monitoring-item.md) | Optional | Tunnel termination monitoring checks for a Mist Edge cluster |
| `TuntermMonitoringDisabled` | `*bool` | Optional | Whether tunnel termination monitoring is disabled for the cluster |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxcluster := models.Mxcluster{
        CreatedTime:               models.ToPointer(float64(209.42)),
        ForSite:                   models.ToPointer(false),
        Id:                        models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MistDas:                   models.ToPointer(models.MxedgeDas{
            CoaServers:           []models.MxedgeDasCoaServer{
                models.MxedgeDasCoaServer{
                    DisableEventTimestampCheck:  models.ToPointer(false),
                    Enabled:                     models.ToPointer(false),
                    Host:                        models.ToPointer("host8"),
                    Port:                        models.ToPointer(28),
                    RequireMessageAuthenticator: models.ToPointer(false),
                },
                models.MxedgeDasCoaServer{
                    DisableEventTimestampCheck:  models.ToPointer(false),
                    Enabled:                     models.ToPointer(false),
                    Host:                        models.ToPointer("host8"),
                    Port:                        models.ToPointer(28),
                    RequireMessageAuthenticator: models.ToPointer(false),
                },
            },
            Enabled:              models.ToPointer(false),
        }),
        MistNac:                   models.ToPointer(models.MxclusterNac{
            AcctServerPort:       models.ToPointer(70),
            AuthServerPort:       models.ToPointer(34),
            ClientIps:            map[string]models.MxclusterNacClientIp{
                "key0": models.MxclusterNacClientIp{
                    RequireMessageAuthenticator: models.ToPointer(false),
                    Secret:                      models.ToPointer("secret4"),
                    SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                    Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
                },
                "key1": models.MxclusterNacClientIp{
                    RequireMessageAuthenticator: models.ToPointer(false),
                    Secret:                      models.ToPointer("secret4"),
                    SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                    Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
                },
                "key2": models.MxclusterNacClientIp{
                    RequireMessageAuthenticator: models.ToPointer(false),
                    Secret:                      models.ToPointer("secret4"),
                    SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                    Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
                },
            },
            Enabled:              models.ToPointer(false),
            Secret:               models.ToPointer("secret6"),
        }),
        OrgId:                     models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:                    models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        TuntermHostsSelection:     models.ToPointer(models.MxclusterTuntermHostsSelectionEnum_SHUFFLE),
        AdditionalProperties:      map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

