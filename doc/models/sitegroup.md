
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
  "name": "name8",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 2.38,
  "id": "000007d8-0000-0000-0000-000000000000",
  "modified_time": 76.58,
  "site_ids": [
    "00000322-0000-0000-0000-000000000000",
    "00000323-0000-0000-0000-000000000000",
    "00000324-0000-0000-0000-000000000000"
  ]
}
```

