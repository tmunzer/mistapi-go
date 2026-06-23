
# Alarms

Alarm identifiers and optional note used for bulk alarm state changes

*This model accepts additional fields of type interface{}.*

## Structure

`Alarms`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmIds` | `[]uuid.UUID` | Required | List of alarm identifiers |
| `Note` | `*string` | Optional | Some text note describing the intent |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    alarms := models.Alarms{
        AlarmIds:             []uuid.UUID{
            uuid.MustParse("00001a02-0000-0000-0000-000000000000"),
            uuid.MustParse("00001a03-0000-0000-0000-000000000000"),
            uuid.MustParse("00001a04-0000-0000-0000-000000000000"),
        },
        Note:                 models.ToPointer("note4"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

