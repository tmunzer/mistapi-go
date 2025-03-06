
# Audit Log

*This model accepts additional fields of type interface{}.*

## Structure

`AuditLog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `uuid.UUID` | Required | ID of the administrator |
| `AdminName` | `string` | Required | - |
| `After` | `*interface{}` | Optional | Field values after the change |
| `Before` | `*interface{}` | Optional | Field values prior to the change |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Message` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "admin_id": "456b7016-a916-a4b1-78dd-72b947c152b7",
  "admin_name": "admin_name6",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "message": "message8",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 132.46,
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

