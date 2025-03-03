
# Capture New Assoc

Initiate a packet Capture for New Wireless Client Associations

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureNewAssoc`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | - |
| `ClientMac` | `*string` | Optional | Client mac, required if `type`==`client`; optional otherwise |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `>= 60`, `<= 86400` |
| `IncludesMcast` | `*bool` | Optional | **Default**: `false` |
| `MaxPktLen` | `models.Optional[int]` | Optional | **Default**: `512`<br>**Constraints**: `>= 64`, `<= 2048` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br>**Default**: `1024`<br>**Constraints**: `>= 0`, `<= 10000` |
| `Ssid` | `*string` | Optional | Optional filter by ssid |
| `Type` | `string` | Required, Constant | enum: `new_assoc`<br>**Value**: `"new_assoc"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_mac": "a83a79a947ee",
  "client_mac": "60a10a773412",
  "duration": 300,
  "includes_mcast": false,
  "max_pkt_len": 128,
  "num_packets": 1000,
  "type": "new_assoc",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

