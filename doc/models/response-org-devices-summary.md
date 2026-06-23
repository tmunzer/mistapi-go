
# Response Org Devices Summary

Organization device count summary

## Structure

`ResponseOrgDevicesSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `*int` | Optional, Read-only | Number of APs in the organization |
| `NumGateways` | `*int` | Optional, Read-only | Number of gateways in the organization |
| `NumMxedges` | `*int` | Optional, Read-only | Number of Mist Edges in the organization |
| `NumSwitches` | `*int` | Optional, Read-only | Number of switches in the organization |
| `NumUnassignedAps` | `*int` | Optional, Read-only | Number of APs in organization inventory that are not assigned to a site |
| `NumUnassignedGateways` | `*int` | Optional, Read-only | Number of gateways in organization inventory that are not assigned to a site |
| `NumUnassignedSwitches` | `*int` | Optional, Read-only | Number of switches in organization inventory that are not assigned to a site |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseOrgDevicesSummary := models.ResponseOrgDevicesSummary{
        NumAps:                models.ToPointer(86),
        NumGateways:           models.ToPointer(166),
        NumMxedges:            models.ToPointer(126),
        NumSwitches:           models.ToPointer(184),
        NumUnassignedAps:      models.ToPointer(88),
    }

}
```

