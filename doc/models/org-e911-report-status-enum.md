
# Org E911 Report Status Enum

Current status of E911 report generation. enum: `disabled`, `scheduled`, `available`

## Enumeration

`OrgE911ReportStatusEnum`

## Fields

| Name |
|  --- |
| `disabled` |
| `scheduled` |
| `available` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgE911ReportStatus := models.OrgE911ReportStatusEnum_DISABLED

}
```

