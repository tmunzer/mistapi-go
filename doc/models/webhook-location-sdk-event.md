
# Webhook Location Sdk Event

SDK client location update with map coordinates

## Structure

`WebhookLocationSdkEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MapId` | `*uuid.UUID` | Optional | Map where the SDK client location was calculated |
| `Name` | `*string` | Optional | Display name of the SDK client when available |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | Location object type for the SDK client event; defaults to `sdk`<br><br>**Default**: `"sdk"` |
| `X` | `*float64` | Optional | Horizontal map coordinate of the SDK client, in meters |
| `Y` | `*float64` | Optional | Vertical map coordinate of the SDK client, in meters |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationSdkEvent := models.WebhookLocationSdkEvent{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MapId:                models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
        Name:                 models.ToPointer("optional"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Timestamp:            models.ToPointer(float64(176.26)),
        Type:                 models.ToPointer("sdk"),
        X:                    models.ToPointer(float64(13.5)),
        Y:                    models.ToPointer(float64(3.2)),
    }

}
```

