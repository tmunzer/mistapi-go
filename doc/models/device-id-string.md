
# Device Id String

Request body containing a device identifier

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceIdString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceId` | `uuid.UUID` | Required | Device identifier supplied by the request |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    deviceIdString := models.DeviceIdString{
        DeviceId:             uuid.MustParse("00001bcc-0000-0000-0000-000000000000"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

