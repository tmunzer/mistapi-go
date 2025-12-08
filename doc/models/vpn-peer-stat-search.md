
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
  "end": 142.42,
  "limit": 8,
  "results": [
    {
      "last_seen": 1470417522.0,
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "peer_site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "is_active": false,
      "jitter": 27.56,
      "latency": 250.14,
      "loss": 36.76
    }
  ],
  "start": 98.48,
  "total": 102,
  "next": "next8"
}
```

