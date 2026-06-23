
# Mac Table Stats

MAC table capacity and usage statistics

## Structure

`MacTableStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MacTableCount` | `*int` | Optional | Number of learned MAC table entries currently present |
| `MaxMacEntriesSupported` | `*int` | Optional | Maximum number of MAC table entries supported |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    macTableStats := models.MacTableStats{
        MacTableCount:          models.ToPointer(78),
        MaxMacEntriesSupported: models.ToPointer(222),
    }

}
```

