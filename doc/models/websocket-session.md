
# Websocket Session

Response containing the WebSocket session handle for asynchronous command output

## Structure

`WebsocketSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Session` | `string` | Required | Identifier used to correlate output on the WebSocket stream |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    websocketSession := models.WebsocketSession{
        Session:              "19e73828-937f-05e6-f709-e29efdb0a82b",
    }

}
```

