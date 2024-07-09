
# Template Applies

where this template should be applied to, can be org_id, site_ids, sitegroup_ids

## Structure

`TemplateApplies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteIds` | `[]uuid.UUID` | Optional | list of site ids |
| `SitegroupIds` | `[]uuid.UUID` | Optional | list of sitegroup ids |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_ids": [
    "00000fcc-0000-0000-0000-000000000000"
  ],
  "sitegroup_ids": [
    "000005f6-0000-0000-0000-000000000000",
    "000005f7-0000-0000-0000-000000000000",
    "000005f8-0000-0000-0000-000000000000"
  ]
}
```

