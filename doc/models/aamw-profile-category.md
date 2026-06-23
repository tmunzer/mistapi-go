
# Aamw Profile Category

File category rule for Advanced Anti Malware inspection

## Structure

`AamwProfileCategory`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Category` | [`*models.AamwProfileCategoryCategoryEnum`](../../doc/models/aamw-profile-category-category-enum.md) | Optional | enum: `archive`, `document`, `pdf`, `executable`, `rich_application`, `library`, `os_package`, `mobile`, `java`, `configuration`, `script` |
| `HashLookupOnly` | `*bool` | Optional | Whether files in this category use hash lookup without full file analysis<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    aamwProfileCategory := models.AamwProfileCategory{
        Category:             models.ToPointer(models.AamwProfileCategoryCategoryEnum_CONFIGURATION),
        HashLookupOnly:       models.ToPointer(false),
    }

}
```

