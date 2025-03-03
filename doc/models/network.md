
# Network

Networks are usually subnets that have cross-site significance. `networks`in Org Settings will got merged into `networks`in Site Setting. For gateways, they can be used to define Service Routes.

*This model accepts additional fields of type interface{}.*

## Structure

`Network`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DisallowMistServices` | `*bool` | Optional | Whether to disallow Mist Devices in the network<br>**Default**: `false` |
| `Gateway` | `*string` | Optional | - |
| `Gateway6` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `InternalAccess` | [`*models.NetworkInternalAccess`](../../doc/models/network-internal-access.md) | Optional | - |
| `InternetAccess` | [`*models.NetworkInternetAccess`](../../doc/models/network-internet-access.md) | Optional | Whether this network has direct internet access |
| `Isolation` | `*bool` | Optional | Whether to allow clients in the network to talk to each other |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Multicast` | [`*models.NetworkMulticast`](../../doc/models/network-multicast.md) | Optional | Whether to enable multicast support (only PIM-sparse mode is supported) |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RoutedForNetworks` | `[]string` | Optional | For a Network (usually LAN), it can be routable to other networks (e.g. OSPF) |
| `Subnet` | `*string` | Optional | - |
| `Subnet6` | `*string` | Optional | - |
| `Tenants` | [`map[string]models.NetworkTenant`](../../doc/models/network-tenant.md) | Optional | Property key must be the user/tenant name (i.e. "printer-1") or a Variable (i.e. "{{myvar}}") |
| `VlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |
| `VpnAccess` | [`map[string]models.NetworkVpnAccessConfig`](../../doc/models/network-vpn-access-config.md) | Optional | Property key is the VPN name. Whether this network can be accessed from vpn |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disallow_mist_services": false,
  "gateway": "192.168.70.1",
  "gateway6": "fdad:b0bc:f29e::1",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "subnet": "192.168.70.0/24",
  "subnet6": "fdad:b0bc:f29e::/32",
  "created_time": 103.24,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

