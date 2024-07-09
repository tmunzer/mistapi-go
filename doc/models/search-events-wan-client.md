
# Search Events Wan Client

## Structure

`SearchEventsWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`*models.EventsClientWan`](../../doc/models/events-client-wan.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 200,
  "limit": 226,
  "results": {
    "When": "When8",
    "ev_type": "ev_type4",
    "metadata": {
      "key1": "val1",
      "key2": "val2"
    },
    "org_id": "00002492-0000-0000-0000-000000000000",
    "random_mac": false
  },
  "start": 158,
  "total": 64
}
```

