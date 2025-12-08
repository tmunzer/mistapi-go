
# Response Port Stats Search

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
      "rx_bps": 60003,
      "rx_bytes": 8515104416,
      "rx_pkts": 57770567,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "speed": 1000,
      "tx_bps": 634301,
      "tx_bytes": 211217389682,
      "tx_pkts": 812204062,
      "type": "gateway",
      "xcvr_model": "SFP+-10G-SR",
      "xcvr_part_number": "740-021487",
      "xcvr_serial": "N6AA9HT",
      "active": false,
      "auth_state": "authenticated",
      "disabled": false,
      "for_site": false
    }
  ],
  "start": 162,
  "total": 196,
  "next": "next8"
}
```

