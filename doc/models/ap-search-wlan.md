
# Ap Search Wlan

WLAN summary included in an AP search result

## Structure

`ApSearchWlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | WLAN identifier included in the AP search result |
| `Ssid` | `*string` | Optional | Wireless network SSID shown for this WLAN |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    apSearchWlan := models.ApSearchWlan{
        Id:                   models.ToPointer(uuid.MustParse("00000b1a-0000-0000-0000-000000000000")),
        Ssid:                 models.ToPointer("ssid0"),
    }

}
```

