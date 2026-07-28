
# Flow Record

Network flow record reported by a switch device

## Structure

`FlowRecord`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | MAC address of the device |
| `Direction` | [`*models.FlowRecordDirectionEnum`](../../doc/models/flow-record-direction-enum.md) | Optional | Flow direction. enum: `egress`, `ingress` |
| `DstIp` | `*string` | Optional | Destination IP address |
| `DstPort` | `*int` | Optional | Destination port number |
| `Duration` | `*int64` | Optional | Flow duration in seconds |
| `EndTime` | `*int64` | Optional | Flow end time in epoch seconds |
| `FlowId` | `*int64` | Optional | Unique flow identifier |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Protocol` | `*string` | Optional | Protocol (e.g. `tcp`, `udp`, `icmp`) |
| `SamplingPercentage` | `*float64` | Optional | Percentage of packets sampled (e.g. `0.1` means 0.1% of packets are captured via sFlow; `100.0` means all packets are captured via FBT) |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SrcIp` | `*string` | Optional | Source IP address |
| `SrcPort` | `*int` | Optional | Source port number |
| `StartTime` | `*int64` | Optional | Flow start time in epoch seconds |
| `State` | [`*models.FlowRecordStateEnum`](../../doc/models/flow-record-state-enum.md) | Optional | Flow state. enum: `active`, `aged-out` |
| `Timestamp` | `*int64` | Optional | Epoch time (in seconds) when the flow record was last updated or completed |
| `TotalBytes` | `*int64` | Optional | Total number of bytes in the flow |
| `TotalPkts` | `*int64` | Optional | Total number of packets in the flow |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    flowRecord := models.FlowRecord{
        DeviceMac:            models.ToPointer("device_mac2"),
        Direction:            models.ToPointer(models.FlowRecordDirectionEnum_EGRESS),
        DstIp:                models.ToPointer("dst_ip6"),
        DstPort:              models.ToPointer(10),
        Duration:             models.ToPointer(int64(196)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

