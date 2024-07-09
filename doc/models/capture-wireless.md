
# Capture Wireless

Initiate a Wireless Packet Capture

## Structure

`CaptureWireless`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Band` | [`*models.CaptureWirelessBandEnum`](../../doc/models/capture-wireless-band-enum.md) | Optional | **Default**: `"24"` |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `Format` | [`*models.CaptureWirelessFormatEnum`](../../doc/models/capture-wireless-format-enum.md) | Optional | pcap format<br>**Default**: `"pcap"` |
| `MaxPktLen` | `*int` | Optional | max_len of each packet to capture<br>**Default**: `128`<br>**Constraints**: `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `1024` |
| `Ssid` | `*string` | Optional | - |
| `Type` | `string` | Required, Constant | wireless<br>**Default**: `"wireless"` |
| `WlanId` | `*uuid.UUID` | Optional | WLAN ID |

## Example (as JSON)

```json
{
  "band": "24",
  "duration": 600,
  "format": "pcap",
  "max_pkt_len": 68,
  "num_packets": 100,
  "type": "wireless",
  "ap_mac": "ap_mac8"
}
```

