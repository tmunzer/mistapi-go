
# Response Pcap Status

## Structure

`ResponsePcapStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Aps` | `[]string` | Optional | List of target APs to capture packets |
| `ClientMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `Enabled` | `*bool` | Optional | - |
| `Expiry` | `*float64` | Optional | Expiry time of the capture session, in epoch seconds |
| `Failed` | `[]string` | Optional | List of APs where configuration attempt failed |
| `Format` | [`*models.CaptureMxedgeFormatEnum`](../../doc/models/capture-mxedge-format-enum.md) | Optional | PCAP format. enum:<br><br>* `stream`: to Mist cloud<br>* `tzsp`: stream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)<br><br>**Default**: `"stream"` |
| `Gateways` | `[]string` | Optional | Information on gateways to capture packets on if a gateway capture type is specified |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `IncludesMcast` | `*bool` | Optional | - |
| `InvalidMxedges` | `*interface{}` | Optional | Map of Mist Edge IDs that could not be configured for capture |
| `MaxNumPackets` | `*int` | Optional | Max number of packets configured by user |
| `MaxPktLen` | `*int` | Optional | - |
| `MxedgeCount` | `*int` | Optional | Number of Mist Edges in the capture session |
| `Mxedges` | [`map[string]models.ResponsePcapStatusMxedgesItem`](../../doc/models/response-pcap-status-mxedges-item.md) | Optional | Dict of Mist Edges to capture on, property key is the Mist Edge ID |
| `NumPackets` | `*int` | Optional | total number of packets captured by all AP, not applicable for type [client, new_assoc] |
| `Ok` | `[]string` | Optional | List of target APs successfully configured to capture packets |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PcapAps` | [`map[string]models.ResponsePcapAp`](../../doc/models/response-pcap-ap.md) | Optional | - |
| `RadiotapTcpdumpExpression` | `*string` | Optional | When `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user |
| `Raw` | `*bool` | Optional | - |
| `ScanTcpdumpExpression` | `*string` | Optional | When `type`==`scan`, scan_tcpdump_expression provided by the user |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `models.Optional[string]` | Optional | - |
| `StartedTime` | `*int` | Optional | - |
| `Switches` | `[]string` | Optional | Information on switches to capture packets on if a switch capture type is specified. irb port interface is automatically added to capture as needed to ensure all desired packets are captured. |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression provided by the user (common) |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Type` | [`models.PcapTypeEnum`](../../doc/models/pcap-type-enum.md) | Required | enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless` |
| `TzspHost` | `*string` | Optional | Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets. |
| `TzspPort` | `*int` | Optional | If `format`==`tzsp`. Port on remote host for receiving the captured packets<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `WiredTcpdumpExpression` | `*string` | Optional | When `type`==`wired`, wired_tcpdump_expression provided by the user |
| `WirelessTcpdumpExpression` | `*string` | Optional | When `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user |

## Example (as JSON)

```json
{
  "client_mac": "60a10a773412",
  "duration": 300,
  "expiry": 1695838060.309526,
  "format": "stream",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "max_num_packets": 1000,
  "max_pkt_len": 512,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "pcap_aps": {
    "5c5b35000010": {
      "band": 6,
      "bandwidth": 20,
      "channel": 133,
      "tcpdump_expression": null
    }
  },
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "started_time": 1435080709,
  "type": "client",
  "tzsp_host": "192.168.1.2",
  "ap_mac": "ap_mac4",
  "aps": [
    "aps3",
    "aps4",
    "aps5"
  ],
  "enabled": false
}
```

