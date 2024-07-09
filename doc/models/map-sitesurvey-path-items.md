
# Map Sitesurvey Path Items

## Structure

`MapSitesurveyPathItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coordinate` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Nodes` | [`[]models.MapNode`](../../doc/models/map-node.md) | Optional | **Constraints**: *Minimum Items*: `0` |

## Example (as JSON)

```json
{
  "coordinate": "actual",
  "id": "63eda950-c6da-11e4-a628-60f81dd250ce",
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

