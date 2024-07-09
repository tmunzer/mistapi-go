
# Response Map Import Floorplan

## Structure

`ResponseMapImportFloorplan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | `string` | Required | - |
| `Id` | `uuid.UUID` | Required | - |
| `MapId` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `Reason` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "action": "action0",
  "id": "00001e2a-0000-0000-0000-000000000000",
  "map_id": "00000002-0000-0000-0000-000000000000",
  "name": "name2",
  "reason": "reason2"
}
```

