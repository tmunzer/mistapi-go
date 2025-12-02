
# Capture Client

Initiate a Client Packet Capture

## Structure

`CaptureClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `ClientMac` | `models.Optional[string]` | Optional | Client mac, required if `type`==`client`; optional otherwise |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `IncludesMcast` | `*bool` | Optional | **Default**: `false` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `models.Optional[string]` | Optional | Optional filter by ssid |
| `Type` | `string` | Required, Constant | enum: `client`<br><br>**Value**: `"client"` |

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

