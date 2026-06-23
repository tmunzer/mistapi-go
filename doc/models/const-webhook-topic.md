
# Const Webhook Topic

Webhook topic definition returned by the constants API

## Structure

`ConstWebhookTopic`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowsSingleEventPerMessage` | `*bool` | Optional | Whether this topic can be configured to enforce one event per webhook message |
| `ForOrg` | `*bool` | Optional | Whether this topic can be used in organization-level webhooks |
| `HasDeliveryResults` | `*bool` | Optional | Whether delivery-result search is available for this webhook topic |
| `Internal` | `*bool` | Optional | Whether this topic is internal and not selectable in site or organization webhooks |
| `Key` | `*string` | Optional | Machine-readable key that identifies the webhook topic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constWebhookTopic := models.ConstWebhookTopic{
        AllowsSingleEventPerMessage: models.ToPointer(false),
        ForOrg:                      models.ToPointer(false),
        HasDeliveryResults:          models.ToPointer(false),
        Internal:                    models.ToPointer(false),
        Key:                         models.ToPointer("alarms"),
    }

}
```

