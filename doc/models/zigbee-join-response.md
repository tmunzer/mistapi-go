
# Zigbee Join Response

Response containing the session identifier for a Zigbee join operation

## Structure

`ZigbeeJoinResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SessionId` | `*uuid.UUID` | Optional | Session ID for the Zigbee join operation |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    zigbeeJoinResponse := models.ZigbeeJoinResponse{
        SessionId:            models.ToPointer(uuid.MustParse("19e73828-937f-05e6-f709-e29efdb0a82b")),
    }

}
```

