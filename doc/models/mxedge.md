
# Mxedge

Mist Edge appliance configuration and registration state

*This model accepts additional fields of type interface{}.*

## Structure

`Mxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this Mist Edge is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional, Read-only | Mist Edge MAC address |
| `Magic` | `*string` | Optional, Read-only | Registration claim code for the Mist Edge |
| `Model` | `string` | Required | Mist Edge hardware or virtual appliance model |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MxagentRegistered` | `*bool` | Optional, Read-only | Whether the Mist Edge agent has registered with Mist cloud |
| `MxclusterId` | `models.Optional[uuid.UUID]` | Optional | Mist Edge cluster identifier that this appliance belongs to |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | Management settings for a Mist Edge appliance |
| `Name` | `string` | Required | Display name of the Mist Edge |
| `Notes` | `*string` | Optional | Free-form notes for the Mist Edge |
| `NtpServers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `OobIpConfig` | [`*models.MxedgeOobIpConfig`](../../doc/models/mxedge-oob-ip-config.md) | Optional | IP configuration for the Mist Edge out-of-band management interface |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Proxy` | [`*models.Proxy`](../../doc/models/proxy.md) | Optional | Proxy Configuration to talk to Mist |
| `Services` | `[]string` | Optional | List of services to run, tunterm only for now |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TuntermDhcpdConfig` | [`*models.MxedgeTuntermDhcpdConfig`](../../doc/models/mxedge-tunterm-dhcpd-config.md) | Optional | Global and per-VLAN DHCP relay settings for Mist Tunneled VLANs; property key is the VLAN ID |
| `TuntermExtraRoutes` | [`map[string]models.MxedgeTuntermExtraRoute`](../../doc/models/mxedge-tunterm-extra-route.md) | Optional | Extra routes for Mist Tunneled VLAN traffic; property key is a CIDR |
| `TuntermIgmpSnoopingConfig` | [`*models.MxedgeTuntermIgmpSnoopingConfig`](../../doc/models/mxedge-tunterm-igmp-snooping-config.md) | Optional | IGMP snooping settings for Mist Tunnel VLANs |
| `TuntermIpConfig` | [`*models.MxedgeTuntermIpConfig`](../../doc/models/mxedge-tunterm-ip-config.md) | Optional | IP configuration for the Mist Tunnel interface |
| `TuntermMonitoring` | [`[][]models.TuntermMonitoringItem`](../../doc/models/tunterm-monitoring-item.md) | Optional | Tunnel termination monitoring checks for a Mist Edge |
| `TuntermMulticastConfig` | [`*models.MxedgeTuntermMulticastConfig`](../../doc/models/mxedge-tunterm-multicast-config.md) | Optional | Multicast forwarding settings for tunnel termination |
| `TuntermOtherIpConfigs` | [`map[string]models.MxedgeTuntermOtherIpConfig`](../../doc/models/mxedge-tunterm-other-ip-config.md) | Optional | IPconfigs by VLAN ID. Property key is the VLAN ID |
| `TuntermPortConfig` | [`*models.TuntermPortConfig`](../../doc/models/tunterm-port-config.md) | Optional | Ethernet port configuration for tunnel termination interfaces |
| `TuntermRegistered` | `*bool` | Optional, Read-only | Whether the tunnel termination service has registered with Mist cloud |
| `TuntermSwitchConfig` | [`*models.MxedgeTuntermSwitchConfigs`](../../doc/models/mxedge-tunterm-switch-configs.md) | Optional | Custom VLAN settings for tunnel termination switch ports, if desired; property key is the port identifier |
| `Versions` | [`*models.MxedgeVersions`](../../doc/models/mxedge-versions.md) | Optional, Read-only | Read-only Mist Edge service versions |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxedge := models.Mxedge{
        CreatedTime:               models.ToPointer(float64(184.84)),
        ForSite:                   models.ToPointer(false),
        Id:                        models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Mac:                       models.ToPointer("0200009fbe65"),
        Magic:                     models.ToPointer("L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD"),
        Model:                     "ME-100",
        MxclusterId:               models.NewOptional(models.ToPointer(uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4"))),
        Name:                      "Guest",
        Notes:                     models.ToPointer("note for mxedge"),
        OrgId:                     models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:                    models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties:      map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

