
# Response Auto Zone Status Enum

Status of auto-zone generation for a map. enum: `in_progress`, `awaiting_review`, `not_started`, `error`. `not_started` means the service has not run or results were cleared, `in_progress` means generation is active, `awaiting_review` means suggested zones are ready for review, and `error` means generation failed

## Enumeration

`ResponseAutoZoneStatusEnum`

## Fields

| Name |
|  --- |
| `in_progress` |
| `awaiting_review` |
| `not_started` |
| `error` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoZoneStatus := models.ResponseAutoZoneStatusEnum_NOTSTARTED

}
```

