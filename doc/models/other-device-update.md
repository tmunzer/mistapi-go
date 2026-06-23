
# Other Device Update

Manual site association update for one other device

*This model accepts additional fields of type interface{}.*

## Structure

`OtherDeviceUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | Other-device MAC address to associate with a site |
| `SiteId` | `*uuid.UUID` | Optional | Site ID to associate with the other device |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    otherDeviceUpdate := models.OtherDeviceUpdate{
        DeviceMac:            models.ToPointer("device_mac2"),
        SiteId:               models.ToPointer(uuid.MustParse("43e9c864-a7e4-4310-8031-d9817d2c5a43")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

