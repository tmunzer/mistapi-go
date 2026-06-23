
# Webhook Type Enum

enum: `aws-sns`, `google-pubsub`, `http-post`, `oauth2`, `splunk`

## Enumeration

`WebhookTypeEnum`

## Fields

| Name |
|  --- |
| `aws-sns` |
| `google-pubsub` |
| `http-post` |
| `oauth2` |
| `splunk` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookType := models.WebhookTypeEnum_GOOGLEPUBSUB

}
```

