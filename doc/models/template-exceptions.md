
# Template Exceptions

Where this template should not be applied to (takes precedence)

*This model accepts additional fields of type interface{}.*

## Structure

`TemplateExceptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

