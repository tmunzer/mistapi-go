
# Webhook Minis Application

Sample of the `minis-application` webhook payload.

## Structure

`WebhookMinisApplication`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookMinisApplicationEvent`](../../doc/models/webhook-minis-application-event.md) | Required | Marvis Minis application test events included in a webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for Minis application test deliveries. enum: `minis-application`<br><br>**Value**: `"minis-application"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookMinisApplication := models.WebhookMinisApplication{
        Events:               []models.WebhookMinisApplicationEvent{
            models.WebhookMinisApplicationEvent{
                DeviceMac:            models.ToPointer("device_mac4"),
                Ip:                   models.ToPointer("ip4"),
                Latency:              models.ToPointer(60),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                ProbeName:            models.ToPointer("connectivitycheck.gstatic.com"),
                ProbeType:            models.ToPointer("application"),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                TestType:             models.ToPointer(models.SynthetictestConfigCustomProbeTypeEnum_ICMP),
                Vlan:                 models.ToPointer(12),
            },
        },
        Topic:                "minis-application",
    }

}
```

