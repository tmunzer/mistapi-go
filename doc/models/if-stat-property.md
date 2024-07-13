
# If Stat Property

## Structure

`IfStatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AddressMode` | `*string` | Optional | - |
| `Ips` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `NatAddresses` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `NetworkName` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `PortUsage` | `*string` | Optional | - |
| `RedundancyState` | `*string` | Optional | - |
| `RxBytes` | `*int` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `ServpInfo` | [`*models.IfStatPropertyServpInfo`](../../doc/models/if-stat-property-servp-info.md) | Optional | - |
| `TxBytes` | `*int` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `Vlan` | `*int` | Optional | - |
| `WanName` | `*string` | Optional | - |
| `WanType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "address_mode": "address_mode2",
  "ips": [
    "ips4"
  ],
  "nat_addresses": [
    "nat_addresses3"
  ],
  "network_name": "network_name8",
  "port_id": "port_id0"
}
```

