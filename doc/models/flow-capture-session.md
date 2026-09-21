
# Flow Capture Session

Current flow capture session status for a site

## Structure

`FlowCaptureSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CaptureFilter` | [`*models.FlowCaptureFilter`](../../doc/models/flow-capture-filter.md) | Optional | Normalized filter applied to a flow capture session |
| `Duration` | `*int` | Optional | - |
| `Enabled` | `*bool` | Optional | - |
| `Expiry` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `InvalidSwitches` | `*interface{}` | Optional | Switches that failed flow capture validation |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SwitchCount` | `*int` | Optional | - |
| `Timestamp` | `*int` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    flowCaptureSession := models.FlowCaptureSession{
        CaptureFilter:        models.ToPointer(models.FlowCaptureFilter{
            DstIp:                models.ToPointer("dst_ip6"),
            DstPort:              models.ToPointer(108),
            Protocol:             models.ToPointer(models.FlowCaptureProtocolEnum_TCP),
            SrcIp:                models.ToPointer("src_ip8"),
            SrcPort:              models.ToPointer(32),
        }),
        Duration:             models.ToPointer(110),
        Enabled:              models.ToPointer(false),
        Expiry:               models.ToPointer(190),
        Id:                   models.ToPointer(uuid.MustParse("00000a5e-0000-0000-0000-000000000000")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

