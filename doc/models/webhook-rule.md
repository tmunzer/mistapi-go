
# Webhook Rule

Filtering rule that permits or blocks webhook events for a topic

## Structure

`WebhookRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.WebhookActionEnum`](../../doc/models/webhook-action-enum.md) | Optional | Webhook filtering action. enum: `permit`, `block`<br><br>**Default**: `"permit"` |
| `Matching` | `map[string][]string` | Optional | Event payload matching criteria. Property key is the event field name and the value is the list of accepted values (e.g. matching field `type` to [`AP_DISCONNECTED`]) |
| `Topic` | `string` | Required | Webhook topic this rule applies to |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookRule := models.WebhookRule{
        Action:               models.ToPointer(models.WebhookActionEnum_PERMIT),
        Matching:             map[string][]string{
            "type": map[string][]string{
                "": "",
            },
        },
        Topic:                "topic2",
    }

}
```

