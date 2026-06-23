
# Response Upgrade Id

Response containing an upgrade job identifier

## Structure

`ResponseUpgradeId`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `UpgradeId` | `uuid.UUID` | Required | Unique value identifying the queued upgrade job |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseUpgradeId := models.ResponseUpgradeId{
        UpgradeId:            uuid.MustParse("4316c116-0acb-4c43-8f06-6723154e741e"),
    }

}
```

