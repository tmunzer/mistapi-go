
# Ap Stats Ble Stat

## Structure

`ApStatsBleStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconRate` | `*int` | Optional | **Default**: `0` |
| `EddystoneUidEnabled` | `*bool` | Optional | **Default**: `false` |
| `EddystoneUidFreqMsec` | `*int` | Optional | **Default**: `0` |
| `EddystoneUidInstance` | `*string` | Optional | - |
| `EddystoneUidNamespace` | `*string` | Optional | - |
| `EddystoneUrlEnabled` | `*bool` | Optional | **Default**: `false` |
| `EddystoneUrlFreqMsec` | `*int` | Optional | Frequency (msec) of data emmit by Eddystone-UID beacon<br>**Default**: `0` |
| `EddystoneUrlUrl` | `*string` | Optional | - |
| `IbeaconEnabled` | `*bool` | Optional | **Default**: `false` |
| `IbeaconMajor` | `*int` | Optional | **Default**: `0` |
| `IbeaconMinor` | `*int` | Optional | **Default**: `0` |
| `IbeaconUuid` | `*uuid.UUID` | Optional | - |
| `Major` | `*int` | Optional | **Default**: `0` |
| `Minors` | `[]int` | Optional | - |
| `Power` | `*int` | Optional | **Default**: `9` |
| `RxBytes` | `*int` | Optional | **Default**: `0` |
| `RxPkts` | `*int` | Optional | **Default**: `0` |
| `TxBytes` | `*int64` | Optional | **Default**: `0` |
| `TxPkts` | `*int` | Optional | **Default**: `0` |
| `TxResets` | `*int` | Optional | resets due to tx hung<br>**Default**: `0` |
| `Uuid` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "beacon_rate": 3,
  "eddystone_uid_enabled": false,
  "eddystone_uid_freq_msec": 200,
  "eddystone_uid_instance": "5c5b35000001",
  "eddystone_uid_namespace": "2818e3868dec25629ede",
  "eddystone_url_enabled": true,
  "eddystone_url_freq_msec": 100,
  "eddystone_url_url": "https://www.abc.com",
  "ibeacon_enabled": true,
  "ibeacon_major": 13,
  "ibeacon_minor": 138,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "major": 12345,
  "power": 10,
  "rx_bytes": 135,
  "rx_pkts": 135,
  "tx_bytes": 5231513353,
  "tx_pkts": 135135135,
  "tx_resets": 0,
  "uuid": "ada72f8f-1643-e5c6-94db-f2a5636f1a64"
}
```

