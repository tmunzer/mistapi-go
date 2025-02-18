
# Response Pcap Search Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePcapSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMacs` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Aps` | `[]string` | Optional | - |
| `Duration` | `*float64` | Optional | - |
| `Format` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `MaxNumPackets` | `*float64` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PcapAps` | [`map[string]models.ResponsePcapSearchItemPcapApsItem`](../../doc/models/response-pcap-search-item-pcap-aps-item.md) | Optional | - |
| `PcapUrl` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TerminationReason` | `*string` | Optional | - |
| `Timestamp` | `float64` | Required | - |
| `Type` | `string` | Required | - |
| `Url` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 600.0,
  "format": "stream",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "max_num_packets": 1024,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "pcap_aps": {
    "5c5b35000010": {
      "band": "6",
      "bandwidth": "20",
      "channel": 133,
      "tcpdump_expression": null
    }
  },
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "termination_reason": "default",
  "timestamp": 93.82,
  "type": "type0",
  "url": "url4",
  "ap_macs": [
    "ap_macs7",
    "ap_macs8",
    "ap_macs9"
  ],
  "aps": [
    "aps1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

