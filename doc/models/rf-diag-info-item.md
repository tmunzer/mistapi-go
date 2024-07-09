
# Rf Diag Info Item

## Structure

`RfDiagInfoItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetId` | `*uuid.UUID` | Optional | if `type`==`asset`, id of the asset |
| `AssetName` | `*string` | Optional | if `type`==`asset`, name of the asset |
| `ClientName` | `*string` | Optional | if `type`==`client`, hostname of the client |
| `Duration` | `int` | Required | recording length in seconds, max is 120 |
| `EndTime` | `int` | Required | timestamp of end of recording |
| `FrameCount` | `int` | Required | Number of frames in the output |
| `Id` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | if `type`==`client` or `asset`, mac of the device |
| `MapId` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `Next` | `*string` | Optional | Optional. id of the next recoding if present. Only valid for site survey. |
| `RawEvents` | `string` | Required | URL to a JSON file that contains array of raw location diag events |
| `Ready` | `bool` | Required | whether it’s ready for playback |
| `SdkclientId` | `*uuid.UUID` | Optional | if `type`==`sdkclient`, sdkclient_id of this recording |
| `SdkclientName` | `*string` | Optional | if `type`==`sdkclient`, name of the sdkclient |
| `SdkclientUuid` | `*uuid.UUID` | Optional | if `type`==`sdkclient`, device_id of sdkclient |
| `StartTime` | `int` | Required | timestamp of the recording (the start) |
| `Type` | [`models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Required | - |
| `Url` | `string` | Required | URL to a JSON file that contains an array of frames, each frame is the same format |

## Example (as JSON)

```json
{
  "asset_id": "00000968-0000-0000-0000-000000000000",
  "asset_name": "asset_name0",
  "client_name": "client_name6",
  "duration": 102,
  "end_time": 86,
  "frame_count": 12,
  "id": "id2",
  "mac": "mac6",
  "map_id": "000000da-0000-0000-0000-000000000000",
  "name": "name2",
  "raw_events": "raw_events4",
  "ready": false,
  "start_time": 126,
  "type": "asset",
  "url": "url6"
}
```
