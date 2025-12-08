
# Wifi Beacon Extended Info Items

## Structure

`WifiBeaconExtendedInfoItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FrameCtrl` | `*int` | Optional | Frame control field of 802.11 header |
| `Payload` | `*string` | Optional | Extended Info Payload associated with frame |
| `SeqCtrl` | `*int` | Optional | Sequence control field of 802.11 header |

## Example (as JSON)

```json
{
  "frame_ctrl": 202,
  "payload": "payload8",
  "seq_ctrl": 114
}
```

