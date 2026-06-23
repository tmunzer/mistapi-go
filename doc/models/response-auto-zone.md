
# Response Auto Zone

Auto zones status and suggested zone response

## Structure

`ResponseAutoZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`*models.ResponseAutoZoneStatusEnum`](../../doc/models/response-auto-zone-status-enum.md) | Optional | Status of auto-zone generation for a map. enum: `in_progress`, `awaiting_review`, `not_started`, `error`. `not_started` means the service has not run or results were cleared, `in_progress` means generation is active, `awaiting_review` means suggested zones are ready for review, and `error` means generation failed |
| `Zones` | [`[]models.ResponseAutoZoneZone`](../../doc/models/response-auto-zone-zone.md) | Optional | Suggested zones returned by the auto zones service |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoZone := models.ResponseAutoZone{
        Status:               models.ToPointer(models.ResponseAutoZoneStatusEnum_INPROGRESS),
        Zones:                []models.ResponseAutoZoneZone{
            models.ResponseAutoZoneZone{
                Name:                 models.ToPointer("name8"),
                Vertices:             []models.ResponseAutoZoneZoneVertex{
                    models.ResponseAutoZoneZoneVertex{
                        X:                    models.ToPointer(16),
                        Y:                    models.ToPointer(88),
                    },
                    models.ResponseAutoZoneZoneVertex{
                        X:                    models.ToPointer(16),
                        Y:                    models.ToPointer(88),
                    },
                },
            },
        },
    }

}
```

