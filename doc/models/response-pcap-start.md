
# Response Pcap Start

Packet capture session created by a start request

## Structure

`ResponsePcapStart`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApCount` | `*int` | Optional | Number of APs targeted by the packet capture |
| `Aps` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `ClientMac` | `models.Optional[string]` | Optional | Client MAC address filter applied to the packet capture, or null when no client filter is used |
| `Duration` | `*float64` | Optional | Packet capture duration in seconds |
| `Enabled` | `*bool` | Optional | Whether the packet capture session is enabled after the start request |
| `Expiry` | `*float64` | Optional | Epoch timestamp, in seconds, when the capture session expires |
| `Format` | `*string` | Optional | Output format for packet capture data |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `IncludeMcast` | `*bool` | Optional | Whether multicast traffic is included in the packet capture |
| `MaxPktLen` | `*int` | Optional | Maximum number of bytes captured from each packet |
| `NumPackets` | `*int` | Optional | Maximum number of packets to capture; use 0 for unlimited |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `Raw` | `*bool` | Optional | Whether raw packet data is included in the capture output |
| `SiteId` | `*string` | Required | Site associated with the packet capture session, when site-scoped |
| `Ssid` | `models.Optional[string]` | Optional | Wireless network SSID filter applied to the packet capture, or null when no SSID filter is used |
| `TcpdumpParserExpression` | `models.Optional[string]` | Optional | Tcpdump parser expression applied to the packet capture, or null when no parser expression is used |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required | Packet capture type requested by the start operation |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePcapStart := models.ResponsePcapStart{
        ApCount:                 models.ToPointer(232),
        Aps:                     []string{
            "aps5",
            "aps6",
        },
        ClientMac:               models.NewOptional(models.ToPointer("client_mac2")),
        Duration:                models.ToPointer(float64(155.8)),
        Enabled:                 models.ToPointer(false),
        Id:                      uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        OrgId:                   uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:                  models.ToPointer("site_id0"),
        Timestamp:               float64(54.02),
        Type:                    "type6",
    }

}
```

