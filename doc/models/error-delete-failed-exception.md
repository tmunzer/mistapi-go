
# Error Delete Failed Exception

## Structure

`ErrorDeleteFailedException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |

## Example (as JSON)

```json
{
  "detail": "inventory not empty",
  "org_id": "000001d8-0000-0000-0000-000000000000"
}
```
