
# Response Port Stats Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePortStatsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.StatsSwitchPort`](../../doc/models/stats-switch-port.md) | Required | - |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 204,
  "limit": 222,
  "results": [
    {
      "full_duplex": true,
      "mac": "5c4527a96580",
      "neighbor_mac": "64d814353400",
      "neighbor_port_desc": "GigabitEthernet1/0/21",
      "neighbor_system_name": "CORP-D-SW-2",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "port_id": "ge-0/0/0",
      "port_mac": "5c4527a96580",
      "rx_bytes": 4563443626,
      "rx_pkts": 106,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "speed": 1000,
      "tx_bytes": 11299516780,
      "tx_pkts": 492176,
      "type": "gateway",
      "xcvr_model": "SFP+-10G-SR",
      "xcvr_part_number": "740-021487",
      "xcvr_serial": "N6AA9HT",
      "active": false,
      "auth_state": "authenticated",
      "for_site": false,
      "jitter": 27.56,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 162,
  "total": 196,
  "next": "next8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

