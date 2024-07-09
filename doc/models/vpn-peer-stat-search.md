
# Vpn Peer Stat Search

## Structure

`VpnPeerStatSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.VpnPeerStat`](../../doc/models/vpn-peer-stat.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 121.26,
  "limit": 76,
  "results": [
    {
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "peer_site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "is_active": false,
      "last_seen": 165.16,
      "latency": 250.14,
      "mac": "mac0",
      "mos": 27.76
    }
  ],
  "start": 77.32,
  "total": 170,
  "next": "next6"
}
```

