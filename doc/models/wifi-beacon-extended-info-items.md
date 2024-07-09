
# Wifi Beacon Extended Info Items

## Structure

`WifiBeaconExtendedInfoItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FrameCtrl` | `*int` | Optional | frame control field of 802.11 header |
| `Payload` | `*string` | Optional | Extended Info Payload associated with frame |
| `SeqCtrl` | `*int` | Optional | sequence control field of 802.11 header |

## Example (as JSON)

```json
{
  "frame_ctrl": 156,
  "payload": "payload0",
  "seq_ctrl": 68
}
```

