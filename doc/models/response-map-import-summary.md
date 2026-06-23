
# Response Map Import Summary

Counts summarizing assignments made during the map import

## Structure

`ResponseMapImportSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumApAssigned` | `int` | Required | Number of AP placements assigned during the map import |
| `NumInvAssigned` | `int` | Required | Number of inventory records assigned to the site during the map import |
| `NumMapAssigned` | `int` | Required | Number of map floorplans assigned during the map import |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseMapImportSummary := models.ResponseMapImportSummary{
        NumApAssigned:        100,
        NumInvAssigned:       208,
        NumMapAssigned:       198,
    }

}
```

