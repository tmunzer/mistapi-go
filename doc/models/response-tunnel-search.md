
# Response Tunnel Search

## Structure

`ResponseTunnelSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseTunnelSearchItem2`](../../doc/models/containers/response-tunnel-search-item-2.md) | Required | - |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 238,
  "limit": 188,
  "results": [
    {
      "ap": "ap4",
      "for_site": false,
      "last_seen": 153.2,
      "mxcluster_id": "000017de-0000-0000-0000-000000000000",
      "mxedge_id": "00001712-0000-0000-0000-000000000000"
    },
    {
      "ap": "ap4",
      "for_site": false,
      "last_seen": 153.2,
      "mxcluster_id": "000017de-0000-0000-0000-000000000000",
      "mxedge_id": "00001712-0000-0000-0000-000000000000"
    }
  ],
  "start": 196,
  "total": 230,
  "next": "next2"
}
```

