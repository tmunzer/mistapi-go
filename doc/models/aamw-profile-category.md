
# Aamw Profile Category

*This model accepts additional fields of type interface{}.*

## Structure

`AamwProfileCategory`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Category` | [`*models.AamwProfileCategoryCategoryEnum`](../../doc/models/aamw-profile-category-category-enum.md) | Optional | enum: `archive`, `document`, `pdf`, `executable`, `rich_application`, `library`, `os_package`, `mobile`, `java`, `configuration`, `script` |
| `HashLookupOnly` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "hash_lookup_only": false,
  "category": "pdf",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

