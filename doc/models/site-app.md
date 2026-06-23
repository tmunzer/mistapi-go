
# Site App

Application summary returned by the site applications endpoint

## Structure

`SiteApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Group` | `string` | Required | Application category group for the site app<br><br>**Constraints**: *Minimum Length*: `1` |
| `Key` | `string` | Required | Stable application key for the site app<br><br>**Constraints**: *Minimum Length*: `1` |
| `Name` | `string` | Required | Display name of the site app<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteApp := models.SiteApp{
        Group:                "group4",
        Key:                  "key6",
        Name:                 "name6",
    }

}
```

