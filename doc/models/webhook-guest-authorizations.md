
# Webhook Guest Authorizations

Sample of the `guest-authorizations` webhook payload.

## Structure

`WebhookGuestAuthorizations`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookGuestAuthorizationsEvent`](../../doc/models/webhook-guest-authorizations-event.md) | Optional | Guest authorization events included in this webhook delivery |
| `Topic` | [`*models.WebhookGuestAuthorizationsTopicEnum`](../../doc/models/webhook-guest-authorizations-topic-enum.md) | Optional | Webhook topic name for guest authorization deliveries. enum: `guest-authorizations` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookGuestAuthorizations := models.WebhookGuestAuthorizations{
        Events:               []models.WebhookGuestAuthorizationsEvent{
            models.WebhookGuestAuthorizationsEvent{
                Ap:                     models.ToPointer("ap6"),
                AuthMethod:             models.ToPointer("auth_method4"),
                AuthorizedExpiringTime: models.ToPointer(42),
                AuthorizedTime:         models.ToPointer(252),
                Carrier:                models.ToPointer("carrier2"),
            },
            models.WebhookGuestAuthorizationsEvent{
                Ap:                     models.ToPointer("ap6"),
                AuthMethod:             models.ToPointer("auth_method4"),
                AuthorizedExpiringTime: models.ToPointer(42),
                AuthorizedTime:         models.ToPointer(252),
                Carrier:                models.ToPointer("carrier2"),
            },
            models.WebhookGuestAuthorizationsEvent{
                Ap:                     models.ToPointer("ap6"),
                AuthMethod:             models.ToPointer("auth_method4"),
                AuthorizedExpiringTime: models.ToPointer(42),
                AuthorizedTime:         models.ToPointer(252),
                Carrier:                models.ToPointer("carrier2"),
            },
        },
        Topic:                models.ToPointer(models.WebhookGuestAuthorizationsTopicEnum_GUESTAUTHORIZATIONS),
    }

}
```

