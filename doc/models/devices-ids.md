
# Devices Ids

Request body containing device identifiers

*This model accepts additional fields of type interface{}.*

## Structure

`DevicesIds`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Required | Device identifiers to operate on in the request |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    devicesIds := models.DevicesIds{
        DeviceIds:            []uuid.UUID{
            uuid.MustParse("000021c5-0000-0000-0000-000000000000"),
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

