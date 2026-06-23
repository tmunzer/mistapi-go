
# Alarm Ack

Request body for acknowledging or unacknowledging alarms

*This model accepts additional fields of type interface{}.*

## Structure

`AlarmAck`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmIds` | `[]uuid.UUID` | Required | Alarm identifiers included in an alarm acknowledgement request |
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
    alarmAck := models.AlarmAck{
        AlarmIds:             []uuid.UUID{
            uuid.MustParse("ccb8c94d-ca56-4075-932f-1f2ab444ff2c"),
            uuid.MustParse("98ff4a3d-ec9b-4138-a42e-54fc3335179d"),
        },
        Note:                 models.ToPointer("maintenance window"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

