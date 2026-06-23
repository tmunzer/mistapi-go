
# Alarm Template Rule

Per-alarm enablement and delivery override in an alarm template

## Structure

`AlarmTemplateRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Delivery` | [`*models.Delivery`](../../doc/models/delivery.md) | Optional | Delivery object to configure the alarm delivery |
| `Enabled` | `*bool` | Optional | Whether this alarm rule is enabled in the template |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    alarmTemplateRule := models.AlarmTemplateRule{
        Delivery:             models.ToPointer(models.Delivery{
            AdditionalEmails:     []string{
                "additional_emails9",
                "additional_emails0",
                "additional_emails1",
            },
            Enabled:              false,
            ToOrgAdmins:          models.ToPointer(false),
            ToSiteAdmins:         models.ToPointer(false),
        }),
        Enabled:              models.ToPointer(false),
    }

}
```

