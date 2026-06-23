
# Stats Wired Client

Wired client statistics for a client connected through an AP Ethernet port

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWiredClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthState` | `*string` | Optional | Authorization state reported for the wired client<br><br>**Constraints**: *Minimum Length*: `1` |
| `DeviceId` | `*string` | Optional | Identifier of the AP the wired client is connected to<br><br>**Constraints**: *Minimum Length*: `1` |
| `EthPort` | `*string` | Optional | AP Ethernet port where the wired client is connected<br><br>**Constraints**: *Minimum Length*: `1` |
| `LastSeen` | `*float64` | Optional | Time when transmit or receive traffic was last observed for the wired client |
| `Mac` | `string` | Required | Wired client MAC address observed on the AP Ethernet port<br><br>**Constraints**: *Minimum Length*: `1` |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Uptime` | `*float64` | Optional | Elapsed time since the wired client connected, in seconds |
| `VlanId` | `*float64` | Optional | VLAN identifier used by the wired client, when present |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWiredClient := models.StatsWiredClient{
        AuthState:            models.ToPointer("auth_state8"),
        DeviceId:             models.ToPointer("device_id4"),
        EthPort:              models.ToPointer("eth_port6"),
        LastSeen:             models.ToPointer(float64(148.08)),
        Mac:                  "mac2",
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

