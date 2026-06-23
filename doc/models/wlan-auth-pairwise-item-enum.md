
# Wlan Auth Pairwise Item Enum

enum: `wpa1-ccmp`, `wpa1-tkip`, `wpa2-ccmp`, `wpa2-tkip`, `wpa3`

## Enumeration

`WlanAuthPairwiseItemEnum`

## Fields

| Name |
|  --- |
| `wpa1-ccmp` |
| `wpa1-tkip` |
| `wpa2-ccmp` |
| `wpa2-tkip` |
| `wpa3` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAuthPairwiseItem := models.WlanAuthPairwiseItemEnum_WPA1CCMP

}
```

