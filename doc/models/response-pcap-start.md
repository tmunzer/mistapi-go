
# Response Pcap Start

## Structure

`ResponsePcapStart`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApCount` | `*int` | Optional | - |
| `Aps` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `ClientMac` | `models.Optional[string]` | Optional | - |
| `Duration` | `*float64` | Optional | - |
| `Enabled` | `*bool` | Optional | - |
| `Expiry` | `*float64` | Optional | - |
| `Format` | `*string` | Optional | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `IncludeMcast` | `*bool` | Optional | - |
| `MaxPktLen` | `*int` | Optional | - |
| `NumPackets` | `*int` | Optional | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `Raw` | `*bool` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Ssid` | `models.Optional[string]` | Optional | - |
| `TcpdumpParserExpression` | `models.Optional[string]` | Optional | - |
| `Timestamp` | `float64` | Required | - |
| `Type` | `string` | Required | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 1.12,
  "type": "type4",
  "ap_count": 62,
  "aps": [
    "aps5",
    "aps6"
  ],
  "client_mac": "client_mac2",
  "duration": 102.9,
  "enabled": false
}
```

