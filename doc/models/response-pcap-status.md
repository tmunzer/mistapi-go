
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
| `Failed` | `[]string` | Optional | List of APs where configuration attempt failed |
| `Gateways` | `[]string` | Optional | List of target Gateways to capture packets if a gateway capture type is specified |
| `Id` | `uuid.UUID` | Required | unique id for the capture |
| `IncludesMcast` | `*bool` | Optional | - |
| `MaxNumPackets` | `*int` | Optional | max number of packets configured by user |
| `MaxPktLen` | `*int` | Optional | - |
| `NumPackets` | `*int` | Optional | total number of packets captured by all AP, not applicable for type [client, new_assoc] |
| `Ok` | `[]string` | Optional | List of target APs successfully configured to capture packets |
| `PcapAps` | [`map[string]models.ResponsePcapAp`](../../doc/models/response-pcap-ap.md) | Optional | - |
| `RadiotapTcpdumpExpression` | `*string` | Optional | when `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user |
| `ScanTcpdumpExpression` | `*string` | Optional | when `type`==`scan`, scan_tcpdump_expression provided by the user |
| `Ssid` | `models.Optional[string]` | Optional | - |
| `StartedTime` | `*int` | Optional | - |
| `Switches` | `[]string` | Optional | List of target Switches to capture packets if a switch capture type is specified |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression provided by the user (common) |
| `Type` | [`models.PcapTypeEnum`](../../doc/models/pcap-type-enum.md) | Required | - |
| `WiredTcpdumpExpression` | `*string` | Optional | when `type`==`wired`, wired_tcpdump_expression provided by the user |
| `WirelessTcpdumpExpression` | `*string` | Optional | when `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user |

## Example (as JSON)

```json
{
  "client_mac": "60a10a773412",
  "duration": 300,
  "id": "000001a4-0000-0000-0000-000000000000",
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
  "ap_mac": "ap_mac2",
  "aps": [
    "aps1",
    "aps2",
    "aps3"
  ],
  "failed": [
    "failed6"
  ]
}
```

