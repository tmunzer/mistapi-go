
# Service Packet

Service data packet observed from an asset or beacon

## Structure

`ServicePacket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceData` | `*string` | Optional | ata from service data |
| `ServiceUuid` | `*string` | Optional | UUID from service data |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePacket := models.ServicePacket{
        ServiceData:          models.ToPointer("service_data2"),
        ServiceUuid:          models.ToPointer("service_uuid0"),
    }

}
```

