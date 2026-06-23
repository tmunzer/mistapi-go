
# Wlan Auth Server Selection Enum

When ordered, AP will prefer and go back to the first server if possible. enum: `ordered`, `unordered`

## Enumeration

`WlanAuthServerSelectionEnum`

## Fields

| Name |
|  --- |
| `ordered` |
| `unordered` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAuthServerSelection := models.WlanAuthServerSelectionEnum_ORDERED

}
```

