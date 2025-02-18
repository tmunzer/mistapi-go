
# Stats Switch Clients Stats

*This model accepts additional fields of type interface{}.*

## Structure

`StatsSwitchClientsStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Total` | [`*models.StatsSwitchClientsStatsTotal`](../../doc/models/stats-switch-clients-stats-total.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "total": {
    "num_aps": [
      23,
      22,
      21
    ],
    "num_wired_clients": 222,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

