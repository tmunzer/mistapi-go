
# Ble Config

Bluetooth Low Energy beacon and asset advertising settings for an AP

## Structure

`BleConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconEnabled` | `*bool` | Optional | Whether Mist beacons is enabled<br><br>**Default**: `true` |
| `BeaconRate` | `*int` | Optional | Required if `beacon_rate_mode`==`custom`, 1-10, in number-beacons-per-second |
| `BeaconRateMode` | [`*models.BleConfigBeaconRateModeEnum`](../../doc/models/ble-config-beacon-rate-mode-enum.md) | Optional | Beacon rate mode for Mist BLE beacons; use custom to set beacon_rate. enum: `custom`, `default`<br><br>**Default**: `"default"` |
| `BeamDisabled` | `[]int` | Optional | List of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam) |
| `CustomBlePacketEnabled` | `*bool` | Optional | Can be enabled if `beacon_enabled`==`true`, whether to send custom packet<br><br>**Default**: `false` |
| `CustomBlePacketFrame` | `*string` | Optional | The custom frame to be sent out in this beacon. The frame must be a hexstring |
| `CustomBlePacketFreqMsec` | `*int` | Optional | Frequency (msec) of data emitted by custom ble beacon<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0` |
| `EddystoneUidAdvPower` | `*int` | Optional | Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default<br><br>**Default**: `0`<br><br>**Constraints**: `>= -100`, `<= 20` |
| `EddystoneUidBeams` | `*string` | Optional | BLE beams used to transmit Eddystone-UID advertisements, expressed as ranges such as `2-4,7` |
| `EddystoneUidEnabled` | `*bool` | Optional | Only if `beacon_enabled`==`false`, Whether Eddystone-UID beacon is enabled<br><br>**Default**: `false` |
| `EddystoneUidFreqMsec` | `*int` | Optional | Frequency (msec) of data emit by Eddystone-UID beacon<br><br>**Default**: `0` |
| `EddystoneUidInstance` | `*string` | Optional | Eddystone-UID instance for the device |
| `EddystoneUidNamespace` | `*string` | Optional | Eddystone-UID namespace broadcast by the AP, as a 10-byte hex string |
| `EddystoneUrlAdvPower` | `*int` | Optional | Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default<br><br>**Default**: `0`<br><br>**Constraints**: `>= -100`, `<= 20` |
| `EddystoneUrlBeams` | `*string` | Optional | BLE beams used to transmit Eddystone-URL advertisements, expressed as ranges such as `2-4,7` |
| `EddystoneUrlEnabled` | `*bool` | Optional | Only if `beacon_enabled`==`false`, Whether Eddystone-URL beacon is enabled<br><br>**Default**: `false` |
| `EddystoneUrlFreqMsec` | `*int` | Optional | Frequency (msec) of data emitted by Eddystone-URL beacon<br><br>**Default**: `0` |
| `EddystoneUrlUrl` | `*string` | Optional | URL pointed by Eddystone-URL beacon |
| `IbeaconAdvPower` | `*int` | Optional | Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default<br><br>**Default**: `0`<br><br>**Constraints**: `>= -100`, `<= 20` |
| `IbeaconBeams` | `*string` | Optional | BLE beams used to transmit iBeacon advertisements, expressed as ranges such as `2-4,7` |
| `IbeaconEnabled` | `*bool` | Optional | Can be enabled if `beacon_enabled`==`true`, whether to send iBeacon<br><br>**Default**: `false` |
| `IbeaconFreqMsec` | `*int` | Optional | Frequency (msec) of data emit for iBeacon<br><br>**Default**: `0` |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `*uuid.UUID` | Optional | Optional, if not specified, the same UUID as the beacon will be used |
| `Power` | `*int` | Optional | Required if `power_mode`==`custom`; else use `power_mode` as default<br><br>**Default**: `9`<br><br>**Constraints**: `>= 1`, `<= 10` |
| `PowerMode` | [`*models.BleConfigPowerModeEnum`](../../doc/models/ble-config-power-mode-enum.md) | Optional | Transmit power mode for BLE beacons; use `custom` to set explicit power. enum: `custom`, `default`<br><br>**Default**: `"default"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    bleConfig := models.BleConfig{
        BeaconEnabled:           models.ToPointer(true),
        BeaconRate:              models.ToPointer(3),
        BeaconRateMode:          models.ToPointer(models.BleConfigBeaconRateModeEnum_CUSTOM),
        BeamDisabled:            []int{
            1,
            3,
            6,
        },
        CustomBlePacketEnabled:  models.ToPointer(false),
        CustomBlePacketFrame:    models.ToPointer("0x........"),
        CustomBlePacketFreqMsec: models.ToPointer(300),
        EddystoneUidAdvPower:    models.ToPointer(-65),
        EddystoneUidBeams:       models.ToPointer("2-4,7"),
        EddystoneUidEnabled:     models.ToPointer(false),
        EddystoneUidFreqMsec:    models.ToPointer(200),
        EddystoneUidInstance:    models.ToPointer("5c5b35000001"),
        EddystoneUidNamespace:   models.ToPointer("2818e3868dec25629ede"),
        EddystoneUrlAdvPower:    models.ToPointer(-65),
        EddystoneUrlBeams:       models.ToPointer("2-4,7"),
        EddystoneUrlEnabled:     models.ToPointer(false),
        EddystoneUrlFreqMsec:    models.ToPointer(1000),
        EddystoneUrlUrl:         models.ToPointer("https://www.abc.com"),
        IbeaconAdvPower:         models.ToPointer(-65),
        IbeaconBeams:            models.ToPointer("2-4,7"),
        IbeaconEnabled:          models.ToPointer(false),
        IbeaconFreqMsec:         models.ToPointer(0),
        IbeaconMajor:            models.NewOptional(models.ToPointer(1234)),
        IbeaconMinor:            models.NewOptional(models.ToPointer(1234)),
        IbeaconUuid:             models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3")),
        Power:                   models.ToPointer(6),
        PowerMode:               models.ToPointer(models.BleConfigPowerModeEnum_CUSTOM),
    }

}
```

