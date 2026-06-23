
# Websocket Session with Url

Response containing a WebSocket session handle and connection URL

## Structure

`WebsocketSessionWithUrl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Session` | `string` | Required | Identifier used to correlate output on the WebSocket stream |
| `Url` | `string` | Required | WebSocket URL returned for connecting to this session |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    websocketSessionWithUrl := models.WebsocketSessionWithUrl{
        Session:              "19e73828-937f-05e6-f709-e29efdb0a82b",
        Url:                  "wss://api-ws.mist.com/ssh?jwt=xxxx",
    }

}
```

