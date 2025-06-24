
# Capture Radiotapwired

Initiate a Radiotap Packet Capture and Wired Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureRadiotapwired`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Band` | [`*models.CaptureRadiotapwiredBandEnum`](../../doc/models/capture-radiotapwired-band-enum.md) | Optional | only used for radiotap. enum: `24`, `24,5,6`, `5`, `6`<br><br>**Default**: `"24"` |
| `ClientMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureRadiotapwiredFormatEnum`](../../doc/models/capture-radiotapwired-format-enum.md) | Optional | enum: `pcap`, `stream`<br><br>**Default**: `"pcap"` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `RadiotapTcpdumpExpression` | `*string` | Optional | tcpdump expression for radiotap interface (802.11 + radio headers) |
| `Ssid` | `models.Optional[string]` | Optional | - |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump expression |
| `Type` | `string` | Required, Constant | enum: `radiotap,wired`<br><br>**Value**: `"radiotap,wired"` |
| `WiredTcpdumpExpression` | `models.Optional[string]` | Optional | tcpdump expression |
| `WirelessTcpdumpExpression` | `*string` | Optional | tcpdump expression for radiotap interface (802.11) |
| `WlanId` | `models.Optional[string]` | Optional | WLAN id associated with the respective ssid. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "24",
  "client_mac": "38f9d3972ff1",
  "duration": 300,
  "format": "stream",
  "max_pkt_len": 128,
  "num_packets": 1000,
  "radiotap_tcpdump_expression": "type",
  "ssid": "test",
  "tcpdump_expression": "tcp port 80",
  "type": "radiotap,wired",
  "wired_tcpdump_expression": "tcp port 80",
  "wlan_id": "fac8e973-feb9-421a-b381-aabbc4b61f5a",
  "ap_mac": "ap_mac0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

