
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
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `ServpInfo` | [`*models.IfStatPropertyServpInfo`](../../doc/models/if-stat-property-servp-info.md) | Optional | - |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | - |
| `Vlan` | `*int` | Optional | - |
| `WanName` | `*string` | Optional | - |
| `WanType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "address_mode": "address_mode4",
  "ips": [
    "ips8",
    "ips9",
    "ips0"
  ],
  "nat_addresses": [
    "nat_addresses7",
    "nat_addresses8",
    "nat_addresses9"
  ],
  "network_name": "network_name2",
  "port_id": "port_id4"
}
```

