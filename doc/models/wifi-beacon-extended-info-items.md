
# Wifi Beacon Extended Info Items

*This model accepts additional fields of type interface{}.*

## Structure

`WifiBeaconExtendedInfoItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FrameCtrl` | `*int` | Optional | Frame control field of 802.11 header |
| `Payload` | `*string` | Optional | Extended Info Payload associated with frame |
| `SeqCtrl` | `*int` | Optional | Sequence control field of 802.11 header |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "frame_ctrl": 202,
  "payload": "payload8",
  "seq_ctrl": 114,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

