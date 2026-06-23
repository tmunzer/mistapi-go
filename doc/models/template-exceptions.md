
# Template Exceptions

Where this template should not be applied to (takes precedence)

## Structure

`TemplateExceptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    templateExceptions := models.TemplateExceptions{
        SiteIds:              []uuid.UUID{
            uuid.MustParse("0000238e-0000-0000-0000-000000000000"),
            uuid.MustParse("0000238f-0000-0000-0000-000000000000"),
            uuid.MustParse("00002390-0000-0000-0000-000000000000"),
        },
        SitegroupIds:         []uuid.UUID{
            uuid.MustParse("000020bc-0000-0000-0000-000000000000"),
        },
    }

}
```

