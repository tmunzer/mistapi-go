
# Response Search Item

MSP search result entry

## Structure

`ResponseSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Text` | `string` | Required | Display text for the MSP search result |
| `Type` | `string` | Required | Category of MSP search result, such as `orgs` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSearchItem := models.ResponseSearchItem{
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Text:                 "text4",
        Type:                 "type4",
    }

}
```

