
# Wxtag Client

Client associated with a WxLAN tag

## Structure

`WxtagClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Client MAC address associated with the WxLAN tag entry |
| `Since` | `float64` | Required | Time when the client became associated with the WxLAN tag, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wxtagClient := models.WxtagClient{
        Mac:                  "5684dae9ac8b",
        Since:                float64(1428939600),
    }

}
```

