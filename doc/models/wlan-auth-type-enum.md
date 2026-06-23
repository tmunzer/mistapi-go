
# Wlan Auth Type Enum

enum: `eap`, `eap192`, `open`, `psk`, `psk-tkip`, `psk-wpa2-tkip`, `wep`

## Enumeration

`WlanAuthTypeEnum`

## Fields

| Name |
|  --- |
| `eap` |
| `eap192` |
| `open` |
| `psk` |
| `psk-tkip` |
| `psk-wpa2-tkip` |
| `wep` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAuthType := models.WlanAuthTypeEnum_EAP192

}
```

