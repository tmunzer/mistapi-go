
# Search Wan Client

## Structure

`SearchWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.StatsWanClient`](../../doc/models/stats-wan-client.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 206,
  "limit": 220,
  "next": "next2",
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
      "ip_src": "ip_src6"
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
      "ip_src": "ip_src6"
    }
  ],
  "start": 164
}
```

