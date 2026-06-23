
# If Stat Property

Interface statistics and metadata reported by a device

## Structure

`IfStatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AddressMode` | `*string` | Optional | Address assignment mode reported for the interface |
| `Ips` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `NatAddresses` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `NetworkName` | `*string` | Optional | Mist network name associated with this interface |
| `PortId` | `*string` | Optional | Physical or logical port identifier for this interface |
| `PortUsage` | `*string` | Optional | Configured usage for this interface, such as LAN or WAN |
| `RedundancyState` | `*string` | Optional | Redundancy state reported for this interface |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `ServpInfo` | [`*models.IfStatPropertyServpInfo`](../../doc/models/if-stat-property-servp-info.md) | Optional | Service-provider and geolocation details associated with an interface address |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | Whether the interface is operationally up |
| `Vlan` | `*int` | Optional | Associated VLAN ID for this interface |
| `WanName` | `*string` | Optional | Configured WAN name associated with this interface |
| `WanType` | `*string` | Optional | Configured WAN uplink type associated with this interface |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ifStatProperty := models.IfStatProperty{
        AddressMode:          models.ToPointer("address_mode2"),
        Ips:                  []string{
            "ips4",
            "ips5",
            "ips6",
        },
        NatAddresses:         []string{
            "nat_addresses3",
            "nat_addresses4",
            "nat_addresses5",
        },
        NetworkName:          models.ToPointer("network_name8"),
        PortId:               models.ToPointer("port_id0"),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
    }

}
```

