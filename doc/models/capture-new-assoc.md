
# Capture New Assoc

Initiate a packet Capture for New Wireless Client Associations

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureNewAssoc`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | - |
| `ClientMac` | `*string` | Optional | client mac, required if `type`==`client`; optional otherwise |
| `Duration` | `*int` | Optional | duration of the capture, in seconds<br>**Default**: `600`<br>**Constraints**: `<= 86400` |
| `IncludesMcast` | `*bool` | Optional | **Default**: `false` |
| `MaxPktLen` | `*int` | Optional | **Default**: `128`<br>**Constraints**: `<= 2048` |
| `NumPackets` | `*int` | Optional | number of packets to capture, 0 for unlimited<br>**Default**: `100` |
| `Ssid` | `*string` | Optional | optional filter by ssid |
| `Type` | `string` | Required, Constant | enum: `new_assoc`<br>**Default**: `"new_assoc"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_mac": "a83a79a947ee",
  "client_mac": "60a10a773412",
  "duration": 600,
  "includes_mcast": false,
  "max_pkt_len": 128,
  "num_packets": 100,
  "type": "new_assoc",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

