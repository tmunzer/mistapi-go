
# User Mac Import

Result of importing user MAC entries

## Structure

`UserMacImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Added` | `[]string` | Optional | MAC addresses added during a user MAC import |
| `Detail` | `*string` | Optional | Status message returned for asynchronous imports |
| `Errors` | `[]string` | Optional | Error messages returned during a user MAC import |
| `Updated` | `[]string` | Optional | MAC addresses updated during a user MAC import |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    userMacImport := models.UserMacImport{
        Added:                []string{
            "921b638445cd",
        },
        Detail:               models.ToPointer("detail8"),
        Errors:               []string{
            "921b638445ce - mac invalid",
            "921b638445cf - mac already provided",
        },
        Updated:              []string{
            "721b638445ef",
            "721b638445ee",
        },
    }

}
```

