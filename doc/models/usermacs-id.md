
# Usermacs Id

Request body for deleting multiple user MAC entries

*This model accepts additional fields of type interface{}.*

## Structure

`UsermacsId`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `UsermacIds` | `[]uuid.UUID` | Optional | UUID string values used as identifiers |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    usermacsId := models.UsermacsId{
        UsermacIds:           []uuid.UUID{
            uuid.MustParse("00000e83-0000-0000-0000-000000000000"),
            uuid.MustParse("00000e84-0000-0000-0000-000000000000"),
            uuid.MustParse("00000e85-0000-0000-0000-000000000000"),
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

