
# Response Async License Status Enum

Processing state for an asynchronous license claim. enum: `prepared`, `ongoing`, `done`

## Enumeration

`ResponseAsyncLicenseStatusEnum`

## Fields

| Name |
|  --- |
| `prepared` |
| `ongoing` |
| `done` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAsyncLicenseStatus := models.ResponseAsyncLicenseStatusEnum_ONGOING

}
```

