
# Map Sitesurvey Path Items

## Structure

`MapSitesurveyPathItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coordinate` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Name` | `*string` | Optional | - |
| `Nodes` | [`[]models.MapNode`](../../doc/models/map-node.md) | Optional | **Constraints**: *Minimum Items*: `0` |

## Example (as JSON)

```json
{
  "coordinate": "actual",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Default",
  "nodes": [
    {
      "edges": {
        "key0": "edges6"
      },
      "name": "name6",
      "position": {
        "x": 224.72,
        "y": 100.0
      }
    },
    {
      "edges": {
        "key0": "edges6"
      },
      "name": "name6",
      "position": {
        "x": 224.72,
        "y": 100.0
      }
    }
  ]
}
```

