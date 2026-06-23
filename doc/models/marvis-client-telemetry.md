
# Marvis Client Telemetry

Note: some stats are not collected when it's not connected to Mist infrastructure

## Structure

`MarvisClientTelemetry`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether telemetry collection is enabled for Marvis Client |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisClientTelemetry := models.MarvisClientTelemetry{
        Enabled:              models.ToPointer(false),
    }

}
```

