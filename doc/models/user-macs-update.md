
# User Macs Update

Result of a bulk user MAC update

## Structure

`UserMacsUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Status message returned for asynchronous batch updates |
| `Errors` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Updated` | `[]uuid.UUID` | Optional | UUID string values used as identifiers |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    userMacsUpdate := models.UserMacsUpdate{
        Detail:               models.ToPointer("detail8"),
        Errors:               []string{
            "errors7",
            "errors8",
            "errors9",
        },
        Updated:              []uuid.UUID{
            uuid.MustParse("000001c7-0000-0000-0000-000000000000"),
            uuid.MustParse("000001c6-0000-0000-0000-000000000000"),
            uuid.MustParse("000001c5-0000-0000-0000-000000000000"),
        },
    }

}
```

