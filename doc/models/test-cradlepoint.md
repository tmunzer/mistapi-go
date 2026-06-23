
# Test Cradlepoint

Cradlepoint integration connectivity test result

## Structure

`TestCradlepoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `*string` | Optional, Read-only | if status is `inactive` this field returns the reason for it being inactive. |
| `LastStatus` | [`*models.TestCradlepointLastStatusEnum`](../../doc/models/test-cradlepoint-last-status-enum.md) | Optional, Read-only | status of integration detected during last sync. enum: `active`, `inactive` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    testCradlepoint := models.TestCradlepoint{
        Error:                models.ToPointer("Cradlepoint API keys are no longer valid, please verify and update the keys under organization settings."),
        LastStatus:           models.ToPointer(models.TestCradlepointLastStatusEnum_INACTIVE),
    }

}
```

