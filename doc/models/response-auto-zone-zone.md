
# Response Auto Zone Zone

A list of suggested zones to review and accept for a given map

## Structure

`ResponseAutoZoneZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | The name of the suggested zone |
| `Vertices` | [`[]models.ResponseAutoZoneZoneVertex`](../../doc/models/response-auto-zone-zone-vertex.md) | Optional | A list of of points comprising the zones map location in pixels |

## Example (as JSON)

```json
{
  "name": "zone1",
  "vertices": [
    {
      "x": 16,
      "y": 88
    },
    {
      "x": 16,
      "y": 88
    },
    {
      "x": 16,
      "y": 88
    }
  ]
}
```

