
# Rf Diag Info Item

RF diagnostic recording metadata and output links

## Structure

`RfDiagInfoItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetId` | `*uuid.UUID` | Optional | If `type`==`asset`, id of the asset |
| `AssetName` | `*string` | Optional | If `type`==`asset`, name of the asset |
| `ClientName` | `*string` | Optional | If `type`==`client`, hostname of the client |
| `Duration` | `int` | Required | recording length in seconds, max is 120 |
| `EndTime` | `int` | Required | Timestamp of end of recording |
| `FrameCount` | `int` | Required | Number of frames in the output |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | If `type`==`client` or `asset`, MAC address of the recorded device |
| `MapId` | `uuid.UUID` | Required | Map identifier associated with the recording |
| `Name` | `string` | Required | Recording name returned for the RF diagnostic recording |
| `Next` | `*string` | Optional | Optional. id of the next recoding if present. Only valid for site survey. |
| `RawEvents` | `string` | Required | URL to a JSON file that contains array of raw location diag events |
| `Ready` | `bool` | Required | Whether it’s ready for playback |
| `SdkclientId` | `*uuid.UUID` | Optional | If `type`==`sdkclient`, sdkclient_id of this recording |
| `SdkclientName` | `*string` | Optional | If `type`==`sdkclient`, name of the sdkclient |
| `SdkclientUuid` | `*uuid.UUID` | Optional | If `type`==`sdkclient`, device_id of sdkclient |
| `StartTime` | `int` | Required | Timestamp of the recording (the start) |
| `Type` | [`models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Required | enum: `asset`, `client`, `sdkclient` |
| `Url` | `string` | Required | JSON file URL for frame data captured by the recording |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    rfDiagInfoItem := models.RfDiagInfoItem{
        AssetId:              models.ToPointer(uuid.MustParse("0000020a-0000-0000-0000-000000000000")),
        AssetName:            models.ToPointer("asset_name0"),
        ClientName:           models.ToPointer("client_name4"),
        Duration:             180,
        EndTime:              164,
        FrameCount:           90,
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Mac:                  models.ToPointer("mac6"),
        MapId:                uuid.MustParse("00000a98-0000-0000-0000-000000000000"),
        Name:                 "name2",
        RawEvents:            "raw_events4",
        Ready:                false,
        StartTime:            52,
        Type:                 models.RfClientTypeEnum_CLIENT,
        Url:                  "url6",
    }

}
```

