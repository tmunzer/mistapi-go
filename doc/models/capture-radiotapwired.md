
# Capture Radiotapwired

Initiate a Radiotap Packet Capture and Wired Packet Capture

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureRadiotapwired`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Band` | [`*models.CaptureRadiotapwiredBandEnum`](../../doc/models/capture-radiotapwired-band-enum.md) | Optional | only used for radiotap. enum: `24`, `24,5,6`, `5`, `6`<br>**Default**: `"24"` |
| `ClientMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureRadiotapwiredFormatEnum`](../../doc/models/capture-radiotapwired-format-enum.md) | Optional | enum: `pcap`, `stream`<br>**Default**: `"pcap"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024`<br>**Constraints**: `>= 0` |
| `RadiotapTcpdumpExpression` | `*string` | Optional | tcpdump expression for radiotap interface (802.11 + radio headers) |
| `Ssid` | `models.Optional[string]` | Optional | - |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression common for wired,radiotap |
| `Type` | `string` | Required, Constant | enum: `radiotap,wired`<br>**Default**: `"radiotap,wired"` |
| `WiredTcpdumpExpression` | `*string` | Optional | tcpdump expression for wired |
| `WirelessTcpdumpExpression` | `*string` | Optional | tcpdump expression for radiotap interface (802.11) |
| `WlanId` | `models.Optional[string]` | Optional | wlan id associated with the respective ssid. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "24",
  "client_mac": "38f9d3972ff1",
  "duration": 600,
  "format": "stream",
  "max_pkt_len": 68,
  "num_packets": 100,
  "radiotap_tcpdump_expression": "type",
  "ssid": "atest",
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

