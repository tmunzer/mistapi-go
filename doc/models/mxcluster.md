
# Mxcluster

mxCluster

## Structure

`Mxcluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `MistDas` | [`*models.MxedgeDas`](../../doc/models/mxedge-das.md) | Optional | configure cloud-assisted dynamic authorization service on this cluster of mist edges |
| `MistNac` | [`*models.MxclusterNac`](../../doc/models/mxcluster-nac.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Proxy` | [`*models.Proxy`](../../doc/models/proxy.md) | Optional | Proxy Configuration to talk to Mist |
| `Radsec` | [`*models.MxclusterRadsec`](../../doc/models/mxcluster-radsec.md) | Optional | MxEdge Radsec Configuration |
| `RadsecTls` | [`*models.MxclusterRadsecTls`](../../doc/models/mxcluster-radsec-tls.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TuntermApSubnets` | `[]string` | Optional | list of subnets where we allow AP to establish Mist Tunnels from |
| `TuntermDhcpdConfig` | [`*models.TuntermDhcpdConfig`](../../doc/models/tunterm-dhcpd-config.md) | Optional | DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID |
| `TuntermExtraRoutes` | [`map[string]models.MxclusterTuntermExtraRoute`](../../doc/models/mxcluster-tunterm-extra-route.md) | Optional | extra routes for Mist Tunneled VLANs. Property key is a CIDR |
| `TuntermHosts` | `[]string` | Optional | hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP) |
| `TuntermHostsOrder` | `[]int` | Optional | list of index of tunterm_hosts |
| `TuntermHostsSelection` | [`*models.MxclusterTuntermHostsSelectionEnum`](../../doc/models/mxcluster-tunterm-hosts-selection-enum.md) | Optional | Ordering of tunterm_hosts for mxedge within the same mxcluster. enum:<br><br>* `shuffle`: the ordering of tunterm_hosts is randomized by the device''s MAC<br>* `shuffle-by-site`: shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels)<br>* `ordered`: order decided by tunterm_hosts_order<br>**Default**: `"shuffle"` |
| `TuntermMonitoring` | [`[]models.TuntermMonitoringItem`](../../doc/models/tunterm-monitoring-item.md) | Optional | - |
| `TuntermMonitoringDisabled` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "tunterm_hosts_selection": "shuffle",
  "created_time": 142.92,
  "for_site": false,
  "id": "0000067e-0000-0000-0000-000000000000",
  "mist_das": {
    "coa_servers": [
      {
        "disable_event_timestamp_check": false,
        "enabled": false,
        "host": "host8",
        "port": 28,
        "secret": "secret2"
      },
      {
        "disable_event_timestamp_check": false,
        "enabled": false,
        "host": "host8",
        "port": 28,
        "secret": "secret2"
      }
    ],
    "enabled": false
  },
  "mist_nac": {
    "acct_server_port": 70,
    "auth_server_port": 34,
    "client_ips": {
      "key0": {
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-meraki"
      },
      "key1": {
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-meraki"
      },
      "key2": {
        "secret": "secret4",
        "site_id": "0000197c-0000-0000-0000-000000000000",
        "vendor": "cisco-meraki"
      }
    },
    "enabled": false,
    "secret": "secret6"
  }
}
```

