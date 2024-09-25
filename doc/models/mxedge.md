
# Mxedge

MxEdge

## Structure

`Mxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Magic` | `*string` | Optional | - |
| `Model` | `string` | Required | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `MxagentRegistered` | `*bool` | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | MxCluster this MxEdge belongs to |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | - |
| `Name` | `string` | Required | - |
| `Note` | `*string` | Optional | - |
| `NtpServers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OobIpConfig` | [`*models.MxedgeOobIpConfig`](../../doc/models/mxedge-oob-ip-config.md) | Optional | ip configuration of the Mist Edge out-of_band management interface |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Proxy` | [`*models.Proxy`](../../doc/models/proxy.md) | Optional | Proxy Configuration to talk to Mist |
| `Services` | `[]string` | Optional | list of services to run, tunterm only for now |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TuntermDhcpdConfig` | [`*models.MxedgeTuntermDhcpdConfig`](../../doc/models/mxedge-tunterm-dhcpd-config.md) | Optional | global and per-VLAN. Property key is the VLAN ID |
| `TuntermExtraRoutes` | [`map[string]models.MxedgeTuntermExtraRoute`](../../doc/models/mxedge-tunterm-extra-route.md) | Optional | Property key is a CIDR |
| `TuntermIgmpSnoopingConfig` | [`*models.MxedgeTuntermIgmpSnoopingConfig`](../../doc/models/mxedge-tunterm-igmp-snooping-config.md) | Optional | - |
| `TuntermIpConfig` | [`*models.MxedgeTuntermIpConfig`](../../doc/models/mxedge-tunterm-ip-config.md) | Optional | ip configuration of the Mist Tunnel interface |
| `TuntermMonitoring` | [`[][]models.TuntermMonitoringItem`](../../doc/models/tunterm-monitoring-item.md) | Optional | - |
| `TuntermMulticastConfig` | [`*models.MxedgeTuntermMulticastConfig`](../../doc/models/mxedge-tunterm-multicast-config.md) | Optional | - |
| `TuntermOtherIpConfigs` | [`map[string]models.MxedgeTuntermOtherIpConfig`](../../doc/models/mxedge-tunterm-other-ip-config.md) | Optional | ip configs by VLAN ID. Property key is the VLAN ID |
| `TuntermPortConfig` | [`*models.TuntermPortConfig`](../../doc/models/tunterm-port-config.md) | Optional | ethernet port configurations |
| `TuntermRegistered` | `*bool` | Optional | - |
| `TuntermSwitchConfig` | [`*models.MxedgeTuntermSwitchConfigs`](../../doc/models/mxedge-tunterm-switch-configs.md) | Optional | if custom vlan settings are desired |
| `Versions` | [`*models.MxedgeVersions`](../../doc/models/mxedge-versions.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
  "model": "ME-100",
  "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
  "name": "Guest",
  "note": "note for mxedge",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 184.84,
  "for_site": false,
  "modified_time": 150.12
}
```

