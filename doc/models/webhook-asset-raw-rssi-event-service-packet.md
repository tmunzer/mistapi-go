
# Webhook Asset Raw Rssi Event Service Packet

BLE service data decoded from an asset raw RSSI packet

## Structure

`WebhookAssetRawRssiEventServicePacket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceData` | `models.Optional[string]` | Optional | optional, data from service data |
| `ServiceUuid` | `models.Optional[string]` | Optional | optional, UUID from service data |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookAssetRawRssiEventServicePacket := models.WebhookAssetRawRssiEventServicePacket{
        ServiceData:          models.NewOptional(models.ToPointer("service_data2")),
        ServiceUuid:          models.NewOptional(models.ToPointer("service_uuid0")),
    }

}
```

