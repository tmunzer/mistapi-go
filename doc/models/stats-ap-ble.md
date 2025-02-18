
# Stats Ap Ble

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApBle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconEnabled` | `models.Optional[bool]` | Optional | - |
| `BeaconRate` | `models.Optional[int]` | Optional | - |
| `EddystoneUidEnabled` | `models.Optional[bool]` | Optional | - |
| `EddystoneUidFreqMsec` | `models.Optional[int]` | Optional | - |
| `EddystoneUidInstance` | `models.Optional[string]` | Optional | - |
| `EddystoneUidNamespace` | `models.Optional[string]` | Optional | - |
| `EddystoneUrlEnabled` | `models.Optional[bool]` | Optional | - |
| `EddystoneUrlFreqMsec` | `models.Optional[int]` | Optional | Frequency (msec) of data emmit by Eddystone-UID beacon |
| `EddystoneUrlUrl` | `models.Optional[string]` | Optional | - |
| `IbeaconEnabled` | `models.Optional[bool]` | Optional | - |
| `IbeaconFreqMsec` | `models.Optional[int]` | Optional | - |
| `IbeaconMajor` | `models.Optional[int]` | Optional | - |
| `IbeaconMinor` | `models.Optional[int]` | Optional | - |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | - |
| `Major` | `models.Optional[int]` | Optional | - |
| `Minors` | `[]int` | Optional | - |
| `Power` | `models.Optional[int]` | Optional | - |
| `RxBytes` | `models.Optional[int]` | Optional | - |
| `RxPkts` | `models.Optional[int]` | Optional | - |
| `TxBytes` | `models.Optional[int64]` | Optional | - |
| `TxPkts` | `models.Optional[int]` | Optional | - |
| `TxResets` | `models.Optional[int]` | Optional | Resets due to tx hung |
| `Uuid` | `models.Optional[uuid.UUID]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "beacon_rate": 3,
  "eddystone_uid_enabled": false,
  "eddystone_uid_freq_msec": 2000,
  "eddystone_uid_instance": "5c5b35000001",
  "eddystone_uid_namespace": "2818e3868dec25629ede",
  "eddystone_url_enabled": true,
  "eddystone_url_freq_msec": 100,
  "eddystone_url_url": "https://www.abc.com",
  "ibeacon_enabled": true,
  "ibeacon_freq_msec": 2000,
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
  "uuid": "ada72f8f-1643-e5c6-94db-f2a5636f1a64",
  "beacon_enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

