
# Response Pcap Status

Current status of a packet capture session

## Structure

`ResponsePcapStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | AP MAC address targeted by the packet capture, or null when no single AP filter is set |
| `Aps` | `[]string` | Optional | List of target APs to capture packets |
| `ClientMac` | `models.Optional[string]` | Optional | Client MAC address filter applied to the packet capture, or null when no client filter is used |
| `Duration` | `*int` | Optional | Configured packet capture duration, in seconds |
| `Enabled` | `*bool` | Optional | Whether the packet capture session is currently enabled |
| `Expiry` | `*float64` | Optional | Epoch timestamp, in seconds, when the capture session expires |
| `Failed` | `[]string` | Optional | List of APs where configuration attempt failed |
| `Format` | [`*models.CaptureMxedgeFormatEnum`](../../doc/models/capture-mxedge-format-enum.md) | Optional | PCAP format. enum:<br><br>* `stream`: to Mist cloud<br>* `tzsp`: stream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)<br><br>**Default**: `"stream"` |
| `Gateways` | `[]string` | Optional | Information on gateways to capture packets on if a gateway capture type is specified |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `IncludesMcast` | `*bool` | Optional | Whether multicast traffic is included in the packet capture |
| `InvalidMxedges` | `*interface{}` | Optional | Map of Mist Edge IDs that could not be configured for capture |
| `MaxNumPackets` | `*int` | Optional | Max number of packets configured by user |
| `MaxPktLen` | `*int` | Optional | Maximum number of bytes captured from each packet |
| `MxedgeCount` | `*int` | Optional | Number of Mist Edges in the capture session |
| `Mxedges` | [`map[string]models.ResponsePcapStatusMxedgesItem`](../../doc/models/response-pcap-status-mxedges-item.md) | Optional | Dict of Mist Edges to capture on, property key is the Mist Edge ID |
| `NumPackets` | `*int` | Optional | total number of packets captured by all AP, not applicable for type [client, new_assoc] |
| `Ok` | `[]string` | Optional | List of target APs successfully configured to capture packets |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PcapAps` | [`map[string]models.ResponsePcapAp`](../../doc/models/response-pcap-ap.md) | Optional | Per-AP radio capture settings keyed by AP MAC address |
| `RadiotapTcpdumpExpression` | `*string` | Optional | When `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user |
| `Raw` | `*bool` | Optional | Whether raw packet data is included in the capture output |
| `ScanTcpdumpExpression` | `*string` | Optional | When `type`==`scan`, scan_tcpdump_expression provided by the user |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Ssid` | `models.Optional[string]` | Optional | Wireless network SSID filter applied to the packet capture, or null when no SSID filter is used |
| `StartedTime` | `*int` | Optional | Epoch timestamp, in seconds, when the capture session started |
| `Switches` | `[]string` | Optional | Information on switches to capture packets on if a switch capture type is specified. irb port interface is automatically added to capture as needed to ensure all desired packets are captured. |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression provided by the user (common) |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | [`models.PcapTypeEnum`](../../doc/models/pcap-type-enum.md) | Required | enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless` |
| `TzspHost` | `*string` | Optional | Required if `format`==`tzsp`. Remote host accessible to Mist Edges over the network for receiving the captured packets. |
| `TzspPort` | `*int` | Optional | If `format`==`tzsp`. Port on remote host for receiving the captured packets<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `WiredTcpdumpExpression` | `*string` | Optional | When `type`==`wired`, wired_tcpdump_expression provided by the user |
| `WirelessTcpdumpExpression` | `*string` | Optional | When `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePcapStatus := models.ResponsePcapStatus{
        ApMac:                     models.NewOptional(models.ToPointer("ap_mac6")),
        Aps:                       []string{
            "aps5",
            "aps6",
            "aps7",
        },
        ClientMac:                 models.NewOptional(models.ToPointer("60a10a773412")),
        Duration:                  models.ToPointer(300),
        Enabled:                   models.ToPointer(false),
        Expiry:                    models.ToPointer(float64(1695838060.309526)),
        Format:                    models.ToPointer(models.CaptureMxedgeFormatEnum_STREAM),
        Id:                        uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        MaxNumPackets:             models.ToPointer(1000),
        MaxPktLen:                 models.ToPointer(512),
        OrgId:                     models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PcapAps:                   map[string]models.ResponsePcapAp{
            "5c5b35000010": models.ResponsePcapAp{
                Band:                 models.ToPointer(6),
                Bandwidth:            models.ToPointer(20),
                Channel:              models.ToPointer(133),
                TcpdumpExpression:    models.NewOptional[string](nil),
            },
        },
        SiteId:                    models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        StartedTime:               models.ToPointer(1435080709),
        Type:                      models.PcapTypeEnum_CLIENT,
        TzspHost:                  models.ToPointer("192.168.1.2"),
    }

}
```

