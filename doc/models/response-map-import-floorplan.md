
# Response Map Import Floorplan

## Structure

`ResponseMapImportFloorplan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | `string` | Required | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `MapId` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `Reason` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "action": "action8",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "00000158-0000-0000-0000-000000000000",
  "name": "name0",
  "reason": "reason6"
}
```

