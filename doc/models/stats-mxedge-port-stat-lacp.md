
# Stats Mxedge Port Stat Lacp

LACP state and counters for a Mist Edge port

## Structure

`StatsMxedgePortStatLacp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MuxState` | `*string` | Optional | LACP multiplexer state reported for the port |
| `RxLacpdu` | `*int` | Optional | Number of LACPDUs received on the port |
| `RxState` | `*string` | Optional | LACP receive state reported for the port |
| `TxLacpdu` | `*int` | Optional | Number of LACPDUs transmitted from the port |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgePortStatLacp := models.StatsMxedgePortStatLacp{
        MuxState:             models.ToPointer("mux_state0"),
        RxLacpdu:             models.ToPointer(226),
        RxState:              models.ToPointer("rx_state6"),
        TxLacpdu:             models.ToPointer(252),
    }

}
```

