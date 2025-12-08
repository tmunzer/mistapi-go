
# Template Exceptions

Where this template should not be applied to (takes precedence)

## Structure

`TemplateExceptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids |

## Example (as JSON)

```json
{
  "site_ids": [
    "00000d74-0000-0000-0000-000000000000",
    "00000d73-0000-0000-0000-000000000000",
    "00000d72-0000-0000-0000-000000000000"
  ],
  "sitegroup_ids": [
    "0000039e-0000-0000-0000-000000000000"
  ]
}
```

