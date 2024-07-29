
# Capture Radiotap

Initiate a Radiotap Packet Capture

## Structure

`CaptureRadiotap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | - |
| `Band` | [`*models.CaptureRadiotapBandEnum`](../../doc/models/capture-radiotap-band-enum.md) | Optional | enum: `24`, `24,5,6`, `5`, `6`<br>**Default**: `"24"` |
| `ClientMac` | `*string` | Optional | - |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureRadiotapFormatEnum`](../../doc/models/capture-radiotap-format-enum.md) | Optional | enum: `pcap`, `stream`<br>**Default**: `"pcap"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024`<br>**Constraints**: `>= 0` |
| `Ssid` | `*string` | Optional | - |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression specific to radiotap |
| `Type` | `string` | Required, Constant | enum: `radiotap`<br>**Default**: `"radiotap"` |
| `WlanId` | `*uuid.UUID` | Optional | wlan id associated with the respective ssid. |

## Example (as JSON)

```json
{
  "ap_mac": "a83a79a947ee",
  "band": "24",
  "client_mac": "38f9d3972ff1",
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 68,
  "num_packets": 100,
  "ssid": "atest",
  "tcpdump_expression": "tcp port 80",
  "type": "radiotap",
  "wlan_id": "fac8e973-feb9-421a-b381-aabbc4b61f5a"
}
```

