
# Ap Port Config Dynamic Vlan Type Enum

Mapping mode for interpreting dynamic VLAN attributes returned by RADIUS.\ \ enum: `airespace-interface-name`, where the VLAN is determined by parsing\ \ the RADIUS attribute as an Airespace interface name (e.g. "guest"\ \ would map to VLAN 100), or `standard`, where the VLAN is determined by parsing\ \ the RADIUS attribute as a standard VLAN ID number

## Enumeration

`ApPortConfigDynamicVlanTypeEnum`

## Fields

| Name |
|  --- |
| `airespace-interface-name` |
| `standard` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apPortConfigDynamicVlanType := models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME

}
```

