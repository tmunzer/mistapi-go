
# Zone

Zone

## Structure

`Zone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | map where this zone is defined |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | The name of the zone |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Vertices` | [`[]models.ZoneVertex`](../../doc/models/zone-vertex.md) | Optional | vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "vertices": [
    {
      "x": 732,
      "y": 1821
    },
    {
      "x": 732.5,
      "y": 1731
    },
    {
      "x": 837.5,
      "y": 1731.5
    },
    {
      "x": 839,
      "y": 1821
    }
  ],
  "created_time": 89.02,
  "for_site": false,
  "id": "00001880-0000-0000-0000-000000000000",
  "map_id": "000005ac-0000-0000-0000-000000000000",
  "modified_time": 245.94
}
```

