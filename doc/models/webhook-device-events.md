
# Webhook Device Events

Sample of the `device-events` webhook payload.

## Structure

`WebhookDeviceEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.DeviceEvent`](../../doc/models/device-event.md) | Required | List of device event payloads<br><br>**Constraints**: *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for device event deliveries. enum: `device-events`<br><br>**Value**: `"device-events"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookDeviceEvents := models.WebhookDeviceEvents{
        Events:               []models.DeviceEvent{
            models.DeviceEvent{
                Ap:                   models.ToPointer("ap6"),
                ApName:               models.ToPointer("ap_name8"),
                Apfw:                 models.ToPointer("apfw6"),
                AuditId:              models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                Bandwidth:            models.ToPointer(64),
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Timestamp:            float64(188.18),
                Type:                 "type0",
            },
        },
        Topic:                "device-events",
    }

}
```

