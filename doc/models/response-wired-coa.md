
# Response Wired Coa

Response returned after triggering wired client CoA reauthentication

## Structure

`ResponseWiredCoa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | MAC address of the switch or gateway hosting the wired client session |
| `PortId` | `*string` | Optional | Interface identifier for the wired client session |
| `Session` | `*uuid.UUID` | Optional | RADIUS session identifier used for the CoA request |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseWiredCoa := models.ResponseWiredCoa{
        DeviceMac:            models.ToPointer("5c5b35000002"),
        PortId:               models.ToPointer("ge-0/0/0"),
        Session:              models.ToPointer(uuid.MustParse("0a2a11b8-4b30-40d8-a6d1-e91ea540d86f")),
    }

}
```

