
# Response Org Devices

Response containing organization device records

## Structure

`ResponseOrgDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.OrgDevice`](../../doc/models/org-device.md) | Required | List of organization device identifiers<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseOrgDevices := models.ResponseOrgDevices{
        Results:              []models.OrgDevice{
            models.OrgDevice{
                Mac:                  "mac0",
                Name:                 "name6",
            },
        },
    }

}
```

