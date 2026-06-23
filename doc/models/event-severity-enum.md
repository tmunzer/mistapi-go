
# Event Severity Enum

Severity level for an event. enum: `normal`, `critical`, `high`, `warning`

## Enumeration

`EventSeverityEnum`

## Fields

| Name |
|  --- |
| `normal` |
| `critical` |
| `high` |
| `warning` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    eventSeverity := models.EventSeverityEnum_HIGH

}
```

