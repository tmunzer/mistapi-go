
# Stats Ap Ble

BLE beacon and traffic statistics reported by an AP

## Structure

`StatsApBle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconEnabled` | `models.Optional[bool]` | Optional, Read-only | Whether Mist BLE beacon transmission is enabled |
| `BeaconRate` | `models.Optional[int]` | Optional, Read-only | Mist BLE beacon transmit rate, in beacons per second |
| `EddystoneUidEnabled` | `models.Optional[bool]` | Optional, Read-only | Whether Eddystone-UID beacon transmission is enabled |
| `EddystoneUidFreqMsec` | `models.Optional[int]` | Optional, Read-only | Interval for Eddystone-UID advertisements, in milliseconds |
| `EddystoneUidInstance` | `models.Optional[string]` | Optional, Read-only | Eddystone-UID instance value broadcast by the AP |
| `EddystoneUidNamespace` | `models.Optional[string]` | Optional, Read-only | Eddystone-UID namespace value broadcast by the AP |
| `EddystoneUrlEnabled` | `models.Optional[bool]` | Optional, Read-only | Whether Eddystone-URL beacon transmission is enabled |
| `EddystoneUrlFreqMsec` | `models.Optional[int]` | Optional, Read-only | Interval for Eddystone-URL advertisements, in milliseconds |
| `EddystoneUrlUrl` | `models.Optional[string]` | Optional, Read-only | URL broadcast by the Eddystone-URL beacon |
| `IbeaconEnabled` | `models.Optional[bool]` | Optional, Read-only | Whether iBeacon transmission is enabled |
| `IbeaconFreqMsec` | `models.Optional[int]` | Optional, Read-only | Interval for iBeacon advertisements, in milliseconds |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | iBeacon UUID value, or null when no iBeacon UUID is configured |
| `Major` | `models.Optional[int]` | Optional, Read-only | Reported iBeacon major value for BLE statistics |
| `Minors` | `[]int` | Optional | List of integer values |
| `Power` | `models.Optional[int]` | Optional, Read-only | BLE transmit power setting reported by the AP |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `TxResets` | `models.Optional[int]` | Optional, Read-only | Resets due to tx hung |
| `Uuid` | `models.Optional[uuid.UUID]` | Optional, Read-only | Beacon UUID reported by BLE statistics |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsApBle := models.StatsApBle{
        BeaconEnabled:         models.NewOptional(models.ToPointer(false)),
        BeaconRate:            models.NewOptional(models.ToPointer(3)),
        EddystoneUidEnabled:   models.NewOptional(models.ToPointer(false)),
        EddystoneUidFreqMsec:  models.NewOptional(models.ToPointer(2000)),
        EddystoneUidInstance:  models.NewOptional(models.ToPointer("5c5b35000001")),
        EddystoneUidNamespace: models.NewOptional(models.ToPointer("2818e3868dec25629ede")),
        EddystoneUrlEnabled:   models.NewOptional(models.ToPointer(true)),
        EddystoneUrlFreqMsec:  models.NewOptional(models.ToPointer(100)),
        EddystoneUrlUrl:       models.NewOptional(models.ToPointer("https://www.abc.com")),
        IbeaconEnabled:        models.NewOptional(models.ToPointer(true)),
        IbeaconFreqMsec:       models.NewOptional(models.ToPointer(2000)),
        IbeaconMajor:          models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:          models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:           models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
        Major:                 models.NewOptional(models.ToPointer(12345)),
        Power:                 models.NewOptional(models.ToPointer(10)),
        RxBytes:               models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:                models.NewOptional(models.ToPointer(int64(57770567))),
        TxBytes:               models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:                models.NewOptional(models.ToPointer(int64(812204062))),
        TxResets:              models.NewOptional(models.ToPointer(0)),
        Uuid:                  models.NewOptional(models.ToPointer(uuid.MustParse("ada72f8f-1643-e5c6-94db-f2a5636f1a64"))),
    }

}
```

