
# Secintel Profile Profile

SecIntel action setting for a specific feed category

## Structure

`SecintelProfileProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.SecintelProfileProfileActionEnum`](../../doc/models/secintel-profile-profile-action-enum.md) | Optional | enum: `default`, `standard`, `strict`<br><br>**Default**: `"default"` |
| `Category` | [`*models.SecintelProfileProfileCategoryEnum`](../../doc/models/secintel-profile-profile-category-enum.md) | Optional | enum: `CC`, `IH` (Infected Host), `DNS` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    secintelProfileProfile := models.SecintelProfileProfile{
        Action:               models.ToPointer(models.SecintelProfileProfileActionEnum_ENUMDEFAULT),
        Category:             models.ToPointer(models.SecintelProfileProfileCategoryEnum_CC),
    }

}
```

