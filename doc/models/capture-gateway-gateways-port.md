
# Capture Gateway Gateways Port

Per-port tcpdump filter for a gateway packet capture

## Structure

`CaptureGatewayGatewaysPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression per port |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureGatewayGatewaysPort := models.CaptureGatewayGatewaysPort{
        TcpdumpExpression:    models.ToPointer("tcpdump_expression2"),
    }

}
```

