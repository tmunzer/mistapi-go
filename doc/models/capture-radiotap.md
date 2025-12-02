
# Capture Radiotap

Initiate a Radiotap Packet Capture

## Structure

`CaptureRadiotap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | - |
| `Band` | [`*models.CaptureRadiotapBandEnum`](../../doc/models/capture-radiotap-band-enum.md) | Optional | enum: `24`, `24,5,6`, `5`, `6`<br><br>**Default**: `"24"` |
| `ClientMac` | `*string` | Optional | - |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureRadiotapFormatEnum`](../../doc/models/capture-radiotap-format-enum.md) | Optional | enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `*string` | Optional | - |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump expression |
| `Type` | `string` | Required, Constant | enum: `radiotap`<br><br>**Value**: `"radiotap"` |
| `WlanId` | `*uuid.UUID` | Optional | WLAN id associated with the respective ssid. |

## Example (as JSON)

```json
{
  "ap_mac": "a83a79a947ee",
  "band": "24",
  "client_mac": "38f9d3972ff1",
  "duration": 300,
  "format": "stream",
  "max_pkt_len": 128,
  "num_packets": 1000,
  "ssid": "test",
  "tcpdump_expression": "tcp port 80",
  "type": "radiotap",
  "wlan_id": "fac8e973-feb9-421a-b381-aabbc4b61f5a"
}
```

