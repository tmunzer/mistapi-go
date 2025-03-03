
# Capture Wireless

Initiate a Wireless Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureWireless`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Band` | [`*models.CaptureWirelessBandEnum`](../../doc/models/capture-wireless-band-enum.md) | Optional | enum: `24`, `5`, `6`<br>**Default**: `"24"` |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureWirelessFormatEnum`](../../doc/models/capture-wireless-format-enum.md) | Optional | pcap format. enum: `pcap`, `stream`<br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br>**Default**: `1024`<br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `*string` | Optional | - |
| `Type` | `string` | Required, Constant | enum: `wireless`<br>**Value**: `"wireless"` |
| `WlanId` | `*uuid.UUID` | Optional | WLAN ID |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "24",
  "duration": 300,
  "format": "pcap",
  "max_pkt_len": 128,
  "num_packets": 1000,
  "type": "wireless",
  "ap_mac": "ap_mac4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

