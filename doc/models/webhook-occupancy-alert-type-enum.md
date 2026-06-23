
# Webhook Occupancy Alert Type Enum

Occupancy compliance state reported for the zone. enum: `COMPLIANCE-OK`, `COMPLIANCE-VIOLATION`

## Enumeration

`WebhookOccupancyAlertTypeEnum`

## Fields

| Name |
|  --- |
| `COMPLIANCE-OK` |
| `COMPLIANCE-VIOLATION` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookOccupancyAlertType := models.WebhookOccupancyAlertTypeEnum_COMPLIANCEOK

}
```

