
# Mxedge Das

Cloud-assisted Dynamic Authorization Service settings for a Mist Edge cluster

## Structure

`MxedgeDas`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CoaServers` | [`[]models.MxedgeDasCoaServer`](../../doc/models/mxedge-das-coa-server.md) | Optional | Dynamic authorization clients configured to send CoA\|DM to mist edges on port 3799 |
| `Enabled` | `*bool` | Optional | Whether cloud-assisted DAS is enabled for the Mist Edge cluster<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeDas := models.MxedgeDas{
        CoaServers:           []models.MxedgeDasCoaServer{
            models.MxedgeDasCoaServer{
                DisableEventTimestampCheck:  models.ToPointer(false),
                Enabled:                     models.ToPointer(false),
                Host:                        models.ToPointer("host8"),
                Port:                        models.ToPointer(28),
                RequireMessageAuthenticator: models.ToPointer(false),
            },
            models.MxedgeDasCoaServer{
                DisableEventTimestampCheck:  models.ToPointer(false),
                Enabled:                     models.ToPointer(false),
                Host:                        models.ToPointer("host8"),
                Port:                        models.ToPointer(28),
                RequireMessageAuthenticator: models.ToPointer(false),
            },
            models.MxedgeDasCoaServer{
                DisableEventTimestampCheck:  models.ToPointer(false),
                Enabled:                     models.ToPointer(false),
                Host:                        models.ToPointer("host8"),
                Port:                        models.ToPointer(28),
                RequireMessageAuthenticator: models.ToPointer(false),
            },
        },
        Enabled:              models.ToPointer(false),
    }

}
```

