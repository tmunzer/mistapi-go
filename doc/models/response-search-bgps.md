
# Response Search Bgps

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSearchBgps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.BgpStats`](../../doc/models/bgp-stats.md) | Optional | - |
| `Start` | `*float64` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 237.36,
  "limit": 242,
  "next": "next4",
  "results": [
    {
      "evpn_overlay": false,
      "for_overlay": false,
      "local_as": "String3",
      "mac": "mac0",
      "model": "model4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 193.42,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

