
# Search Wan Usage

## Structure

`SearchWanUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.WanUsages`](../../doc/models/wan-usages.md) | Optional | - |
| `Start` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "end": 173.0,
  "limit": 22,
  "results": [
    {
      "mac": "mac0",
      "path_type": "path_type8",
      "path_weight": 242,
      "peer_mac": "peer_mac6",
      "peer_port_id": "peer_port_id4"
    }
  ],
  "start": 129.06
}
```

