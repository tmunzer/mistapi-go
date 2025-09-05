
# Response Log Search Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseLogSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `models.Optional[uuid.UUID]` | Optional | admin id |
| `AdminName` | `models.Optional[string]` | Optional | Name of the admin that performs the action |
| `After` | `*interface{}` | Optional | field values after the change |
| `Before` | `*interface{}` | Optional | field values prior to the change |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Message` | `string` | Required | log message |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `models.Optional[uuid.UUID]` | Optional | - |
| `SrcIp` | `*string` | Optional | sender source ip address |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "message": "message0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "timestamp": 162.78,
  "admin_id": "00001792-0000-0000-0000-000000000000",
  "admin_name": "admin_name8",
  "after": {
    "key1": "val1",
    "key2": "val2"
  },
  "before": {
    "key1": "val1",
    "key2": "val2"
  },
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

