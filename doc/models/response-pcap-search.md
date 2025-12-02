
# Response Pcap Search

## Structure

`ResponsePcapSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponsePcapSearchItem`](../../doc/models/response-pcap-search-item.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 196,
  "limit": 230,
  "results": [
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
      "timestamp": 2.64,
      "type": "type4",
      "url": "url0",
      "ap_macs": [
        "ap_macs9",
        "ap_macs8"
      ],
      "aps": [
        "aps7",
        "aps8"
      ]
    }
  ],
  "start": 154,
  "next": "next6",
  "total": 68
}
```

