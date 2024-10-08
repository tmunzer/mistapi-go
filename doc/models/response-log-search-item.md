
# Response Log Search Item

## Structure

`ResponseLogSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `*uuid.UUID` | Required | admin id |
| `AdminName` | `*string` | Required | name of the admin that performs the action |
| `After` | `*interface{}` | Optional | field values after the change |
| `Before` | `*interface{}` | Optional | field values prior to the change |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Message` | `string` | Required | log message |
| `OrgId` | `uuid.UUID` | Required | org id |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | start time, in epoch |

## Example (as JSON)

```json
{
  "admin_id": "00001792-0000-0000-0000-000000000000",
  "admin_name": "admin_name8",
  "message": "message0",
  "org_id": "00001500-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 162.78,
  "after": {
    "key1": "val1",
    "key2": "val2"
  },
  "before": {
    "key1": "val1",
    "key2": "val2"
  },
  "for_site": false,
  "id": "00001428-0000-0000-0000-000000000000"
}
```

