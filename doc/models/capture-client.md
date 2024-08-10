
# Capture Client

Initiate a Client Packet Capture

## Structure

`CaptureClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `ClientMac` | `models.Optional[string]` | Optional | client mac, required if `type`==`client`; optional otherwise |
| `Duration` | `models.Optional[int]` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `IncludesMcast` | `*bool` | Optional | **Default**: `false` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `128`<br>**Constraints**: `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024 for client-capture<br>**Default**: `1024` |
| `Ssid` | `models.Optional[string]` | Optional | optional filter by ssid |
| `Type` | `string` | Required, Constant | enum: `client`<br>**Default**: `"client"` |

## Example (as JSON)

```json
{
  "client_mac": "60a10a773412",
  "duration": 300,
  "includes_mcast": false,
  "max_pkt_len": 128,
  "num_packets": 1000,
  "type": "client",
  "ap_mac": "ap_mac0"
}
```

