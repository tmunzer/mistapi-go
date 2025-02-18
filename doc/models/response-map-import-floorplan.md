
# Response Map Import Floorplan

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "action": "action8",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "00000158-0000-0000-0000-000000000000",
  "name": "name0",
  "reason": "reason6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

