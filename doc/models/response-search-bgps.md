
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
  "end": 237.36,
  "limit": 242,
  "results": [
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": 18,
      "mac": "mac0",
      "model": "model4"
    }
  ],
  "start": 193.42,
  "total": 80
}
```

