
# Response Msp Inventory Device

## Structure

`ResponseMspInventoryDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForSite` | `*bool` | Optional | - |
| `Mac` | `string` | Required | - |
| `Model` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `Serial` | `string` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Type` | `string` | Required | - |

## Example (as JSON)

```json
{
  "mac": "mac0",
  "model": "model4",
  "org_id": "0000032c-0000-0000-0000-000000000000",
  "serial": "serial6",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "type4",
  "for_site": false
}
```

