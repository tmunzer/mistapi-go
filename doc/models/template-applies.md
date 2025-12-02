
# Template Applies

Where this template should be applied to, can be org_id, site_ids, sitegroup_ids

## Structure

`TemplateApplies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_ids": [
    "00000a84-0000-0000-0000-000000000000",
    "00000a83-0000-0000-0000-000000000000",
    "00000a82-0000-0000-0000-000000000000"
  ],
  "sitegroup_ids": [
    "000020da-0000-0000-0000-000000000000",
    "000020db-0000-0000-0000-000000000000"
  ]
}
```

