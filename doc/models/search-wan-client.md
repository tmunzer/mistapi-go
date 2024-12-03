
# Search Wan Client

*This model accepts additional fields of type interface{}.*

## Structure

`SearchWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.StatsWanClient`](../../doc/models/stats-wan-client.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 206,
  "limit": 220,
  "results": [
    {
      "dhcp_expire_time": 124.26,
      "dhcp_start_time": 9.94,
      "hostname": [
        "hostname6",
        "hostname7",
        "hostname8"
      ],
      "ip": [
        "ip7",
        "ip8"
      ],
      "ip_src": "ip_src6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "dhcp_expire_time": 124.26,
      "dhcp_start_time": 9.94,
      "hostname": [
        "hostname6",
        "hostname7",
        "hostname8"
      ],
      "ip": [
        "ip7",
        "ip8"
      ],
      "ip_src": "ip_src6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 164,
  "total": 58,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

