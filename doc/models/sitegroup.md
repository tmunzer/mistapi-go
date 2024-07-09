
# Sitegroup

Sites Group

## Structure

`Sitegroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteIds` | `[]uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 64.44,
  "id": "00000ee6-0000-0000-0000-000000000000",
  "modified_time": 14.52,
  "site_ids": [
    "000003ec-0000-0000-0000-000000000000",
    "000003eb-0000-0000-0000-000000000000",
    "000003ea-0000-0000-0000-000000000000"
  ]
}
```

