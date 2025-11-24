
# Ospf Peer Stats Search Result

*This model accepts additional fields of type interface{}.*

## Structure

`OspfPeerStatsSearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.OspfPeerStatsSearchResultsItems`](../../doc/models/ospf-peer-stats-search-results-items.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1711035686,
  "limit": 10,
  "start": 1710949286,
  "total": 232,
  "next": "next0",
  "results": [
    {
      "dead_time": 140,
      "mac": "mac0",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "peer_ip": "peer_ip4",
      "port_id": "port_id6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "dead_time": 140,
      "mac": "mac0",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "peer_ip": "peer_ip4",
      "port_id": "port_id6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "dead_time": 140,
      "mac": "mac0",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "peer_ip": "peer_ip4",
      "port_id": "port_id6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

