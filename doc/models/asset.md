
# Asset

Asset

## Structure

`Asset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Mac` | `string` | Required | bluetooth MAC |
| `MapId` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | name / label of the device |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TagId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "mac4",
  "name": "name0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 59.3,
  "for_site": false,
  "id": "00001e14-0000-0000-0000-000000000000",
  "map_id": "00000018-0000-0000-0000-000000000000",
  "modified_time": 19.66
}
```

