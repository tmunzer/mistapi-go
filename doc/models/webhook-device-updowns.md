
# Webhook Device Updowns

Sample of the `device-updowns` webhook payload.

## Structure

`WebhookDeviceUpdowns`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookDeviceUpdownsEvent`](../../doc/models/webhook-device-updowns-event.md) | Required | Device up/down events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for device up/down deliveries. enum: `device-updowns`<br><br>**Value**: `"device-updowns"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookDeviceUpdowns := models.WebhookDeviceUpdowns{
        Events:               []models.WebhookDeviceUpdownsEvent{
            models.WebhookDeviceUpdownsEvent{
                Ap:                   "ap6",
                ApName:               "ap_name8",
                ForSite:              models.ToPointer(false),
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                SiteName:             "site_name2",
                Timestamp:            float64(188.18),
                Type:                 "type0",
            },
        },
        Topic:                "device-updowns",
    }

}
```

