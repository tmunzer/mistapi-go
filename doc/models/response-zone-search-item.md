
# Response Zone Search Item

Zone visit record for a client, asset, or SDK client

## Structure

`ResponseZoneSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enter` | `*int` | Optional | Epoch timestamp, in seconds, when the zone visit started |
| `Scope` | `*string` | Optional | Zone scope represented by this visit record |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `User` | `*string` | Optional | Client, asset, or SDK client identifier for the visit |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseZoneSearchItem := models.ResponseZoneSearchItem{
        Enter:                models.ToPointer(1541705254),
        Scope:                models.ToPointer("map"),
        Timestamp:            models.ToPointer(float64(253.96)),
        User:                 models.ToPointer("c4b301c81166"),
    }

}
```

