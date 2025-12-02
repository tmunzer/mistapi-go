
# Response Mxedge Search

## Structure

`ResponseMxedgeSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.SearchMxedge`](../../doc/models/search-mxedge.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1694708579,
  "limit": 10,
  "start": 1694622179,
  "total": 2,
  "next": "next0",
  "results": [
    {
      "distro": "distro8",
      "last_seen": 165.16,
      "model": "model4",
      "mxcluster_id": "00001c8a-0000-0000-0000-000000000000",
      "mxedge_id": "00001bbe-0000-0000-0000-000000000000"
    },
    {
      "distro": "distro8",
      "last_seen": 165.16,
      "model": "model4",
      "mxcluster_id": "00001c8a-0000-0000-0000-000000000000",
      "mxedge_id": "00001bbe-0000-0000-0000-000000000000"
    },
    {
      "distro": "distro8",
      "last_seen": 165.16,
      "model": "model4",
      "mxcluster_id": "00001c8a-0000-0000-0000-000000000000",
      "mxedge_id": "00001bbe-0000-0000-0000-000000000000"
    }
  ]
}
```

