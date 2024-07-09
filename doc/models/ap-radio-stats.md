
# Ap Radio Stats

radio stat

## Structure

`ApRadioStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band * `80` is only applicable for band_5 and band_6 * `160` is only for band_6 |
| `Channel` | `*int` | Optional | current channel the radio is running on |
| `DynamicChainingEnalbed` | `*bool` | Optional | Use dynamic chaining for downlink |
| `Mac` | `*string` | Optional | radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af) |
| `NumClients` | `*int` | Optional | - |
| `Power` | `*int` | Optional | transmit power (in dBm) |
| `RxBytes` | `*float64` | Optional | - |
| `RxPkts` | `*float64` | Optional | - |
| `TxBytes` | `*float64` | Optional | - |
| `TxPkts` | `*float64` | Optional | - |
| `UtilAll` | `*int` | Optional | all utilization in percentage |
| `UtilNonWifi` | `*int` | Optional | reception of “No Packets” utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise |
| `UtilRxInBss` | `*int` | Optional | reception of “In BSS” utilization in percentage, only frames that are received from AP/STAs within the BSS |
| `UtilRxOtherBss` | `*int` | Optional | reception of “Other BSS” utilization in percentage, all frames received from AP/STAs that are outside the BSS |
| `UtilTx` | `*int` | Optional | transmission utilization in percentage |
| `UtilUndecodableWifi` | `*int` | Optional | reception of “UnDecodable Wifi“ utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio |
| `UtilUnknownWifi` | `*int` | Optional | reception of “No Category” utilization in percentage, all 802.11 frames that are corrupted at the receiver |

## Example (as JSON)

```json
{
  "bandwidth": 20,
  "channel": 190,
  "dynamic_chaining_enalbed": false,
  "mac": "mac4",
  "num_clients": 80
}
```

