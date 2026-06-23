
# Capture Switch Switches

Switch-specific packet capture settings keyed under a switch MAC address

## Structure

`CaptureSwitchSwitches`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.CaptureSwitchPortsTcpdumpExpression`](../../doc/models/capture-switch-ports-tcpdump-expression.md) | Optional | Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureSwitchSwitches := models.CaptureSwitchSwitches{
        Ports:                map[string]models.CaptureSwitchPortsTcpdumpExpression{
            "key0": models.CaptureSwitchPortsTcpdumpExpression{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
            },
            "key1": models.CaptureSwitchPortsTcpdumpExpression{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression0"),
            },
        },
    }

}
```

