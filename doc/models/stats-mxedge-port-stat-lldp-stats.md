
# Stats Mxedge Port Stat Lldp Stats

LLDP neighbor information reported for a Mist Edge port

## Structure

`StatsMxedgePortStatLldpStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisId` | `*string` | Optional | LLDP chassis identifier advertised by the neighbor |
| `MgmtAddr` | `*string` | Optional | Management address advertised by the LLDP neighbor |
| `PortDesc` | `*string` | Optional | Port description advertised by the LLDP neighbor |
| `PortId` | `*string` | Optional | Port identifier advertised by the LLDP neighbor |
| `SystemDesc` | `*string` | Optional | System description advertised by the LLDP neighbor |
| `SystemName` | `*string` | Optional | System name advertised by the LLDP neighbor |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgePortStatLldpStats := models.StatsMxedgePortStatLldpStats{
        ChassisId:            models.ToPointer("chassis_id0"),
        MgmtAddr:             models.ToPointer("mgmt_addr8"),
        PortDesc:             models.ToPointer("port_desc6"),
        PortId:               models.ToPointer("port_id6"),
        SystemDesc:           models.ToPointer("system_desc2"),
    }

}
```

