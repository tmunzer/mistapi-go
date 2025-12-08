
# Aamw Profile Category

## Structure

`AamwProfileCategory`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Category` | [`*models.AamwProfileCategoryCategoryEnum`](../../doc/models/aamw-profile-category-category-enum.md) | Optional | enum: `archive`, `document`, `pdf`, `executable`, `rich_application`, `library`, `os_package`, `mobile`, `java`, `configuration`, `script` |
| `HashLookupOnly` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "hash_lookup_only": false,
  "category": "pdf"
}
```

