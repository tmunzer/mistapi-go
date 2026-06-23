
# Capture Gateway

Initiate a Gateway (SSR/SRX) Packet Capture

## Structure

`CaptureGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `models.Optional[int]` | Optional | Duration of the capture, in seconds<br><br>**Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 86400` |
| `Format` | [`*models.CaptureGatewayFormatEnum`](../../doc/models/capture-gateway-format-enum.md) | Optional | Output format for the gateway packet capture. enum: `stream`<br><br>**Default**: `"stream"` |
| `Gateways` | [`map[string]models.CaptureGatewayGateways`](../../doc/models/capture-gateway-gateways.md) | Required | List of SSRs. Property key is the SSR MAC |
| `MaxPktLen` | `models.Optional[int]` | Optional | minimum is 64 (SSR) / 68 (SRX) maximum is 10240 (SSR) / 1520 (SRX)<br><br>**Default**: `512`<br><br>**Constraints**: `>= 64`, `<= 1520` |
| `NumPackets` | `models.Optional[int]` | Optional | number of packets to capture, 0 for unlimited, default is 1024, maximum is 10000<br><br>**Default**: `1024`<br><br>**Constraints**: `>= 0`, `<= 10000` |
| `Ports` | [`map[string]models.CaptureGatewayGatewaysPort`](../../doc/models/capture-gateway-gateways-port.md) | Optional | Property key is the port ID |
| `Type` | `string` | Required, Constant | Packet capture type discriminator for gateway captures. enum: `gateway`<br><br>**Value**: `"gateway"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureGateway := models.CaptureGateway{
        Duration:             models.NewOptional(models.ToPointer(300)),
        Format:               models.ToPointer(models.CaptureGatewayFormatEnum_STREAM),
        Gateways:             map[string]models.CaptureGatewayGateways{
            "key0": models.CaptureGatewayGateways{
                Ports:                map[string]models.CaptureGatewayGatewaysPort{
                    "key0": models.CaptureGatewayGatewaysPort{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                    "key1": models.CaptureGatewayGatewaysPort{
                        TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
                    },
                },
            },
        },
        MaxPktLen:            models.NewOptional(models.ToPointer(128)),
        NumPackets:           models.NewOptional(models.ToPointer(1000)),
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
        Type:                 "gateway",
    }

}
```

