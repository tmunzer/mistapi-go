
# Response Tunnel Search

## Structure

`ResponseTunnelSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseTunnelSearchResults`](../../doc/models/containers/response-tunnel-search-results.md) | Required | This is Array of a container for any-of cases. |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 178,
  "limit": 248,
  "results": [
    {
      "ap": "ap4",
      "for_site": false,
      "last_seen": 85.7,
      "mxcluster_id": "00002490-0000-0000-0000-000000000000",
      "mxedge_id": "000023c4-0000-0000-0000-000000000000"
    },
    {
      "ap": "ap4",
      "for_site": false,
      "last_seen": 85.7,
      "mxcluster_id": "00002490-0000-0000-0000-000000000000",
      "mxedge_id": "000023c4-0000-0000-0000-000000000000"
    },
    {
      "ap": "ap4",
      "for_site": false,
      "last_seen": 85.7,
      "mxcluster_id": "00002490-0000-0000-0000-000000000000",
      "mxedge_id": "000023c4-0000-0000-0000-000000000000"
    }
  ],
  "start": 136,
  "total": 86,
  "next": "next6"
}
```

