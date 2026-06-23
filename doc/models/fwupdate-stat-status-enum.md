
# Fwupdate Stat Status Enum

enum: `inprogress`, `failed`, `upgraded`, `success`, `scheduled`, `error`

## Enumeration

`FwupdateStatStatusEnum`

## Fields

| Name |
|  --- |
| `inprogress` |
| `failed` |
| `upgraded` |
| `success` |
| `scheduled` |
| `error` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    fwupdateStatStatus := models.FwupdateStatStatusEnum_INPROGRESS

}
```

