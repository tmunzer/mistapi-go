
# Response Map Import Ap

## Structure

`ResponseMapImportAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`models.ResponseMapImportApActionEnum`](../../doc/models/response-map-import-ap-action-enum.md) | Required | enum: `assigned-named-placed`, `assigned-placed`, `ignored`, `named-placed`, `placed` |
| `FloorplanId` | `uuid.UUID` | Required | - |
| `Height` | `*float64` | Optional | - |
| `Mac` | `string` | Required | - |
| `MapId` | `uuid.UUID` | Required | - |
| `Orientation` | `int` | Required | - |
| `Reason` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "action": "assigned-placed",
  "floorplan_id": "00000c80-0000-0000-0000-000000000000",
  "height": 86.12,
  "mac": "mac0",
  "map_id": "00001390-0000-0000-0000-000000000000",
  "orientation": 240,
  "reason": "reason8"
}
```

