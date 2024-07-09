
# Template Exceptions

where this template should not be applied to (takes precedence)

## Structure

`TemplateExceptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteIds` | `[]uuid.UUID` | Optional | list of site ids |
| `SitegroupIds` | `[]uuid.UUID` | Optional | list of sitegroup ids |

## Example (as JSON)

```json
{
  "site_ids": [
    "00000ad8-0000-0000-0000-000000000000",
    "00000ad9-0000-0000-0000-000000000000"
  ],
  "sitegroup_ids": [
    "00001262-0000-0000-0000-000000000000",
    "00001263-0000-0000-0000-000000000000"
  ]
}
```

