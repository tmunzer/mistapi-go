
# Network

Networks are usually subnets that have cross-site significance. `networks`in Org Settings will got merged into `networks`in Site Setting. For gateways, they can be used to define Service Routes.

## Structure

`Network`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `DisallowMistServices` | `*bool` | Optional | whether to disallow Mist Devices in the network<br>**Default**: `false` |
| `Gateway` | `*string` | Optional | - |
| `Gateway6` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `InternalAccess` | [`*models.NetworkInternalAccess`](../../doc/models/network-internal-access.md) | Optional | - |
| `InternetAccess` | [`*models.NetworkInternetAccess`](../../doc/models/network-internet-access.md) | Optional | whether this network has direct internet access |
| `Isolation` | `*bool` | Optional | whether to allow clients in the network to talk to each other |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RoutedForNetworks` | `[]string` | Optional | for a Network (usually LAN), it can be routable to other networks (e.g. OSPF) |
| `Subnet` | `*string` | Optional | - |
| `Subnet6` | `*string` | Optional | - |
| `Tenants` | [`map[string]models.NetworkTenant`](../../doc/models/network-tenant.md) | Optional | - |
| `VlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |
| `VpnAccess` | [`map[string]models.NetworkVpnAccessConfig`](../../doc/models/network-vpn-access-config.md) | Optional | Property key is the VPN name. Whether this network can be accessed from vpn |

## Example (as JSON)

```json
{
  "disallow_mist_services": false,
  "gateway": "192.168.70.1",
  "gateway6": "fdad:b0bc:f29e::1",
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "subnet": "192.168.70.0/24",
  "subnet6": "fdad:b0bc:f29e::/32",
  "created_time": 103.24,
  "id": "0000082e-0000-0000-0000-000000000000"
}
```

