
# Response Search Bgps

## Structure

`ResponseSearchBgps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.BgpStats`](../../doc/models/bgp-stats.md) | Optional | - |
| `Start` | `*float64` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 127.0,
  "limit": 14,
  "results": [
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": 18,
      "mac": "mac0",
      "neighbor": "neighbor6"
    },
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": 18,
      "mac": "mac0",
      "neighbor": "neighbor6"
    },
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": 18,
      "mac": "mac0",
      "neighbor": "neighbor6"
    }
  ],
  "start": 83.06,
  "total": 108
}
```

