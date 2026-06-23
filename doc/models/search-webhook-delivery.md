
# Search Webhook Delivery

Paginated response for webhook delivery searches

## Structure

`SearchWebhookDelivery`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Upper bound timestamp of the webhook delivery search window, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of webhook delivery records returned by this page |
| `Next` | `*string` | Optional | URL for the next page of webhook delivery records, when more results are available |
| `Results` | [`[]models.WebhookDelivery`](../../doc/models/webhook-delivery.md) | Optional | Webhook delivery records returned by a search response |
| `Start` | `*int` | Optional | Lower bound timestamp of the webhook delivery search window, in epoch seconds |
| `Total` | `*int` | Optional | Count of webhook delivery records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    searchWebhookDelivery := models.SearchWebhookDelivery{
        End:                  models.ToPointer(1688035193),
        Limit:                models.ToPointer(10),
        Next:                 models.ToPointer("next0"),
        Results:              []models.WebhookDelivery{
            models.WebhookDelivery{
                Error:                models.ToPointer("error0"),
                Id:                   models.ToPointer(uuid.MustParse("000023ba-0000-0000-0000-000000000000")),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                ReqHeaders:           models.ToPointer("req_headers6"),
                ReqPayload:           models.ToPointer("req_payload4"),
            },
        },
        Start:                models.ToPointer(1687948793),
    }

}
```

