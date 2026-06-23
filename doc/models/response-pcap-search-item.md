
# Response Pcap Search Item

Packet capture record returned by organization or site packet capture search

## Structure

`ResponsePcapSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMacs` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Aps` | `[]string` | Optional | AP MAC addresses included in a packet capture |
| `Duration` | `*float64` | Optional | Packet capture duration in seconds |
| `Format` | `*string` | Optional | Output format requested for the packet capture |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `*float64` | Optional | Last seen timestamp of the capture |
| `MaxNumPackets` | `*float64` | Optional | Maximum number of packets requested for the capture |
| `Mxedges` | `[]string` | Optional | List of Mist Edge IDs included in the capture |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PcapAps` | [`map[string]models.ResponsePcapSearchItemPcapApsItem`](../../doc/models/response-pcap-search-item-pcap-aps-item.md) | Optional | Per-AP radio capture settings keyed by AP MAC address |
| `PcapUrl` | `*string` | Optional | URL for downloading the generated PCAP file |
| `SiteId` | `models.Optional[string]` | Optional | Site associated with the packet capture, when the capture is site-scoped |
| `TerminationReason` | `*string` | Optional | Reason the packet capture session ended |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required | Packet capture type represented by this record |
| `Url` | `string` | Required | Link for accessing the packet capture output or stream |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePcapSearchItem := models.ResponsePcapSearchItem{
        ApMacs:               []string{
            "ap_macs5",
            "ap_macs6",
        },
        Aps:                  []string{
            "aps7",
        },
        Duration:             models.ToPointer(float64(600)),
        Format:               models.ToPointer("stream"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        LastSeen:             models.ToPointer(float64(1693482149.417)),
        MaxNumPackets:        models.ToPointer(float64(1024)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PcapAps:              map[string]models.ResponsePcapSearchItemPcapApsItem{
            "5c5b35000010": models.ResponsePcapSearchItemPcapApsItem{
                Band:                 models.ToPointer("6"),
                Bandwidth:            models.ToPointer("20"),
                Channel:              models.ToPointer(133),
                TcpdumpExpression:    models.NewOptional[string](nil),
            },
        },
        TerminationReason:    models.ToPointer("default"),
        Timestamp:            float64(54.64),
        Type:                 "type4",
        Url:                  "url0",
    }

}
```

