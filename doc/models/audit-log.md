
# Audit Log

## Structure

`AuditLog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `uuid.UUID` | Required | - |
| `AdminName` | `string` | Required | - |
| `After` | `*interface{}` | Optional | field values after the change |
| `Before` | `*interface{}` | Optional | field values prior to the change |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `uuid.UUID` | Required | - |
| `Message` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "admin_id": "00002458-0000-0000-0000-000000000000",
  "admin_name": "admin_name2",
  "id": "000020ee-0000-0000-0000-000000000000",
  "message": "message0",
  "org_id": "0000054a-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 104.52,
  "after": {
    "key1": "val1",
    "key2": "val2"
  },
  "before": {
    "key1": "val1",
    "key2": "val2"
  },
  "for_site": false
}
```
