
# Response Pcap Status

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePcapStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `models.Optional[string]` | Optional | - |
| `Aps` | `[]string` | Optional | List of target APs to capture packets |
| `ClientMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `Failed` | `[]string` | Optional | List of APs where configuration attempt failed |
| `Format` | [`*models.CaptureMxedgeFormatEnum`](../../doc/models/capture-mxedge-format-enum.md) | Optional | PCAP format. enum:<br><br>* `stream`: to Mist cloud<br>* `tzsp`: tream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)<br>**Default**: `"stream"` |
| `Gateways` | `[]string` | Optional | Information on gateways to capture packets on if a gateway capture type is specified |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `IncludesMcast` | `*bool` | Optional | - |
| `MaxNumPackets` | `*int` | Optional | Max number of packets configured by user |
| `MaxPktLen` | `*int` | Optional | - |
| `Mxedges` | `[]string` | Optional | nformation on mxedges to capture packets on if a mxedge capture type is specified |
| `NumPackets` | `*int` | Optional | total number of packets captured by all AP, not applicable for type [client, new_assoc] |
| `Ok` | `[]string` | Optional | List of target APs successfully configured to capture packets |
| `PcapAps` | [`map[string]models.ResponsePcapAp`](../../doc/models/response-pcap-ap.md) | Optional | - |
| `RadiotapTcpdumpExpression` | `*string` | Optional | When `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user |
| `ScanTcpdumpExpression` | `*string` | Optional | When `type`==`scan`, scan_tcpdump_expression provided by the user |
| `Ssid` | `models.Optional[string]` | Optional | - |
| `StartedTime` | `*int` | Optional | - |
| `Switches` | `[]string` | Optional | Information on switches to capture packets on if a switch capture type is specified. irb port interface is automatically added to capture as needed to ensure all desired packets are captured. |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression provided by the user (common) |
| `Type` | [`models.PcapTypeEnum`](../../doc/models/pcap-type-enum.md) | Required | enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless` |
| `TzspHost` | `*string` | Optional | Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets. |
| `TzspPort` | `*int` | Optional | If `format`==`tzsp`. Port on remote host for receiving the captured packets<br>**Constraints**: `>= 1`, `<= 65535` |
| `WiredTcpdumpExpression` | `*string` | Optional | When `type`==`wired`, wired_tcpdump_expression provided by the user |
| `WirelessTcpdumpExpression` | `*string` | Optional | When `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_mac": "60a10a773412",
  "duration": 300,
  "format": "stream",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "max_num_packets": 1000,
  "max_pkt_len": 128,
  "pcap_aps": {
    "5c5b35000010": {
      "band": 6,
      "bandwidth": 20,
      "channel": 133,
      "tcpdump_expressiin": null
    }
  },
  "started_time": 1435080709,
  "type": "client",
  "tzsp_host": "192.168.1.2",
  "ap_mac": "ap_mac4",
  "aps": [
    "aps3",
    "aps4",
    "aps5"
  ],
  "failed": [
    "failed6"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

