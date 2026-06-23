
# Psk Id List

PSK delete request payload

*This model accepts additional fields of type interface{}.*

## Structure

`PskIdList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PskIds` | `[]uuid.UUID` | Optional | PSK identifiers included in a delete request |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    pskIdList := models.PskIdList{
        PskIds:               []uuid.UUID{
            uuid.MustParse("0039c16c-ca87-4d3f-bb94-b97c58199f18"),
            uuid.MustParse("6562cc8e-5893-418a-acaa-4d7c1af8084f"),
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

