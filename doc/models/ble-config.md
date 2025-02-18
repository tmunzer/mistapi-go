
# Ble Config

BLE AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`BleConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconEnabled` | `*bool` | Optional | Whether Mist beacons is enabled<br>**Default**: `false` |
| `BeaconRate` | `*int` | Optional | Required if `beacon_rate_mode`==`custom`, 1-10, in number-beacons-per-second<br>**Default**: `0` |
| `BeaconRateMode` | [`*models.BleConfigBeaconRateModeEnum`](../../doc/models/ble-config-beacon-rate-mode-enum.md) | Optional | enum: `custom`, `default`<br>**Default**: `"default"` |
| `BeamDisabled` | `[]int` | Optional | List of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam) |
| `CustomBlePacketEnabled` | `*bool` | Optional | Can be enabled if `beacon_enabled`==`true`, whether to send custom packet<br>**Default**: `false` |
| `CustomBlePacketFrame` | `*string` | Optional | The custom frame to be sent out in this beacon. The frame must be a hexstring |
| `CustomBlePacketFreqMsec` | `*int` | Optional | Frequency (msec) of data emitted by custom ble beacon<br>**Default**: `0`<br>**Constraints**: `>= 0` |
| `EddystoneUidAdvPower` | `*int` | Optional | Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default<br>**Default**: `0`<br>**Constraints**: `>= -100`, `<= 20` |
| `EddystoneUidBeams` | `*string` | Optional | - |
| `EddystoneUidEnabled` | `*bool` | Optional | Only if `beacon_enabled`==`false`, Whether Eddystone-UID beacon is enabled<br>**Default**: `false` |
| `EddystoneUidFreqMsec` | `*int` | Optional | Frequency (msec) of data emmit by Eddystone-UID beacon<br>**Default**: `0` |
| `EddystoneUidInstance` | `*string` | Optional | Eddystone-UID instance for the device |
| `EddystoneUidNamespace` | `*string` | Optional | Eddystone-UID namespace |
| `EddystoneUrlAdvPower` | `*int` | Optional | Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default<br>**Default**: `0`<br>**Constraints**: `>= -100`, `<= 20` |
| `EddystoneUrlBeams` | `*string` | Optional | - |
| `EddystoneUrlEnabled` | `*bool` | Optional | Only if `beacon_enabled`==`false`, Whether Eddystone-URL beacon is enabled<br>**Default**: `false` |
| `EddystoneUrlFreqMsec` | `*int` | Optional | Frequency (msec) of data emit by Eddystone-UID beacon<br>**Default**: `0` |
| `EddystoneUrlUrl` | `*string` | Optional | URL pointed by Eddystone-URL beacon |
| `IbeaconAdvPower` | `*int` | Optional | Advertised TX Power, -100 to 20 (dBm), omit this attribute to use default<br>**Default**: `0`<br>**Constraints**: `>= -100`, `<= 20` |
| `IbeaconBeams` | `*string` | Optional | - |
| `IbeaconEnabled` | `*bool` | Optional | Can be enabled if `beacon_enabled`==`true`, whether to send iBeacon<br>**Default**: `false` |
| `IbeaconFreqMsec` | `*int` | Optional | Frequency (msec) of data emmit for iBeacon<br>**Default**: `0` |
| `IbeaconMajor` | `*int` | Optional | Major number for iBeacon<br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `*int` | Optional | Minor number for iBeacon<br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `*uuid.UUID` | Optional | Optional, if not specified, the same UUID as the beacon will be used |
| `Power` | `*int` | Optional | Required if `power_mode`==`custom`; else use `power_mode` as default<br>**Constraints**: `>= 2`, `<= 7` |
| `PowerMode` | [`*models.BleConfigPowerModeEnum`](../../doc/models/ble-config-power-mode-enum.md) | Optional | enum: `custom`, `default`<br>**Default**: `"default"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "beacon_enabled": false,
  "beacon_rate": 3,
  "beacon_rate_mode": "custom",
  "beam_disabled": [
    1,
    3,
    6
  ],
  "custom_ble_packet_enabled": false,
  "custom_ble_packet_frame": "0x........",
  "custom_ble_packet_freq_msec": 300,
  "eddystone_uid_adv_power": -65,
  "eddystone_uid_beams": "2-4,7",
  "eddystone_uid_enabled": false,
  "eddystone_uid_freq_msec": 200,
  "eddystone_uid_instance": "5c5b35000001",
  "eddystone_uid_namespace": "2818e3868dec25629ede",
  "eddystone_url_adv_power": -65,
  "eddystone_url_beams": "2-4,7",
  "eddystone_url_enabled": false,
  "eddystone_url_freq_msec": 1000,
  "eddystone_url_url": "https://www.abc.com",
  "ibeacon_adv_power": -65,
  "ibeacon_beams": "2-4,7",
  "ibeacon_enabled": false,
  "ibeacon_freq_msec": 0,
  "ibeacon_major": 13,
  "ibeacon_minor": 138,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "power": 6,
  "power_mode": "custom",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

