
# Capture Switch Ports Tcpdump Expression

Per-port tcpdump filter for a switch packet capture

## Structure

`CaptureSwitchPortsTcpdumpExpression`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureSwitchPortsTcpdumpExpression := models.CaptureSwitchPortsTcpdumpExpression{
        TcpdumpExpression:    models.ToPointer("port 443"),
    }

}
```

