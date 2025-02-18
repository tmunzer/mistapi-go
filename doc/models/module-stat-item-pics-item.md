
# Module Stat Item Pics Item

*This model accepts additional fields of type interface{}.*

## Structure

`ModuleStatItemPicsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Index` | `*int` | Optional | - |
| `ModelNumber` | `*string` | Optional | - |
| `PortGroups` | [`[]models.ModuleStatItemPicsItemPortGroupsItem`](../../doc/models/module-stat-item-pics-item-port-groups-item.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "index": 36,
  "model_number": "model_number8",
  "port_groups": [
    {
      "count": 32,
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "count": 32,
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "count": 32,
      "type": "type6",
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

