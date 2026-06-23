
# Const Event

Device event definition returned by the constants API

## Structure

`ConstEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Narrative description of the device event type |
| `Display` | `string` | Required | Human-readable label for the device event type |
| `Example` | `*interface{}` | Optional | Sample device event payload for this event type |
| `Group` | `*string` | Optional | Device event group for this definition |
| `Key` | `string` | Required | Machine-readable device event key |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constEvent := models.ConstEvent{
        Description:          models.ToPointer("description2"),
        Display:              "display4",
        Example:              models.ToPointer(interface{}("[key1, val1][key2, val2]")),
        Group:                models.ToPointer("group0"),
        Key:                  "key2",
    }

}
```

