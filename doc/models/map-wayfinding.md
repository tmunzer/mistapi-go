
# Map Wayfinding

Properties related to wayfinding

*This model accepts additional fields of type interface{}.*

## Structure

`MapWayfinding`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Micello` | [`*models.MapWayfindingMicello`](../../doc/models/map-wayfinding-micello.md) | Optional | - |
| `SnapToPath` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "micello": {
    "account_key": "account_key8",
    "default_level_id": 220,
    "map_id": "map_id8",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "snap_to_path": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

