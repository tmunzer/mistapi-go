
# Nac Client Last Status Enum

Latest Authentication status of the client. enum: `denied`, `permitted`, `session_started`, `session_stopped`

## Enumeration

`NacClientLastStatusEnum`

## Fields

| Name |
|  --- |
| `permitted` |
| `session_started` |
| `session_stopped` |
| `denied` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacClientLastStatus := models.NacClientLastStatusEnum_PERMITTED

}
```

