
# Avprofile Fallback Action Enum

Action applied when antivirus scanning cannot complete. enum: `block`, `log-and-permit`, `permit`

## Enumeration

`AvprofileFallbackActionEnum`

## Fields

| Name |
|  --- |
| `block` |
| `log-and-permit` |
| `permit` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    avprofileFallbackAction := models.AvprofileFallbackActionEnum_PERMIT

}
```

