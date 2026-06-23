
# Wired Client Response Device Mac Port Item

Switch or gateway port observation for a wired client

## Structure

`WiredClientResponseDeviceMacPortItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | Switch or gateway MAC address for this wired client observation<br><br>**Constraints**: *Minimum Length*: `1` |
| `Ip` | `*string` | Optional, Read-only | Client IP address observed for this port entry |
| `PortId` | `*string` | Optional, Read-only | Interface identifier where the wired client was observed |
| `PortParent` | `*string` | Optional | Parent interface or port group associated with this port entry |
| `Start` | `*string` | Optional, Read-only | Time when this wired client observation began |
| `Vlan` | `*int` | Optional, Read-only | Client VLAN identifier observed for this port entry |
| `When` | `*string` | Optional, Read-only | Time when this wired client port entry was recorded |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wiredClientResponseDeviceMacPortItem := models.WiredClientResponseDeviceMacPortItem{
        DeviceMac:            models.ToPointer("device_mac0"),
        Ip:                   models.ToPointer("ip0"),
        PortId:               models.ToPointer("port_id6"),
        PortParent:           models.ToPointer("port_parent8"),
        Start:                models.ToPointer("start0"),
    }

}
```

