
# Response Tunnel Search

## Structure

`ResponseTunnelSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseTunnelSearchItem`](../../doc/models/containers/response-tunnel-search-item.md) | Required | - |
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
      "fwupdate": {
        "progress": 100,
        "status": "inprogress",
        "status_id": 70,
        "timestamp": 147.68,
        "will_retry": false
      },
      "last_seen": 249.6,
      "mtu": 34,
      "remote_ip": "remote_ip4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "ap": "ap4",
      "for_site": false,
      "fwupdate": {
        "progress": 100,
        "status": "inprogress",
        "status_id": 70,
        "timestamp": 147.68,
        "will_retry": false
      },
      "last_seen": 249.6,
      "mtu": 34,
      "remote_ip": "remote_ip4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 196,
  "total": 230,
  "next": "next2"
}
```

