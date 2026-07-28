
# Zigbee Trail Response

Response containing the session identifier for a Zigbee event or packet trail operation

## Structure

`ZigbeeTrailResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Session` | `*uuid.UUID` | Optional | Session ID the UI can use to stream trail results |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    zigbeeTrailResponse := models.ZigbeeTrailResponse{
        Session:              models.ToPointer(uuid.MustParse("7a5f7796-83ee-11e5-95c6-1258369c38a9")),
    }

}
```

