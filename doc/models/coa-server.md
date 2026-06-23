
# Coa Server

RADIUS Change of Authorization (CoA) server settings

## Structure

`CoaServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableEventTimestampCheck` | `*bool` | Optional | Whether to disable Event-Timestamp Check<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | Whether this RADIUS CoA server is enabled<br><br>**Default**: `false` |
| `Ip` | `string` | Required | Server IPv4 address for RADIUS CoA messages |
| `Port` | [`*models.RadiusCoaPort`](../../doc/models/containers/radius-coa-port.md) | Optional | RADIUS CoA Port, value from 1 to 65535, default is 3799 |
| `Secret` | `string` | Required | Shared secret used to authenticate RADIUS CoA messages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    coaServer := models.CoaServer{
        DisableEventTimestampCheck: models.ToPointer(false),
        Enabled:                    models.ToPointer(false),
        Ip:                         "1.2.3.4",
        Port:                       models.ToPointer(models.RadiusCoaPortContainer.FromNumber(206)),
        Secret:                     "testing456",
    }

}
```

