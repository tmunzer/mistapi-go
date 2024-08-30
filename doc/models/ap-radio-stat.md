
# Ap Radio Stat

radio stat

## Structure

`ApRadioStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `models.Optional[int]` | Optional | current channel the radio is running on |
| `DynamicChainingEnalbed` | `models.Optional[bool]` | Optional | Use dynamic chaining for downlink |
| `Mac` | `models.Optional[string]` | Optional | radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af) |
| `NoiseFloor` | `models.Optional[int]` | Optional | - |
| `NumClients` | `models.Optional[int]` | Optional | - |
| `Power` | `models.Optional[int]` | Optional | transmit power (in dBm) |
| `RxBytes` | `models.Optional[int]` | Optional | - |
| `RxPkts` | `models.Optional[int]` | Optional | - |
| `TxBytes` | `models.Optional[int]` | Optional | - |
| `TxPkts` | `models.Optional[int]` | Optional | - |
| `Usage` | `models.Optional[string]` | Optional | - |
| `UtilAll` | `models.Optional[int]` | Optional | all utilization in percentage |
| `UtilNonWifi` | `models.Optional[int]` | Optional | reception of “No Packets” utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise |
| `UtilRxInBss` | `models.Optional[int]` | Optional | reception of “In BSS” utilization in percentage, only frames that are received from AP/STAs within the BSS |
| `UtilRxOtherBss` | `models.Optional[int]` | Optional | reception of “Other BSS” utilization in percentage, all frames received from AP/STAs that are outside the BSS |
| `UtilTx` | `models.Optional[int]` | Optional | transmission utilization in percentage |
| `UtilUndecodableWifi` | `models.Optional[int]` | Optional | reception of “UnDecodable Wifi“ utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio |
| `UtilUnknownWifi` | `models.Optional[int]` | Optional | reception of “No Category” utilization in percentage, all 802.11 frames that are corrupted at the receiver |

## Example (as JSON)

```json
{
  "bandwidth": 20,
  "noise_floor": -90,
  "usage": "24",
  "channel": 138,
  "dynamic_chaining_enalbed": false,
  "mac": "mac2"
}
```
