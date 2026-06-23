
# Webhook Device Updowns Event

Device availability event delivered when a monitored AP changes up/down state

## Structure

`WebhookDeviceUpdownsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required, Read-only | MAC address for the AP whose availability changed |
| `ApName` | `string` | Required, Read-only | AP name for the device whose availability changed |
| `ForSite` | `*bool` | Optional, Read-only | Whether this device up/down event is scoped to a site rather than only to the organization |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `SiteName` | `string` | Required, Read-only | Site name associated with the device up/down event |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required, Read-only | Reported up/down event type for the device |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookDeviceUpdownsEvent := models.WebhookDeviceUpdownsEvent{
        Ap:                   "ap6",
        ApName:               "ap_name8",
        ForSite:              models.ToPointer(false),
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        SiteName:             "site_name2",
        Timestamp:            float64(189.98),
        Type:                 "type0",
    }

}
```

