
# Search Wan Client

## Structure

`SearchWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`*models.ClientsWanStats`](../../doc/models/clients-wan-stats.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 166,
  "limit": 4,
  "results": {
    "When": "When8",
    "hostname": [
      "hostname6",
      "hostname7",
      "hostname8"
    ],
    "ip": [
      "ip7",
      "ip8"
    ],
    "last_hostname": "last_hostname8",
    "last_ip": "last_ip6"
  },
  "start": 124,
  "total": 98
}
```

