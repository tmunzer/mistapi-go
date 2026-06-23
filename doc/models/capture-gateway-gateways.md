
# Capture Gateway Gateways

Gateway-specific packet capture settings keyed under a gateway MAC address

## Structure

`CaptureGatewayGateways`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.CaptureGatewayGatewaysPort`](../../doc/models/capture-gateway-gateways-port.md) | Optional | Property key is the port ID |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureGatewayGateways := models.CaptureGatewayGateways{
        Ports:                map[string]models.CaptureGatewayGatewaysPort{
            "key0": models.CaptureGatewayGatewaysPort{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
            },
            "key1": models.CaptureGatewayGatewaysPort{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
            },
            "key2": models.CaptureGatewayGatewaysPort{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
            },
        },
    }

}
```

